package models

import (
	"reflect"
)

type User struct {
	Id       int64
	Name     string `sql:"size:60"`
	Password string `sql:"size:60"`
	Hobby    string `sql:"size:60"`
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
