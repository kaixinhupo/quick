package repository

{{#with Config}}
import (
	"{{Module}}/dao/entity"
	"{{Module}}/model"
	"xorm.io/xorm"
)
{{/with}}
{{#with Meta}}
type {{ModelName}}Repository interface {
	// Count 查询数量
	Count(req *model.{{ModelName}}QueryReq, session *xorm.Session) int64
	// List 查询列表
	List(req *model.{{ModelName}}QueryReq, session *xorm.Session) []*entity.{{ModelName}}Entity
	// Entry 根据主键查询记录
	Entry(id int64, session *xorm.Session) *entity.{{ModelName}}Entity
	// Insert 插入单个记录
	Insert(entity *entity.{{ModelName}}Entity, session *xorm.Session) (*entity.{{ModelName}}Entity, error)
	// Update 更新记录
	Update(entity *entity.{{ModelName}}Entity, allFields bool, session *xorm.Session) (int64, error)
	// Delete 删除记录
	Delete(id int64, session *xorm.Session) error
}
{{/with}}
