package swagger

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	LOGS    = "logs"
	METRICS = "metrics"
)

type File struct {
	Id   bson.ObjectId `bson:"_id"`
	Path string `bson:"path"`
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

func deleteAll(fileType string){
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	_, err := collection.RemoveAll(bson.M{})

	if err != nil{
		panic(err)
	}
}

func deleteFromDB(id string, fileType string){
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)

	_, err := collection.RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil{
		panic(err)
	}
}

func getAll(fileType string) (ids string){
	session, errDB := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	if errDB != nil {
		panic(errDB)
	}

	collection := session.DB("acronisdb").C(fileType)
	
	items := []File{}
	
	collection.Find(bson.M{}).All(&items)

	for _, element := range items {
		ids += strconv.Itoa(int(element.Id.Counter())) + "\n"
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

	if err != nil{
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
		path = "\\home\\files\\logs\\" + strconv.Itoa(int(id.Counter())) + ".log"
	} else {
		path = "\\home\\files\\metrics\\" + strconv.Itoa(int(id.Counter())) + ".json"
	}

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	f.WriteString(text)

	defer f.Close()

	collection.Insert(&File{Id: id, Path: path})
}
