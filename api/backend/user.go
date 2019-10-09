package backend

import (
	"strings"

	"github.com/xiaomeng79/istio-micro/internal/api"
	"github.com/xiaomeng79/istio-micro/internal/jwt"
	"github.com/xiaomeng79/istio-micro/internal/utils"
	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"

	"github.com/labstack/echo"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/go-utils/crypto"
)

//  添加
func UserAdd(c echo.Context) error {
	ctx := c.Request().Context()
	//  解析请求参数
	_req := new(pb.UserBase)
	err := c.Bind(&_req)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	_rsp, err := UserClient.UserAdd(ctx, _req)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RPCErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//  修改
func UserUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	//  解析请求参数
	_req := new(pb.UserBase)
	err := c.Bind(&_req)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	id, err := utils.S2ID(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	_req.Id = id
	_rsp, err := UserClient.UserUpdate(ctx, _req)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RPCErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//  删除
func UserDelete(c echo.Context) error {
	ctx := c.Request().Context()
	// 解析请求参数
	_req := new(pb.UserID)
	id, err := utils.S2ID(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}

	_req.Id = id
	_rsp, err := UserClient.UserDelete(ctx, _req)
	if err != nil {
		// 解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RPCErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//  查询一个
func UserQueryOne(c echo.Context) error {
	ctx := c.Request().Context()
	//  解析请求参数
	_req := new(pb.UserID)
	id, err := utils.S2ID(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	_req.Id = id
	_rsp, err := UserClient.UserQueryOne(ctx, _req)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RPCErr(c, err)
	}
	return api.HandleSuccess(c, _rsp)
}

//  查询全部
func UserQueryAll(c echo.Context) error {
	ctx := c.Request().Context()

	pageIndex, err := utils.S2N(c.QueryParam("page_index"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	pageSize, err := utils.S2N(c.QueryParam("page_size"))
	if err != nil {
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}

	_page := new(pb.Page)
	_page.PageSize = pageSize
	_page.PageIndex = pageIndex

	_req := new(pb.UserAllOption)
	_req.Page = _page
	_rsp, err := UserClient.UserQueryAll(ctx, _req)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.RPCErr(c, err)
	}
	return api.HandleSuccess(c, _rsp.All, _rsp.Page)
}

const (
	UserName = "root"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//  登录
func Login(c echo.Context) error {
	ctx := c.Request().Context()
	user := new(User)
	err := c.Bind(user)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.BusParamConvertError, err.Error())
	}
	if user.UserName != UserName || !strings.EqualFold(crypto.MD5(user.Password), "394810afe2fcb9a8210b80300d7ccc7a") {
		return api.HandleError(c, api.ReqNoAllow, "用户名或者密码不正确")
	}
	j := new(jwt.Msg)
	j.UserName = user.UserName
	token, err := jwt.Encode(*j)
	if err != nil {
		//  解析返回的错误信息
		log.Error(err.Error(), ctx)
		return api.HandleError(c, api.ServiceError, err.Error())
	}
	t := make(map[string]string)
	t["token"] = token
	return api.HandleSuccess(c, t)
}
