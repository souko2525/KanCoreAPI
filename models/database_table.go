package models

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	db "github.com/souko2525/KanCoreAPI/database"
	"reflect"
	"strconv"
	//"strings"
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

func Select(t DatabaseTable, cond map[string]interface{}, cont echo.Context) error {
	s := db.GetSession()
	stmt := s.Table(t.TableName()).Select(colums(t)).Where(cond)
	if cont.QueryParam("offset") != "" {
		if v, e := strconv.ParseUint(cont.QueryParam("offset"), 10, 64); e != nil {
			return e
		} else {
			stmt.Offset(v)
		}

	}

	if cont.QueryParam("limit") != "" {
		if v, e := strconv.ParseUint(cont.QueryParam("limit"), 10, 64); e != nil {
			return e
		} else {
			stmt.Limit(v)
		}

	}

	if cont.QueryParam("order_by") != "" {
		stmt.Order(cont.QueryParam("order_by"))
	}

	result := stmt.Find(t)

	if err := result.Error; err != nil {
		return err
	} else if result.RecordNotFound() {
		return gorm.ErrRecordNotFound
	}

	return nil
}
