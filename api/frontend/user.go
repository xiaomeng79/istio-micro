package frontend

import (
	"github.com/labstack/echo"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/internal/api"
	"github.com/xiaomeng79/istio-micro/internal/utils"
	pb "github.com/xiaomeng79/istio-micro/srv/user/proto"
)

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
