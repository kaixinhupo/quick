package entity

import (
	"time"
)


type UserEntity struct {
	Id        int64 `xorm:"bigint pk autoincr 'id' comment('ID')"`
	Username  string `xorm:"varchar(20) notnull unique 'username' comment('用户名')"`
	Password  string `xorm:"varchar(225) notnull  'password' comment('密码')"`
	CreatedAt time.Time `xorm:"datetime notnull created 'created_at' comment('创建时间')"`
	UpdatedAt time.Time `xorm:"datetime notnull updated 'updated_at' comment('修改时间')"`
}

func(entity *UserEntity) TableName() string {
	return "user"
}