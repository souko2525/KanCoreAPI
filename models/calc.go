package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}

/*
type enemy struct {
	Name string
	HP   int
	A    int
	T    int
	D    int
}

type ally struct {
	Name string
	HP   int
	A    int
	T    int
	D    int
}

type equip struct {
	A int
	H int
	D int
}

type enemy struct {
    Id   int    `json:id`
    Name string `json:name`
}

*/
