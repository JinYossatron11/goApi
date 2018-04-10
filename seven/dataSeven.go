package seven

// import "github.com/globalsign/mgo/bson"

var Config = struct {
	APPName string `default:"seven"`
	HOST    string
	DB      struct {
		Host string
	}
}{}

type Seven struct {
	ID    int  `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Price string        `bson:"price" json:"price"`
}
