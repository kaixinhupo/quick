package contract

import "github.com/kaixinhupo/quick/model"

type GencodeService interface {
	// Generate ็ๆไปฃ็ 
	Generate(meta []*model.TableMeta) error
}
