package contract

import (
	"github.com/kaixinhupo/quick/infrastructure/web"
	"github.com/kaixinhupo/quick/model"
)

type UserService interface {
	// Query 查询记录列表
	Query(req *model.UserQueryReq) (*web.PageResp, error)
	// Item 查询单条记录
	Item(id int64) (*model.UserDetailResp, error)
	// Create 创建记录
	Create(req *model.UserInfoReq) (*model.UserDetailResp, error)
	// Update 更新记录
	Update(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error)
	// Patch 修改记录
	Patch(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error)
	// Delete 删除记录
	Delete(id int64) error
}
