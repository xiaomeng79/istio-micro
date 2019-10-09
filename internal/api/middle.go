package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/jwt"
	metrics2 "github.com/xiaomeng79/istio-micro/internal/metrics"

	"github.com/labstack/echo"
	"github.com/rcrowley/go-metrics"
	"github.com/xiaomeng79/go-log"
)

func VerifyParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获取span
		ctx := c.Request().Context()
		log.Infof("req:%+v", c.Request().Body, ctx)
		// 解析公共参数
		r := new(ReqParam)
		err := c.Bind(r)
		if err != nil {
			log.Info("解析参数错误:"+err.Error(), ctx)
			return HandleError(c, CommonParamConvertError)
		}
		log.Infof("decode req:%+v", r, ctx)
		// 验证公共参数
		err = r.Validate()
		if err != nil {
			log.Info("验证参数错误:"+err.Error(), ctx)
			return HandleError(c, CommonParamConvertError, err.Error())
		}
		// 请求appsecret

		// 验证签名
		b, err := r.CompareSign()
		if !b {
			log.Info("获取appsecret"+err.Error(), ctx)
			return HandleError(c, CommonSignError)
		}
		// 通过验证，绑定参数
		c.Set(cinit.ReqParam, r)
		return next(c)
	}
}

// 验证jwt
func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获取span
		ctx := c.Request().Context()
		// 从请求头获取token信息
		jwtString := c.Request().Header.Get(cinit.JWTName)
		// log.Debug(jwtString, ctx)
		// 解析JWT
		auths := strings.Split(jwtString, " ")
		if !strings.EqualFold(auths[0], "BEARER") || auths[1] == " " {
			return HandleError(c, ReqTokenError, "token验证失败")
		}
		jwtmsg, err := jwt.Decode(strings.Trim(auths[1], " "))
		if err != nil {
			log.Info(err.Error(), ctx)
			return HandleError(c, ReqTokenError, "token验证失败")
		}
		c.Set(cinit.JWTMsg, jwtmsg)
		return next(c)
	}
}

func TraceHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Infof("header:%+v", c.Request().Header)
		return next(c)
	}
}

// verifyparam// trace
func NoSign(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获取span
		// c.Set(cinit.TRACE_CONTEXT, context.Background())
		return next(c)
	}
}

// metrics
func MetricsFunc(m *metrics2.Metrics) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			latency := stop.Sub(start)
			status := res.Status
			// Total request count.
			meter := metrics.GetOrRegisterMeter("status."+strconv.Itoa(status), m.GetRegistry())
			meter.Mark(1)

			// Request size in bytes.
			// meter = metrics.GetOrRegisterMeter(m.WithPrefix("status."+strconv.Itoa(status)), m.GetRegistry())
			// meter.Mark(req.ContentLength)

			// Request duration in nanoseconds.
			h := metrics.GetOrRegisterHistogram("h_status."+strconv.Itoa(status), m.GetRegistry(),
				metrics.NewExpDecaySample(1028, 0.015))
			h.Update(latency.Nanoseconds())

			// Response size in bytes.
			// meter = metrics.GetOrRegisterMeter(m.WithPrefix("status."+strconv.Itoa(status)), m.GetRegistry())
			// meter.Mark(res.Size)

			return nil
		}
	}
}
