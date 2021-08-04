package contract

{{#with Config}}
import (
	"{{Module}}/infrastructure/web"
	"{{Module}}/model"
)
{{/with}}

{{#with Meta}}
type {{ModelName}}Service interface {
	// Query 查询记录列表
	Query(req *model.{{ModelName}}QueryReq) (*web.PageResp,error)
	// Item 查询单条记录
	Item(id int64) (*model.{{ModelName}}DetailResp, error)
	// Create 创建记录
	Create(req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error)
	// Update 更新记录
	Update(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error)
	// Patch 修改记录
	Patch(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error)
	// Delete 删除记录
	Delete(id int64) error
}
{{/with}}