package api

import (
	"github.com/labstack/echo"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/jwt"
	"strings"
)

//opentracing中间件
//func Opentracing(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		ctx, span, err := trace.TraceFromHeader(context.Background(), "api:"+c.Request().URL.Path, c.Request().Header)
//		if err == nil {
//			defer span.Finish()
//			c.Set(cinit.TRACE_CONTEXT, ctx)
//		} else {
//			c.Set(cinit.TRACE_CONTEXT, context.Background())
//		}
//		return next(c)
//	}
//}
//
//verifyparam
//trace
func VerifyParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取span
		ctx := c.Request().Context()
		log.Infof("req:%+v", c.Request().Body, ctx)
		//解析公共参数
		r := new(ReqParam)
		err := c.Bind(r)
		if err != nil {
			log.Info("解析参数错误:"+err.Error(), ctx)
			return HandleError(c, CommonParamConvertError)
		}
		log.Infof("decode req:%+v", r, ctx)
		//验证公共参数
		err = r.Validate()
		if err != nil {
			log.Info("验证参数错误:"+err.Error(), ctx)
			return HandleError(c, CommonParamConvertError, err.Error())
		}
		//请求appsecret

		//验证签名
		b, err := r.CompareSign()
		if !b {
			log.Info("获取appsecret"+err.Error(), ctx)
			return HandleError(c, CommonSignError)
		}
		//通过验证，绑定参数
		c.Set(cinit.REQ_PARAM, r)
		return next(c)
	}
}

//验证jwt
func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取span
		ctx := c.Request().Context()
		//从请求头获取token信息
		jwtString := c.Request().Header.Get(cinit.JWT_NAME)
		//log.Debug(jwtString, ctx)
		//解析JWT
		auths := strings.Split(jwtString, " ")
		if strings.ToUpper(auths[0]) != "BEARER" || auths[1] == " " {
			return HandleError(c, ReqTokenError, "token验证失败")
		}
		jwtmsg, err := jwt.Decode(strings.Trim(auths[1], " "))
		if err != nil {
			log.Info(err.Error(), ctx)
			return HandleError(c, ReqTokenError, "token验证失败")
		}
		c.Set(cinit.JWT_MSG, jwtmsg)
		return next(c)
	}
}

func TraceHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Infof("header:%+v", c.Request().Header)
		return next(c)
	}
}

//verifyparam
//trace
func NoSign(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取span
		//c.Set(cinit.TRACE_CONTEXT, context.Background())
		return next(c)
	}
}
