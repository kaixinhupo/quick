package database

import (
	"log"

	"github.com/kaixinhupo/quick/dao/entity"
	"github.com/kaixinhupo/quick/infrastruture/db"
	"github.com/kaixinhupo/quick/model"
	"xorm.io/xorm"
)


type UserRepositoryImpl struct {
	engine *xorm.Engine
}

func NewUserRepository(engine *xorm.Engine) *UserRepositoryImpl {
	return &UserRepositoryImpl {
		engine: engine,
	}
}

func (impl *UserRepositoryImpl) Insert(user *entity.UserEntity, session *xorm.Session) (*entity.UserEntity, error) {
	_, err := impl.engine.InsertOne(user); if err != nil {
		return nil, err
	}
	return user,nil
}

func (impl *UserRepositoryImpl) Count(param *model.UserQueryReq, session *xorm.Session) (int64) {
	s := impl.createFilter(param, false, session)
	cnt, err := s.Count(new(entity.UserEntity)); if err != nil {
		log.Println("error occurred when count",err)
		return 0
	}
	return cnt
}

func (impl *UserRepositoryImpl) List(param *model.UserQueryReq, session *xorm.Session) ([]*entity.UserEntity) {
	s := impl.createFilter(param, false, session)
	list := make([]*entity.UserEntity, 0)
	err := s.Find(&list);if err != nil {
		log.Println("error occurred when find",err)
		return nil
	}
	return list
}

func (impl *UserRepositoryImpl) createFilter(param *model.UserQueryReq, page bool,session *xorm.Session) *xorm.Session {
	s:=db.OpenSession(impl.engine,session,"t")
	if param.Username !="" {
		s.Where("t.username=?",param.Username)
	}
	if page {
		size,offset := db.CalcLimit(&param.Page)
		s.Limit(size,offset)
		s.OrderBy("t.id desc")
	} 
	return s
}