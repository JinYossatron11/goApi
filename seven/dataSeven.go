package seven

import "github.com/globalsign/mgo/bson"

var Config = struct {
	APPName string `default:"seven"`
	HOST    string
	DB      struct {
		Host string
	}

}{}

type Seven struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
	Score string        `bson:"score" json:"score"`
}
