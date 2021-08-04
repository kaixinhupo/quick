package contract

import (
	"github.com/kaixinhupo/quick/infrastructure/web"
	"github.com/kaixinhupo/quick/model"
)

type RoleService interface {
	// Query 查询记录列表
	Query(req *model.RoleQueryReq) (*web.PageResp, error)
	// Item 查询单条记录
	Item(id int64) (*model.RoleDetailResp, error)
	// Create 创建记录
	Create(req *model.RoleInfoReq) (*model.RoleDetailResp, error)
	// Update 更新记录
	Update(id int64, req *model.RoleInfoReq) (*model.RoleDetailResp, error)
	// Patch 修改记录
	Patch(id int64, req *model.RoleInfoReq) (*model.RoleDetailResp, error)
	// Delete 删除记录
	Delete(id int64) error
}
