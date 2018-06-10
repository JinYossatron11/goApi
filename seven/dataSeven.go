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
	ID    bson.ObjectId  `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Price string        `bson:"price" json:"price"`
	CreateBy string      `bson:"createby" json:"createby"`
}


type AuthenUser struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	UserName   string       `bson:"username" json:"username"`
	Password   string       `bson:"password" json:"password"`
}