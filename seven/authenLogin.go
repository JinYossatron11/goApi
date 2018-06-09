package seven

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func AuthenLogin(c echo.Context) error {
	var session *mgo.Session
	var err error
	session, err = mgo.Dial(Config.DB.Host)
	defer session.Close()
	username := c.FormValue("username")
	password := c.FormValue("password")



	result := AuthenUser{}
	err = session.DB("seven").C("authenUser").Find(bson.M{"username":username,"password":password}).One(&result)
	if err != nil {
		if err.Error() != "not found" {
			return err
		}
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No Find Id"})
	}
	return c.JSON(http.StatusOK, result)
}
