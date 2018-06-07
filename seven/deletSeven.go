package seven

import (
	"net/http"

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
	err = session.DB("seven").C("seven").Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("id"))})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Delete Success")
}
