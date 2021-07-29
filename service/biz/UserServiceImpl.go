package biz

import (
	"github.com/jinzhu/copier"
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/dao/repository"
	"github.com/kaixinhupo/quick/infrastruture/web"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)


type UserServiceImpl struct {
	engine *xorm.Engine
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository,engine *xorm.Engine) *UserServiceImpl {
	return &UserServiceImpl {
		userRepository:repo ,
		engine: engine,
	}
}

func (impl *UserServiceImpl) CreateUser(req *model.UserInfoReq) (*model.UserDetailResp, error) {
	entity := &entity.UserEntity{}
	copier.Copy(entity,req)
	var  err error
	entity,err = impl.userRepository.Insert(entity,nil)
	if err != nil {
		return nil, err
	}
	rst := &model.UserDetailResp{}
	copier.Copy(rst,entity)
	return rst, nil
}

func (impl *UserServiceImpl) Query(param *model.UserQueryReq) (*web.PageResp, error) {
	session := impl.engine.NewSession()
	defer session.Close()
 	total := impl.userRepository.Count(param,session)
	records := impl.userRepository.List(param,session)
	vos := make([]*model.UserDetailResp,len(records))
	copier.Copy(&vos,&records)
	return	&web.PageResp {
		Records: vos,
		Total: total,
		No: param.Page.No,
	}, nil
}