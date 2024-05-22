package module

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/RuthDianaPurnamasari/package_ats/model"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (*mongo.Database) {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
    if err != nil {
         fmt.Errorf("MongoConnect: %v", err)
    }
    return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
    insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
    if err != nil {
        fmt.Printf("InsertOneDoc: %v\n", err)
    }
    return insertResult.InsertedID

}

func InsertPenelitian(judul string, institusi string, penulis string, datetime primitive.DateTime, ringkasan string, biodata model.Peneliti) (InsertedID interface{}) {
    penelitian := map[string]interface{}{
        "_id":       primitive.NewObjectID(),
        "judul":     judul,
        "institusi": institusi,
        "penulis":   penulis,
        "datetime":  datetime,
        "ringkasan": ringkasan,
        "biodata":   biodata,
    }
    return InsertOneDoc("ATS", "penelitian", penelitian)
}

func GetAllPenelitian() (data []model.Penelitian) {
	collection := MongoConnect("ATS").Collection("penelitian")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
