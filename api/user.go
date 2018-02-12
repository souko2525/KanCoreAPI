package api

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/souko2525/KanCoreAPI/models"
)

//GetUsers return echo.HandlerFunc
func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		users := new(models.Users)
		cond := make(map[string]interface{})
		//id, _ := strconv.ParseInt(c.Param(user.PrimaryKey()), 0, 64)
		//cond["id"] = id
		if c.QueryParam("id") != "" {
			cond["id"], _ = strconv.ParseUint(c.QueryParam("id"), 10, 64)
		}
		if c.QueryParam("name") != "" {
			cond["name"] = c.QueryParam("name")
		}

		e := models.Select(users, cond, c)
		if e == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "")
		} else if e != nil {
			return c.String(http.StatusInternalServerError, e.Error())
		}
		return c.JSON(http.StatusOK, users)
	}
}

//GetUser return echo.HandlerFun
func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		user := new(models.User)
		cond := make(map[string]interface{})
		id, _ := strconv.ParseInt(c.Param(user.PrimaryKey()), 0, 64)
		cond["id"] = id
		e := models.Select(user, cond, c)
		if e == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "")
		} else if e != nil {
			return c.String(http.StatusInternalServerError, e.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}
