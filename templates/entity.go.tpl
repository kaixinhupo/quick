package entity

import (
	"time"
)


type {{ModelName}}Entity struct {
{{#each Fields}}
    {{Property}} {{PropertyType}} `xorm:"{{ColType}} {{#ColPk}}pk autoincr{{/ColPk}} {{^ColNull}}not{{/ColNull}} null '{{Col}}' comment('{{ColComment}}')"`
{{/each}}
}

func(entity *{{ModelName}}Entity) TableName() string {
	return "{{TableName}}"
}