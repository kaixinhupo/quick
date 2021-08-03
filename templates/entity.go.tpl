package entity

{{#with Config}}
import (
  "{{Module}}/infrastruture/core"
)
{{/with}}

{{#with Meta}}
type {{ModelName}}Entity struct {
{{#each Fields}}
    {{Property}} {{PropertyType}} `xorm:"{{ColType}} {{#ColPk}}pk autoincr{{/ColPk}} {{^ColNull}}not{{/ColNull}} null {{#ColCreate}}created{{/ColCreate}} {{#ColUpdate}}updated{{/ColUpdate}} '{{Col}}' comment('{{ColComment}}')"`
{{/each}}
}

func(entity *{{ModelName}}Entity) TableName() string {
	return "{{TableName}}"
}
{{/with}}