package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/anushkarastogi/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionStr = "mongodb+srv://anushkarastogi:admin@cluster0.vuyfibj.mongodb.net/?retryWrites=true&w=majority"

const dbName = "course"
const collectionName = "watchlist"

// important
var collection *mongo.Collection

//connect with mongodb
//init is specialised method in go which runs only at the very first time
//as the entire application executes

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionStr)

	//connect with mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection successful")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("collection instance ready")

}

//mongodb helpers-file

// insert 1 record
func insertOneRecord(coursename model.Course) {
	inserted, err := collection.InsertOne(context.Background(), coursename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 course in db with id", inserted.InsertedID)

}

// update 1 record
func updateOneRecord(courseId string) {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Modified count", result.ModifiedCount)

}
