package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string
	Password string
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(user, fields...); err != nil {
		return err
	}
	return nil
}
