package database

import (
	"log"

	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/infrastructure/db"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)

type RoleRepositoryImpl struct {
	engine *xorm.Engine
}

func NewRoleRepository(engine *xorm.Engine) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{
		engine: engine,
	}
}

func (impl *RoleRepositoryImpl) Insert(entity *entity.RoleEntity, session *xorm.Session) (*entity.RoleEntity, error) {
	s := db.OpenSession(impl.engine, session, "t")
	_, err := s.InsertOne(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (impl *RoleRepositoryImpl) Count(param *model.RoleQueryReq, session *xorm.Session) int64 {
	s := impl.createFilter(param, false, session)
	cnt, err := s.Count(new(entity.RoleEntity))
	if err != nil {
		log.Println("error occurred when count", err)
		return 0
	}
	return cnt
}

func (impl *RoleRepositoryImpl) List(param *model.RoleQueryReq, session *xorm.Session) []*entity.RoleEntity {
	s := impl.createFilter(param, false, session)
	list := make([]*entity.RoleEntity, 0)
	err := s.Find(&list)
	if err != nil {
		log.Println("error occurred when find", err)
		return nil
	}
	return list
}

func (impl *RoleRepositoryImpl) Entry(id int64, session *xorm.Session) *entity.RoleEntity {
	s := db.OpenSession(impl.engine, session, "t")
	entity := &entity.RoleEntity{}
	found, err := s.ID(id).Get(entity)
	if err != nil {
		log.Printf("entry for key:%d not found,err:%s \n", id, err.Error())
		return nil
	}
	if !found {
		return nil
	}
	return entity
}

func (impl *RoleRepositoryImpl) Update(entity *entity.RoleEntity, allFields bool, session *xorm.Session) (int64, error) {
	s := db.OpenSession(impl.engine, session, "t")
	if allFields {
		return s.ID(entity.Id).AllCols().Update(entity)
	} else {
		return s.ID(entity.Id).Update(entity)
	}
}

func (impl *RoleRepositoryImpl) Delete(id int64, session *xorm.Session) error {
	s := db.OpenSession(impl.engine, session, "t")
	_, err := s.ID(id).Delete(new(entity.RoleEntity))
	return err
}

func (impl *RoleRepositoryImpl) createFilter(param *model.RoleQueryReq, page bool, session *xorm.Session) *xorm.Session {
	s := db.OpenSession(impl.engine, session, "t")
	if param.RoleName != "" {
		s.Where("t.role_name=?", param.RoleName)
	}
	if page {
		size, offset := db.CalcLimit(&param.Page)
		s.Limit(size, offset)
		s.OrderBy("t.id desc")
	}
	return s
}
