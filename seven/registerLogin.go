package seven

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func RegisterLogin(c echo.Context) error {
	var session *mgo.Session
	var err error
	session, err = mgo.Dial(Config.DB.Host)
	defer session.Close()
	if err != nil {
		return err
	}
	username := c.FormValue("username")
	password := c.FormValue("password")
	newId := bson.NewObjectId()
	AuthenUser := AuthenUser{
		ID:       newId,
		UserName: username,
		Password: password,
	}
	err = session.DB("seven").C("authenUser").Insert(AuthenUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, AuthenUser)
}
