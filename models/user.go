package models

import (
	"reflect"
)

type User struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Hobby    string `json:"hobby"`
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
