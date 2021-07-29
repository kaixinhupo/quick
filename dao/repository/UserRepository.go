package repository

import (
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)


type UserRepository interface {
	// 插入单个记录
	Insert(user *entity.UserEntity, session *xorm.Session) (*entity.UserEntity, error)
	// 查询数量
	Count(param *model.UserQueryReq, session *xorm.Session) (int64)
	// 查询列表
	List(param *model.UserQueryReq, session *xorm.Session) ([]*entity.UserEntity)
}

