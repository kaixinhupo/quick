package database

import (
	"github.com/kaixinhupo/quick/dao/entity"
)


type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return new(UserRepositoryImpl)
}

func (impl *UserRepositoryImpl) InsertUser(user *entity.UserEntity) (*entity.UserEntity, error) {
	return user,nil
}