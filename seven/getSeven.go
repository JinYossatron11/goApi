package seven

import (
	"github.com/globalsign/mgo/bson"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
)

func GetSeven(c echo.Context) error {
	seven := []Seven{}

	var session *mgo.Session
	var err error

	session, err = mgo.Dial(Config.DB.Host)
	defer session.Close()
	createBy := c.FormValue("createby")


	err = session.DB("seven").C("seven").Find(bson.M{"createby":createBy}).All(&seven)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, seven)
	
}
