package entity

import (
	"github.com/kaixinhupo/quick/infrastructure/core"
)

type RoleEntity struct {
	Id        int64          `xorm:"bigint pk autoincr not null   'id' comment('ID')"`
	RoleName  string         `xorm:"varchar(50)  not null   'role_name' comment('角色名称')"`
	CreatedAt core.LocalTime `xorm:"datetime  not null created  'created_at' comment('创建时间')"`
	UpdatedAt core.LocalTime `xorm:"datetime  not null  updated 'updated_at' comment('修改时间')"`
}

func (entity *RoleEntity) TableName() string {
	return "role"
}
