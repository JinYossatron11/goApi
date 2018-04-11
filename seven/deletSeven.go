package seven

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func DeleteId(c echo.Context) error {

	var session *mgo.Session
	var err error
	session, err = mgo.Dial(Config.DB.Host)
	defer session.Close()
	if err != nil {
		return err
	}
	newId, _ := strconv.Atoi(c.Param("id"))

	err = session.DB("snack").C("snacks").Remove(bson.M{"_id": newId})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Delete Success")
}
