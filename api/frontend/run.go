package frontend

import (
	"net/http"
	"time"

	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/api"
	"github.com/xiaomeng79/istio-micro/internal/wrapper"
	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xiaomeng79/go-log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// 定义services名称
const (
	SN = "api-frontend"
)

var (
	UserClient pb.UserServiceClient
)

// 运行
func Run() {
	// 初始化
	cinit.InitOption(SN, cinit.Trace)
	// 建立客户端连接
	grOpts := []grpc_retry.CallOption{
		grpc_retry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(15 * time.Second),
	}
	conn, err := grpc.Dial(cinit.Config.SrvUser.Address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
			wrapper.LoggingUnaryClientInterceptor(),
			grpc_retry.UnaryClientInterceptor(grOpts...),
			// wrapper.TraceingUnaryClientInterceptor(),
		)))
	if err != nil {
		log.Fatal("连接user服务失败" + err.Error())
	}
	// 注册客户端
	UserClient = pb.NewUserServiceClient(conn)

	// 注册路由
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))
	// e.Use(common.Opentracing)
	e.Use(api.TraceHeader)
	// e.Use(api.NoSign)
	// 总分组
	g := e.Group("/frontend/v1")
	// g := e.Group("/backend/v1", api.JWT)
	// 用户
	g.GET("/user", UserQueryAll)

	// check
	check := e.Group("/frontend/check")
	check.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	//  启动service
	if err := e.Start(cinit.Config.APIFrontend.Port); err != nil {
		log.Fatal("启动服务失败" + err.Error())
	}
	defer conn.Close()
}
