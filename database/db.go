package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var cptr *gorm.DB

func Connect() error {
	if cptr != nil {
		return nil
	}
	DBMS := "mysql"
	USER := "api"
	PASS := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "kancole"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return err
	}

	db.SingularTable(true)
	cptr = db

	return nil
}

func Close() error {
	if cptr == nil {
		return nil
	}
	return cptr.Close()
}

func GetSession() *gorm.DB {
	if cptr == nil {
		if e := Connect(); e != nil {
			return nil
		}
	}
	return cptr
}

/*
func main() {
	db := gormConnect()
	enemyEx := enemy{}
	enemyEx.Id = 1
	db.First(&enemyEx)
	fmt.Println(enemyEx.Name)
}
*/
