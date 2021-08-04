package biz

{{#with Config}}
import (
	"github.com/jinzhu/copier"
	"{{Module}}/dao/entity"
	"{{Module}}/dao/repository"
	"{{Module}}/infrastructure/core"
	es "{{Module}}/infrastructure/errors"
	"{{Module}}/infrastructure/web"
	"{{Module}}/model"
	"xorm.io/xorm"
)

{{/with}}
{{#with Meta}}
type {{ModelName}}ServiceImpl struct {
	engine *xorm.Engine
	repo repository.{{ModelName}}Repository
}

func New{{ModelName}}Service(repo repository.{{ModelName}}Repository,engine *xorm.Engine) *{{ModelName}}ServiceImpl {
	return &{{ModelName}}ServiceImpl {
		repo: repo ,
		engine: engine,
	}
}

func (impl *{{ModelName}}ServiceImpl) Query(req *model.{{ModelName}}QueryReq) (*web.PageResp, error) {
	session := impl.engine.NewSession()
    defer func(session *xorm.Session) {
        _ = session.Close()
    }(session)
 	total := impl.repo.Count(req,session)
	records := impl.repo.List(req,session)
	vos := make([]*model.{{ModelName}}DetailResp,len(records))
	err := copier.Copy(&vos, &records)
    if err != nil {
        return nil, err
    }
	return	&web.PageResp {
		Records: vos,
		Total: total,
		No: req.Page.No,
	}, nil
}

func (impl *{{ModelName}}ServiceImpl) Item(id int64) (*model.{{ModelName}}DetailResp, error) {
	item := impl.repo.Entry(id, nil)
	if item == nil {
		return nil, es.NewBizError(core.CodeNotFound,"记录不存在")
	}
	vo := &model.{{ModelName}}DetailResp{}
	err := copier.Copy(vo, item)
    if err != nil {
        return nil, err
    }
	return vo, nil
}

func (impl *{{ModelName}}ServiceImpl) Create(req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error) {
	item := &entity.{{ModelName}}Entity{}
	err := copier.Copy(item, req)
    if err != nil {
        return nil, err
    }
	item, err = impl.repo.Insert(item,nil)
	if err != nil {
		return nil, err
	}
	rst := &model.{{ModelName}}DetailResp{}
	err = copier.Copy(rst, item)
    if err != nil {
        return nil, err
    }
	return rst, nil
}

func (impl *{{ModelName}}ServiceImpl) Update(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error){
	return impl.update(id,req,true)
}

func (impl *{{ModelName}}ServiceImpl) Patch(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error){
	return impl.update(id,req,false)
}

func (impl *{{ModelName}}ServiceImpl) update(id int64, req *model.{{ModelName}}InfoReq,all bool) (*model.{{ModelName}}DetailResp, error){
	session := impl.engine.NewSession()
    defer func(session *xorm.Session) {
        _ = session.Close()
    }(session)
	if err := session.Begin(); err != nil {
        return nil, err
    }
    item := impl.repo.Entry(id, session)
    if item == nil {
        return nil, es.NewBizError(core.CodeNotFound,"记录不存在")
    }
    err := copier.Copy(item, req)
    if err != nil {
        return nil, err
    }
    _, err = impl.repo.Update(item,all, session); if err != nil {
        return nil, err
    }
    if err := session.Commit(); err != nil {
        return nil, err
    }
    vo := &model.{{ModelName}}DetailResp{}
    err = copier.Copy(vo, item)
    if err != nil {
        return nil, err
    }
    return vo, nil
}

func (impl *{{ModelName}}ServiceImpl) Delete(id int64) error {
	return impl.repo.Delete(id, nil)
}
{{/with}}