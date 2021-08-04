package contract

import "github.com/kaixinhupo/quick/model"

type GencodeService interface {
	// Generate 生成代码
	Generate(meta []*model.TableMeta) error
}
