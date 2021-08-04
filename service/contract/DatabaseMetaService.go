package contract

import "github.com/kaixinhupo/quick/model"

type DatabaseMetaService interface {
	// ReadMeta 读取数据库元数据
	ReadMeta(tables []string) ([]*model.TableMeta, error)
}
