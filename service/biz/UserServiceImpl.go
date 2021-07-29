package biz

import (
	"github.com/jinzhu/copier"
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/dao/repository"
	"github.com/kaixinhupo/quick/infrastruture/core"
	es "github.com/kaixinhupo/quick/infrastruture/errors"
	"github.com/kaixinhupo/quick/infrastruture/web"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)


type UserServiceImpl struct {
	engine *xorm.Engine
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository,engine *xorm.Engine) *UserServiceImpl {
	return &UserServiceImpl {
		repo:repo ,
		engine: engine,
	}
}

func (impl *UserServiceImpl) Query(req *model.UserQueryReq) (*web.PageResp, error) {
	session := impl.engine.NewSession()
	defer session.Close()
 	total := impl.repo.Count(req,session)
	records := impl.repo.List(req,session)
	vos := make([]*model.UserDetailResp,len(records))
	copier.Copy(&vos,&records)
	return	&web.PageResp {
		Records: vos,
		Total: total,
		No: req.Page.No,
	}, nil
}
// 查询单条记录
func (impl *UserServiceImpl) Item(id int64) (*model.UserDetailResp, error) {
	entity := impl.repo.Entry(id, nil)
	if entity == nil {
		return nil, es.NewBizError(core.CodeNotFound,"记录不存在")
	}
	vo := &model.UserDetailResp{}
	copier.Copy(vo,entity)
	return vo, nil
}

func (impl *UserServiceImpl) Create(req *model.UserInfoReq) (*model.UserDetailResp, error) {
	entity := &entity.UserEntity{}
	copier.Copy(entity,req)
	var  err error
	entity, err = impl.repo.Insert(entity,nil)
	if err != nil {
		return nil, err
	}
	rst := &model.UserDetailResp{}
	copier.Copy(rst,entity)
	return rst, nil
}

// 更新记录
func (impl *UserServiceImpl) Update(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error){
	return impl.update(id,req,true)
}

// 修改记录
func (impl *UserServiceImpl) Patch(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error){
	return impl.update(id,req,false)
}

// 更新记录
func (impl *UserServiceImpl) update(id int64, req *model.UserInfoReq,all bool) (*model.UserDetailResp, error){
	session := impl.engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
        return nil, err
    }
	entity := impl.repo.Entry(id, session)
	if entity == nil {
		return nil, es.NewBizError(core.CodeNotFound,"记录不存在")
	}
	copier.Copy(entity, req) 
	_, err := impl.repo.Update(entity,all, session); if err != nil {
		return nil, err
	}
	if err := session.Commit(); err != nil {
        return nil, err
    }
	vo := & model.UserDetailResp{}
	copier.Copy(vo,entity) 
	return vo, nil
}

// 删除记录
func (impl *UserServiceImpl) Delete(id int64) error {
	return impl.repo.Delete(id, nil)
}

