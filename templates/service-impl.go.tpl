package biz

{{#with Config}}
import (
	"github.com/jinzhu/copier"
	"{{Module}}/dao/entity"
	"{{Module}}/dao/repository"
	"{{Module}}/infrastruture/core"
	es "{{Module}}/infrastruture/errors"
	"{{Module}}/infrastruture/web"
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
	defer session.Close()
 	total := impl.repo.Count(req,session)
	records := impl.repo.List(req,session)
	vos := make([]*model.{{ModelName}}DetailResp,len(records))
	copier.Copy(&vos,&records)
	return	&web.PageResp {
		Records: vos,
		Total: total,
		No: req.Page.No,
	}, nil
}

// 查询单条记录
func (impl *{{ModelName}}ServiceImpl) Item(id int64) (*model.{{ModelName}}DetailResp, error) {
	entity := impl.repo.Entry(id, nil)
	if entity == nil {
		return nil, es.NewBizError(core.CodeNotFound,"记录不存在")
	}
	vo := &model.{{ModelName}}DetailResp{}
	copier.Copy(vo,entity)
	return vo, nil
}

func (impl *{{ModelName}}ServiceImpl) Create(req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error) {
	entity := &entity.{{ModelName}}Entity{}
	copier.Copy(entity,req)
	var  err error
	entity, err = impl.repo.Insert(entity,nil)
	if err != nil {
		return nil, err
	}
	rst := &model.{{ModelName}}DetailResp{}
	copier.Copy(rst,entity)
	return rst, nil
}

// 更新记录
func (impl *{{ModelName}}ServiceImpl) Update(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error){
	return impl.update(id,req,true)
}

// 修改记录
func (impl *{{ModelName}}ServiceImpl) Patch(id int64, req *model.{{ModelName}}InfoReq) (*model.{{ModelName}}DetailResp, error){
	return impl.update(id,req,false)
}

// 更新记录
func (impl *{{ModelName}}ServiceImpl) update(id int64, req *model.{{ModelName}}InfoReq,all bool) (*model.{{ModelName}}DetailResp, error){
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
	vo := &model.{{ModelName}}DetailResp{}
	copier.Copy(vo,entity) 
	return vo, nil
}

// 删除记录
func (impl *{{ModelName}}ServiceImpl) Delete(id int64) error {
	return impl.repo.Delete(id, nil)
}
{{/with}}