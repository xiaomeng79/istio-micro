package account

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/gateway"
	"github.com/xiaomeng79/istio-micro/internal/sqlupdate"
	"github.com/xiaomeng79/istio-micro/internal/wrapper"
	pb "github.com/xiaomeng79/istio-micro/srv/account/proto"
	"github.com/xiaomeng79/istio-micro/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	SN = "srv-account" //定义services名称
)

func Run() {

	//初始化,选着需要的组件
	cinit.InitOption(SN, "trace", "postgres", "metrics")
	// 更新sql
	err := execUpdateSql()
	if err != nil {
		log.Fatalf("updatesql error:%+v ", err)
	}
	// 更新版本号
	err = updateVersion()
	if err != nil {
		log.Fatalf("update version error:%+v ", err)
	}
	lis, err := net.Listen("tcp", cinit.Config.SrvAccount.Port)
	if err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
			wrapper.RecoveryUnaryInterceptor,
			wrapper.LoggingUnaryInterceptor,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
		)),
	)
	pb.RegisterAccountServiceServer(s, &AccountServer{})
	reflection.Register(s)

	// 开启一个网关服务
	ctx := context.Background()
	go gateway.Run(
		ctx,
		gateway.WithAddr(cinit.Config.SrvAccount.GateWayAddr),
		gateway.WithGRPCServer("tcp", cinit.Config.SrvAccount.Address),
		gateway.WithSwaggerDir(cinit.Config.SrvAccount.GateWaySwaggerDir),
		gateway.WithHandle(pb.RegisterAccountServiceHandler),
	)

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
}

// 获取旧的版本号
func getOldVersion() (string, error) {
	var oldversion string
	err := cinit.Pg.Get(&oldversion, `select version from sys_info where id =1 limit 1`)
	if err != nil {
		log.Errorf("获取旧版本号失败:%+v", err)
		return "", err
	}
	return oldversion, nil
}

// 获取旧的版本号
func updateVersion() error {
	_, err := cinit.Pg.Exec(`update sys_info set version=$1 where id=1`, version.Version)
	if err != nil {
		log.Errorf("更新版本号失败:%+v", err)
		return err
	}
	return nil
}

// 获取sql
func getSql() (string, error) {
	oldVersion, err := getOldVersion()
	if err != nil {
		return "", err
	}
	s := new(sqlupdate.SqlUpdate)
	sqls, err := s.GetSqls("./sqlupdate/record.json", oldVersion, version.Version)
	if err != nil {
		log.Errorf("获取执行sql失败:%+v", err)
		return "", err
	}
	return sqls, nil
}

// 执行sql
func execUpdateSql() error {
	sqls, err := getSql()
	if err != nil {
		if err == sqlupdate.NoSqlNeedUpdate {
			return nil
		}
		return err
	}
	tx, err := cinit.Pg.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sqls)
	if err != nil {
		log.Errorf("执行sql失败:%+v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
