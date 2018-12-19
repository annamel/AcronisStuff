package swagger

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"strconv"
	"strings"
)

var (
	LOGS    = "logs"
	METRICS = "metrics"
)

type File struct {
	Id     bson.ObjectId `bson:"_id"`
	UserId int           `bson:"user_id"`
	AppId  string        `bson:"app_id"`
	Path   string        `bson:"path"`
}

//func setupDB(fileType string) (collection *mgo.Collection){
//	session, errDB := mgo.Dial("mongodb://127.0.0.1")
//
//	if errDB != nil {
//		panic(errDB)
//	}
//
//	collection = session.DB("acronisdb").C(fileType)
//	return
//}

func deleteAll(fileType string) {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	_, err := collection.RemoveAll(bson.M{})

	if err != nil {
		panic(err)
	}
}

func deleteFromDB(id string, fileType string) {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	_, err := collection.RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		panic(err)
	}
}

func getAll(fileType string) (ids string) {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	var items []File

	collection.Find(bson.M{}).All(&items)

	for _, element := range items {
		ids += element.Id.Hex() + "\n"
	}
	return
}

func get(id string, fileType string) File {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	//collection := setupDB(fileType)

	item := File{}
	collection.FindId(bson.ObjectIdHex(id)).One(&item)

	return item

}

func put(text string, fileType string) {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	id := bson.NewObjectId()

	path := ""
	user := 0
	app := ""

	//TODO: Make it relative
	if fileType == LOGS {
		path = "/home/files/logs/" + id.Hex() + ".log"
	} else {
		path = "/home/files/metrics/" + id.Hex() + ".json"
	}

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	texts := strings.Split(text, "\n")

	text = ""

	for index, element := range texts{
		if index == 0 {
			user, _ = strconv.Atoi(element)
		}
		if index == 1 {
			app = element
		}
		text += element + "\n"
	}

	f.WriteString(text)

	defer f.Close()

	collection.Insert(&File{Id: id, UserId: user, AppId: app, Path: path})
}
