package biz

import (
	"github.com/jinzhu/copier"
	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/dao/repository"
	"github.com/kaixinhupo/quick/infrastructure/core"
	es "github.com/kaixinhupo/quick/infrastructure/errors"
	"github.com/kaixinhupo/quick/infrastructure/web"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)

type UserServiceImpl struct {
	engine *xorm.Engine
	repo   repository.UserRepository
}

func NewUserService(repo repository.UserRepository, engine *xorm.Engine) *UserServiceImpl {
	return &UserServiceImpl{
		repo:   repo,
		engine: engine,
	}
}

func (impl *UserServiceImpl) Query(req *model.UserQueryReq) (*web.PageResp, error) {
	session := impl.engine.NewSession()
	defer func(session *xorm.Session) {
		_ = session.Close()
	}(session)
	total := impl.repo.Count(req, session)
	records := impl.repo.List(req, session)
	vos := make([]*model.UserDetailResp, len(records))
	err := copier.Copy(&vos, &records)
	if err != nil {
		return nil, err
	}
	return &web.PageResp{
		Records: vos,
		Total:   total,
		No:      req.Page.No,
	}, nil
}

func (impl *UserServiceImpl) Item(id int64) (*model.UserDetailResp, error) {
	item := impl.repo.Entry(id, nil)
	if item == nil {
		return nil, es.NewBizError(core.CodeNotFound, "记录不存在")
	}
	vo := &model.UserDetailResp{}
	err := copier.Copy(vo, item)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (impl *UserServiceImpl) Create(req *model.UserInfoReq) (*model.UserDetailResp, error) {
	item := &entity.UserEntity{}
	err := copier.Copy(item, req)
	if err != nil {
		return nil, err
	}
	item, err = impl.repo.Insert(item, nil)
	if err != nil {
		return nil, err
	}
	rst := &model.UserDetailResp{}
	err = copier.Copy(rst, item)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (impl *UserServiceImpl) Update(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error) {
	return impl.update(id, req, true)
}

func (impl *UserServiceImpl) Patch(id int64, req *model.UserInfoReq) (*model.UserDetailResp, error) {
	return impl.update(id, req, false)
}

func (impl *UserServiceImpl) update(id int64, req *model.UserInfoReq, all bool) (*model.UserDetailResp, error) {
	session := impl.engine.NewSession()
	defer func(session *xorm.Session) {
		_ = session.Close()
	}(session)
	if err := session.Begin(); err != nil {
		return nil, err
	}
	item := impl.repo.Entry(id, session)
	if item == nil {
		return nil, es.NewBizError(core.CodeNotFound, "记录不存在")
	}
	err := copier.Copy(item, req)
	if err != nil {
		return nil, err
	}
	_, err = impl.repo.Update(item, all, session)
	if err != nil {
		return nil, err
	}
	if err := session.Commit(); err != nil {
		return nil, err
	}
	vo := &model.UserDetailResp{}
	err = copier.Copy(vo, item)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (impl *UserServiceImpl) Delete(id int64) error {
	return impl.repo.Delete(id, nil)
}
