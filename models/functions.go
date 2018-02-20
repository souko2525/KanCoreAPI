package models

import (
	"encoding/json"
	"github.com/labstack/echo"
	"reflect"
	"time"
)

//ParseJSON return map
func ParseJSON(c echo.Context) (map[string]interface{}, error) {
	var jd map[string]interface{}
	err := json.NewDecoder(c.Request().Body).Decode(&jd)
	if err != nil {
		return nil, err
	}
	return jd, nil
}

//column_value return db value
func columnValue(i interface{}) interface{} {
	var v interface{}

	return v
}

// BindJSON map d to m
func BindJSON(m interface{}, d map[string]interface{}) error {
	rt := reflect.TypeOf(m)
	rv := reflect.ValueOf(m)
	for i := 0; i < rt.NumField(); i++ {
		if v := columnValue(rv.Field(i).Interface()); v != nil {
			continue
		}
		f := rt.Field(i)
		name := f.Tag.Get("json")
		dval, found := f.Tag.Lookup("default")
		if name == "" || !found {
			continue
		}
		if dval == "NOW()" {
			dval = time.Now().Format("2006-01-02T03:04:05Z")
		}
		if _, exists := d[name]; !exists {
			d[name] = dval
		}
	}
	return nil

}
