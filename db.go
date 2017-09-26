package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type enemy struct {
	Id   int    `json:id`
	Name string `json:name`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "kancore"
	PASS := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "KanCore"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	enemyEx := enemy{}
	enemyEx.Id = 1
	db.First(&enemyEx)
	fmt.Println(enemyEx.Name)
}
