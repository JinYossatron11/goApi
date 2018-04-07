package seven

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

//----------
// Handlers
//----------

func CreateUser(c echo.Context) error {
	var session *mgo.Session
	var err error
	session, err = mgo.Dial(Config.DB.Host)
	defer session.Close()
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	score := c.FormValue("score")
	newId := bson.NewObjectId()

	seven := Seven{
		ID:   newId,
		Name: name,
		Score: score,
	}

	err = session.DB("snack").C("snacks").Insert(seven)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, seven)
}
