package biz

import (
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/dao/repository"
	"github.com/kaixinhupo/quick/model"
)


type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl {
		userRepository:repo ,
	}
}

func (impl *UserServiceImpl) CreateUser(user *model.UserInfoReq) (*model.UserDetailResp, error) {
	impl.userRepository.InsertUser(&entity.UserEntity{})
	return &model.UserDetailResp{},nil
}