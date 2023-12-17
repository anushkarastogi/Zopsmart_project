package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anushkarastogi/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// insert 1 course
func insertOneCourse(coursename model.Course) {
	inserted, err := collection.InsertOne(context.Background(), coursename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 course in db with id", inserted.InsertedID)

}

// update 1 Course
func updateOneCourse(courseId string) {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Modified count", result.ModifiedCount)

}

// delete 1 Course
func deleteOneCourse(courseId string) {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie with id is deleted", deleteCount)

}

// delete all Courses
func deleteAllCourse() int64 {
	deleteRes, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {

	}
	fmt.Println("No of movies deleted", deleteRes.DeletedCount)
	return deleteRes.DeletedCount
}

// get all Courses from database
func getAllCourses() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var courses []primitive.M

	for cur.Next(context.Background()) {
		var module bson.M
		err = cur.Decode(&module)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, module)
	}
	defer cur.Close(context.Background())
	return courses
}

// actual controller - file
func GetMyAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	allCourses := getAllCourses()
	json.NewEncoder(w).Encode(allCourses)
}

func CreateModule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var module model.Course
	_ = json.NewDecoder(r.Body).Decode(&module)
	insertOneCourse(module)
	json.NewEncoder(w).Encode(module)
}

// module to mark as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneCourse(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneCourse(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovieAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllCourse()
	json.NewEncoder(w).Encode(count)

}
