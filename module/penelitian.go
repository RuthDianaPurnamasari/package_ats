package package_ats

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
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
func InsertPenelitian(judul string, institusi string, penulis string, datetime primitive.DateTime, ringkasan string, biodata map[string]interface{}) (insertedID interface{}) {
	var penelitian = map[string]interface{}{
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

func GetAllPenelitian() []interface{} {
	var penelitians []interface{}
	collection := MongoConnect("ATS").Collection("penelitian")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPenelitian:", err)
		return penelitians
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var penelitian interface{}
		err := cursor.Decode(&penelitian)
		if err != nil {
			fmt.Printf("Error decoding penelitian: %v\n", err)
			continue
		}
		penelitians = append(penelitians, penelitian)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("Cursor error: %v\n", err)
	}
	return penelitians
}
