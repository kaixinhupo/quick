package entity

{{#with Config}}
import (
  "{{Module}}/infrastructure/core"
)
{{/with}}

{{#with Meta}}
type {{ModelName}}Entity struct {
{{#each Fields}}
    {{Property}} {{PropertyType}} `xorm:"{{ColTypeLower}}{{ColTypeLen}} {{#ColPk}}pk autoincr{{/ColPk}} {{^ColNull}}not{{/ColNull}} null {{#ColCreate}}created{{/ColCreate}} {{#ColUpdate}}updated{{/ColUpdate}} '{{Col}}' comment('{{ColComment}}')"`
{{/each}}
}

func(entity *{{ModelName}}Entity) TableName() string {
	return "{{TableName}}"
}
{{/with}}