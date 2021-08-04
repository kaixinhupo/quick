package db

import (
	"github.com/kaixinhupo/quick/infrastructure/web"
	"xorm.io/xorm"
)

func CalcLimit(page *web.PageParam) (int, int) {
	var no int
	var size int
	if page.No < 1 {
		no = 1
	} else {
		no = page.No
	}
	if page.Size < 1 {
		size = 20
	} else {
		size = page.Size
	}
	offset := (no - 1) * size
	return size, offset
}

func OpenSession(engine *xorm.Engine, session *xorm.Session, tableAlias string) *xorm.Session {
	var s *xorm.Session
	if session != nil {
		s = session
		s.Alias(tableAlias)
	} else {
		s = engine.Alias(tableAlias)
	}
	return s
}
