package contract

import "github.com/kaixinhupo/quick/model"

type DatabaseMetaService interface {
	ReadMeta(tables []string) ([]*model.TableMeta, error)
}