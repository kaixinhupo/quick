package repository

import (
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)

type RoleRepository interface {
	// Count 查询数量
	Count(req *model.RoleQueryReq, session *xorm.Session) int64
	// List 查询列表
	List(req *model.RoleQueryReq, session *xorm.Session) []*entity.RoleEntity
	// Entry 根据主键查询记录
	Entry(id int64, session *xorm.Session) *entity.RoleEntity
	// Insert 插入单个记录
	Insert(entity *entity.RoleEntity, session *xorm.Session) (*entity.RoleEntity, error)
	// Update 更新记录
	Update(entity *entity.RoleEntity, allFields bool, session *xorm.Session) (int64, error)
	// Delete 删除记录
	Delete(id int64, session *xorm.Session) error
}
