package models

import (
	//"fmt"
	"reflect"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	db "github.com/souko2525/KanCoreAPI/database"
	//"strings"
)

//DatabaseTable shold has TableName, PrimaryKey, ReflectType
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

//Select return struct t
func Select(t DatabaseTable, cond map[string]interface{}, cont echo.Context) error {
	s := db.GetSession()
	stmt := s.Table(t.TableName()).Select(colums(t)).Where(cond)
	if cont.QueryParam("offset") != "" {
		v, e := strconv.ParseUint(cont.QueryParam("offset"), 10, 64)
		if e != nil {
			return e
		}
		stmt.Offset(v)
	}

	if cont.QueryParam("limit") != "" {
		v, e := strconv.ParseUint(cont.QueryParam("limit"), 10, 64)
		if e != nil {
			return e
		}
		stmt.Limit(v)
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

//Insert user
func Insert(t DatabaseTable) error {
	/*
		cs := make([]string, 0, len(d))
		vs := make([]interface{}, 0, len(d))
		for k, v := range d {
			cs = append(cs, k)
			vs = append(vs, v)
		}
	*/
	s := db.GetSession()
	tx := s.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Table(t.TableName()).Create(t).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
