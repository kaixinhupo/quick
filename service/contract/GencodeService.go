package contract

import "github.com/kaixinhupo/quick/model"

type GencodeService interface {
	Generate(meta []*model.TableMeta) error
}