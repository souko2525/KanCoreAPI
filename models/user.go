package models

import (
	"reflect"
)

type User struct {
	Id       int64  `db:"id" json:"id"`
	Name     string `sql:"size:60" db:"name" json:"name"`
	Password string `sql:"size:60" db:"password" json:"password"`
	Hobby    string `sql:"size:60" db:"hobby" json:"hobby"`
}

func (User) TableName() string  { return "user" }
func (User) PrimaryKey() string { return "id" }
func (User) ReflectType() reflect.Type {
	t := new(User)
	return reflect.TypeOf(*t)
}

type Users []User

func (Users) TableName() string  { return "user" }
func (Users) PrimaryKey() string { return "id" }
func (Users) ReflectType() reflect.Type {
	t := new(User)
	return reflect.TypeOf(*t)
}
