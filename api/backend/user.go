package backend

import (
	"github.com/labstack/echo"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/go-utils/crypto"
	"github.com/xiaomeng79/istio-micro/internal/api"
	"github.com/xiaomeng79/istio-micro/internal/jwt"
	"github.com/xiaomeng79/istio-micro/internal/utils"
	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"
	"strings"
)

//添加
func UserAdd(c echo.Context) error {
	ctx := c.Request().Context()
	//解析请求参数
	_req := new(pb.UserBase)
	err := c.Bind(&_req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	_rsp, err := UserClient.UserAdd(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RpcErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//修改
func UserUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	//解析请求参数
	_req := new(pb.UserBase)
	err := c.Bind(&_req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	id, err := utils.S2Id(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	_req.Id = id
	_rsp, err := UserClient.UserUpdate(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RpcErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//删除
func UserDelete(c echo.Context) error {
	ctx := c.Request().Context()
	//解析请求参数
	_req := new(pb.UserId)
	id, err := utils.S2Id(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}

	_req.Id = id
	_rsp, err := UserClient.UserDelete(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RpcErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//查询一个
func UserQueryOne(c echo.Context) error {

	ctx := c.Request().Context()
	//解析请求参数
	_req := new(pb.UserId)
	id, err := utils.S2Id(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	_req.Id = id
	_rsp, err := UserClient.UserQueryOne(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RpcErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//查询全部
func UserQueryAll(c echo.Context) error {
	ctx := c.Request().Context()

	page_index, err := utils.S2N(c.QueryParam("page_index"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	page_size, err := utils.S2N(c.QueryParam("page_size"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}

	_page := new(pb.Page)
	_page.PageSize = page_size
	_page.PageIndex = page_index

	_req := new(pb.UserAllOption)
	_req.Page = _page
	_rsp, err := UserClient.UserQueryAll(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RpcErr(c, err)
	}
	return api.HandleSuccess(c, _rsp.All, _rsp.Page)
}

const (
	UserName = "root"
	Password = "394810afe2fcb9a8210b80300d7ccc7a"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//登录
func Login(c echo.Context) error {
	ctx := c.Request().Context()
	user := new(User)
	err := c.Bind(user)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	if user.UserName != UserName || strings.ToLower(crypto.MD5(user.Password)) != Password {
		return api.HandleError(c, api.ReqNoAllow, "用户名或者密码不正确")
	}
	j := new(jwt.JWTMsg)
	j.UserName = user.UserName
	token, err := jwt.Encode(*j)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.ServiceError, err.Error())
	}
	t := make(map[string]string)
	t["token"] = token
	return api.HandleSuccess(c, t)
}
