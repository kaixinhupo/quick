package database

{{#with Config}}
import (
	"log"

	"{{Module}}/dao/entity"
	"{{Module}}/infrastruture/db"
	"{{Module}}/model"
	"xorm.io/xorm"
)
{{/with}}

{{#with Meta}}
type {{ModelName}}RepositoryImpl struct {
	engine *xorm.Engine
}

func New{{ModelName}}Repository(engine *xorm.Engine) *{{ModelName}}RepositoryImpl {
	return &{{ModelName}}RepositoryImpl {
		engine: engine,
	}
}

func (impl *{{ModelName}}RepositoryImpl) Insert(entity *entity.{{ModelName}}Entity, session *xorm.Session) (*entity.{{ModelName}}Entity, error) {
	_, err := impl.engine.InsertOne(entity); if err != nil {
		return nil, err
	}
	return entity,nil
}

func (impl *{{ModelName}}RepositoryImpl) Count(param *model.{{ModelName}}QueryReq, session *xorm.Session) (int64) {
	s := impl.createFilter(param, false, session)
	cnt, err := s.Count(new(entity.{{ModelName}}Entity)); if err != nil {
		log.Println("error occurred when count",err)
		return 0
	}
	return cnt
}

func (impl *{{ModelName}}RepositoryImpl) List(param *model.{{ModelName}}QueryReq, session *xorm.Session) ([]*entity.{{ModelName}}Entity) {
	s := impl.createFilter(param, false, session)
	list := make([]*entity.{{ModelName}}Entity, 0)
	err := s.Find(&list);if err != nil {
		log.Println("error occurred when find",err)
		return nil
	}
	return list
}

func (impl *{{ModelName}}RepositoryImpl) Entry(id int64, session *xorm.Session) *entity.{{ModelName}}Entity {
	s:= db.OpenSession(impl.engine,session,"t")
	entity := &entity.{{ModelName}}Entity{}
	found, err := s.ID(id).Get(entity); if err != nil {
		log.Printf("entry for key:%d not found,err:%s \n",id,err.Error())
		return nil
	}
	if !found {
		return nil
	}
	return entity
}

func (impl *{{ModelName}}RepositoryImpl) Update(entity *entity.{{ModelName}}Entity, allFields bool, session *xorm.Session) (int64, error){
	s:= db.OpenSession(impl.engine,session,"t")
	if allFields {
		return s.ID(entity.Id).AllCols().Update(entity)
	} else {
		return s.ID(entity.Id).Update(entity)
	}
}

func (impl *{{ModelName}}RepositoryImpl) Delete(id int64, session *xorm.Session) error{
	s:= db.OpenSession(impl.engine,session,"t")
	_, err := s.ID(id).Delete(new(entity.{{ModelName}}Entity))
	return err
}

func (impl *{{ModelName}}RepositoryImpl) createFilter(param *model.{{ModelName}}QueryReq, page bool,session *xorm.Session) *xorm.Session {
	s := db.OpenSession(impl.engine,session,"t")
	{{#each Queries}}
	if param.{{Property}} != {{{PropertyNullValue}}} {
		s.{{#if First}}Where{{else}}And{{/if}}("t.{{Col}}=?",param.{{Property}})
	}
	{{/each}}
	if page {
		size,offset := db.CalcLimit(&param.Page)
		s.Limit(size,offset)
		s.OrderBy("t.id desc")
	} 
	return s
}
{{/with}}