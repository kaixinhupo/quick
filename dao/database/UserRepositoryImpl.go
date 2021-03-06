package database

import (
	"log"

	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/infrastructure/db"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)

type UserRepositoryImpl struct {
	engine *xorm.Engine
}

func NewUserRepository(engine *xorm.Engine) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		engine: engine,
	}
}

func (impl *UserRepositoryImpl) Insert(entity *entity.UserEntity, session *xorm.Session) (*entity.UserEntity, error) {
	s := db.OpenSession(impl.engine, session, "t")
	_, err := s.InsertOne(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (impl *UserRepositoryImpl) Count(param *model.UserQueryReq, session *xorm.Session) int64 {
	s := impl.createFilter(param, false, session)
	cnt, err := s.Count(new(entity.UserEntity))
	if err != nil {
		log.Println("error occurred when count", err)
		return 0
	}
	return cnt
}

func (impl *UserRepositoryImpl) List(param *model.UserQueryReq, session *xorm.Session) []*entity.UserEntity {
	s := impl.createFilter(param, false, session)
	list := make([]*entity.UserEntity, 0)
	err := s.Find(&list)
	if err != nil {
		log.Println("error occurred when find", err)
		return nil
	}
	return list
}

func (impl *UserRepositoryImpl) Entry(id int64, session *xorm.Session) *entity.UserEntity {
	s := db.OpenSession(impl.engine, session, "t")
	item := &entity.UserEntity{}
	found, err := s.ID(id).Get(item)
	if err != nil {
		log.Printf("entry for key:%d not found,err:%s \n", id, err.Error())
		return nil
	}
	if !found {
		return nil
	}
	return item
}

func (impl *UserRepositoryImpl) Update(entity *entity.UserEntity, allFields bool, session *xorm.Session) (int64, error) {
	s := db.OpenSession(impl.engine, session, "t")
	if allFields {
		return s.ID(entity.Id).AllCols().Update(entity)
	} else {
		return s.ID(entity.Id).Update(entity)
	}
}

func (impl *UserRepositoryImpl) Delete(id int64, session *xorm.Session) error {
	s := db.OpenSession(impl.engine, session, "t")
	_, err := s.ID(id).Delete(new(entity.UserEntity))
	return err
}

func (impl *UserRepositoryImpl) createFilter(param *model.UserQueryReq, page bool, session *xorm.Session) *xorm.Session {
	s := db.OpenSession(impl.engine, session, "t")
	if param.Username != "" {
		s.Where("t.username=?", param.Username)
	}
	if page {
		size, offset := db.CalcLimit(&param.Page)
		s.Limit(size, offset)
		s.OrderBy("t.id desc")
	}
	return s
}
