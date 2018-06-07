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
var numberId int = 1

func CreateUser(c echo.Context) error {
	var session *mgo.Session
	var err error
	session, err = mgo.Dial(Config.DB.Host)
	defer session.Close()
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	price := c.FormValue("price")
	newId := bson.NewObjectId()
	seven := Seven{
		ID: newId,
		Name:  name,
		Price: price,
	}

	err = session.DB("seven").C("seven").Insert(seven)
	if err != nil {
		return err
	}
	numberId++
	return c.JSON(http.StatusOK, seven)
}
