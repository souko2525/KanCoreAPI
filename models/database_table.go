package models

import (
	//"fmt"
	"reflect"
	//"strconv"
	//"strings"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	db "github.com/souko2525/KanCoreAPI/database"
)

type DatabaseTable interface {
	TableName() string
	PrimaryKey() string
	ReflectType() reflect.Type
}

func colums(t DatabaseTable) []string {
	ref := t.ReflectType()
	cols := []string{}
	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i)

		col := f.Tag.Get("json")
		if col == "" {
			continue
		}
		cols = append(cols, col)
	}
	return cols
}

func Select(t DatabaseTable, id int64, cont echo.Context) error {
	//cols := "`" + strings.Join(colums(t), "`, `") + "`"
	s := db.GetSession()
	//result := s.First(t, id)
	result := s.Table(t.TableName()).Select(colums(t)).First(t, id)
	if err := result.Error; err != nil {
		return err
	} else if result.RecordNotFound() {
		return gorm.ErrRecordNotFound
	}

	return nil
}
