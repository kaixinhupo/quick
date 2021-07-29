package contract

import (
	"github.com/kaixinhupo/quick/infrastruture/web"
	"github.com/kaixinhupo/quick/model"
)


type UserService interface {
	// 查询记录列表
	Query(req *model.UserQueryReq) (*web.PageResp,error)
	// 查询单条记录
	Item(id int64) (*model.UserDetailResp, error)
	// 创建记录
	Create(req *model.UserInfoReq) (*model.UserDetailResp, error)
	// 更新记录
	Update(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error)
	// 修改记录
	Patch(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error)
	// 删除记录
	Delete(id int64) error
}