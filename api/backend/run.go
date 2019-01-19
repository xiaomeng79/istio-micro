package backend

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/api"
	"github.com/xiaomeng79/istio-micro/internal/wrapper"
	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net/http"
	"time"
)

//定义services名称
const (
	SN = "api-backend"
)

var (
	UserClient pb.UserServiceClient
)

//运行
func Run() {
	//初始化
	cinit.InitOption(SN, "trace")
	//建立客户端连接
	gr_opts := []grpc_retry.CallOption{
		grpc_retry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(15 * time.Second),
	}
	conn, err := grpc.Dial(cinit.Config.SrvUser.Address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
			wrapper.LoggingUnaryClientInterceptor(),
			grpc_retry.UnaryClientInterceptor(gr_opts...),
		//wrapper.TraceingUnaryClientInterceptor(),
		)))
	if err != nil {
		log.Fatal("连接user服务失败" + err.Error())
	}
	defer conn.Close()
	//注册客户端
	UserClient = pb.NewUserServiceClient(conn)

	//注册路由
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))
	//e.Use(common.Opentracing)
	e.Use(api.TraceHeader)
	//e.Use(api.NoSign)

	//总分组
	g := e.Group("/backend/v1")
	//g := e.Group("/backend/v1", api.JWT)
	//用户
	g.POST("/user", UserAdd)
	g.PUT("/user/:id", UserUpdate)
	g.DELETE("/user/:id", UserDelete)
	g.GET("/user/:id", UserQueryOne)
	g.GET("/user", UserQueryAll)

	//总分组
	gu := e.Group("/backend/v1")
	gu.POST("/login", Login)
	//check
	check := e.Group("/backend/check")
	check.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	//启动service
	if err := e.Start(cinit.Config.ApiBackend.Port); err != nil {
		log.Fatal("启动服务失败" + err.Error())
	}
}
