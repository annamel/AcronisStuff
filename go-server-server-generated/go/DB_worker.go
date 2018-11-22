package swagger

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"strings"
)

var (
	LOGS    = "logs"
	METRICS = "metrics"
)

type File struct {
	Id   bson.ObjectId `bson:"_id"`
	Path string        `bson:"path"`
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

	var items []File

	collection.Find(bson.M{}).All(&items)

	for _, element := range items {
		deleteFromDB(element.Id.Hex(), fileType)
	}

	//if err != nil {
	//	panic(err)
	//}
}

func deleteFromDB(id string, fileType string) {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	item := File{}
	collection.FindId(bson.ObjectIdHex(id)).One(&item)

	path := item.Path

	err := os.Remove(path)
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

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

func get(id string, fileType string) string {
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	//collection := setupDB(fileType)

	item := File{}
	collection.FindId(bson.ObjectIdHex(id)).One(&item)

	path := item.Path

	text, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(text)

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
		if index == 0 || index == 1 || index == 2 || index == len(texts) - 2 {
			continue
		}
		text += element + "\n"
	}

	f.WriteString(text)

	defer f.Close()

	collection.Insert(&File{Id: id, Path: path})
}
