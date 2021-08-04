package repository

import (
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)

type UserRepository interface {
	// Count 查询数量
	Count(req *model.UserQueryReq, session *xorm.Session) int64
	// List 查询列表
	List(req *model.UserQueryReq, session *xorm.Session) []*entity.UserEntity
	// Entry 根据主键查询记录
	Entry(id int64, session *xorm.Session) *entity.UserEntity
	// Insert 插入单个记录
	Insert(entity *entity.UserEntity, session *xorm.Session) (*entity.UserEntity, error)
	// Update 更新记录
	Update(entity *entity.UserEntity, allFields bool, session *xorm.Session) (int64, error)
	// Delete 删除记录
	Delete(id int64, session *xorm.Session) error
}
