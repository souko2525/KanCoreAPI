package models

import (
	"reflect"
)

//User struct has User table
type User struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Hobby    string `json:"hobby"`
}

//TableName usere
func (User) TableName() string { return "user" }

//PrimaryKey id
func (User) PrimaryKey() string { return "id" }

//ReflectType user
func (User) ReflectType() reflect.Type {
	t := new(User)
	return reflect.TypeOf(*t)
}

//Users is multi user
type Users []User

//TableName user
func (Users) TableName() string { return "user" }

//PrimaryKey id
func (Users) PrimaryKey() string { return "id" }

//ReflectType user
func (Users) ReflectType() reflect.Type {
	t := new(User)
	return reflect.TypeOf(*t)
}
