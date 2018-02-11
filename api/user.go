package api

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/souko2525/KanCoreAPI/models"
	"net/http"
	"strconv"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		user := new(models.User)
		id, _ := strconv.ParseInt(c.Param(user.PrimaryKey()), 0, 64)
		e := models.Select(user, id, c)
		if e == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "")
		} else if e != nil {
			return c.String(http.StatusInternalServerError, e.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}
