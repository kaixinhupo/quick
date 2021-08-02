package contract

{{#with Config}}
import (
	"{{Module}}/infrastruture/web"
	"{{Module}}/model"
)
{{/with}}

{{#with Meta}}
type {{ModelName}}Service interface {
	// 查询记录列表
	Query(req *model.{{ModelName}}QueryReq) (*web.PageResp,error)
	// 查询单条记录
	Item(id int64) (*model.{{ModelName}}DetailResp, error)
	// 创建记录
	Create(req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error)
	// 更新记录
	Update(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error)
	// 修改记录
	Patch(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error)
	// 删除记录
	Delete(id int64) error
}
{{/with}}