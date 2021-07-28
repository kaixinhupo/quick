package repository

import "github.com/kaixinhupo/quick/dao/entity"

type UserRepository interface {
	InsertUser(user *entity.UserEntity) (*entity.UserEntity, error)
}

