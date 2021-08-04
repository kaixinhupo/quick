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

type RoleServiceImpl struct {
	engine *xorm.Engine
	repo   repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository, engine *xorm.Engine) *RoleServiceImpl {
	return &RoleServiceImpl{
		repo:   repo,
		engine: engine,
	}
}

func (impl *RoleServiceImpl) Query(req *model.RoleQueryReq) (*web.PageResp, error) {
	session := impl.engine.NewSession()
	defer func(session *xorm.Session) {
		_ = session.Close()
	}(session)
	total := impl.repo.Count(req, session)
	records := impl.repo.List(req, session)
	vos := make([]*model.RoleDetailResp, len(records))
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

func (impl *RoleServiceImpl) Item(id int64) (*model.RoleDetailResp, error) {
	item := impl.repo.Entry(id, nil)
	if item == nil {
		return nil, es.NewBizError(core.CodeNotFound, "记录不存在")
	}
	vo := &model.RoleDetailResp{}
	err := copier.Copy(vo, item)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (impl *RoleServiceImpl) Create(req *model.RoleInfoReq) (*model.RoleDetailResp, error) {
	item := &entity.RoleEntity{}
	err := copier.Copy(item, req)
	if err != nil {
		return nil, err
	}
	item, err = impl.repo.Insert(item, nil)
	if err != nil {
		return nil, err
	}
	rst := &model.RoleDetailResp{}
	err = copier.Copy(rst, item)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (impl *RoleServiceImpl) Update(id int64, req *model.RoleInfoReq) (*model.RoleDetailResp, error) {
	return impl.update(id, req, true)
}

func (impl *RoleServiceImpl) Patch(id int64, req *model.RoleInfoReq) (*model.RoleDetailResp, error) {
	return impl.update(id, req, false)
}

func (impl *RoleServiceImpl) update(id int64, req *model.RoleInfoReq, all bool) (*model.RoleDetailResp, error) {
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
	vo := &model.RoleDetailResp{}
	err = copier.Copy(vo, item)
	if err != nil {
		return nil, err
	}
	return vo, nil
}

func (impl *RoleServiceImpl) Delete(id int64) error {
	return impl.repo.Delete(id, nil)
}
