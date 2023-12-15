package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//we are creating online courses like udemy, coursera etc similar kind of backend api
//where user get feature all the courses like they create new courses, read, update
//and delete courses

// Models for course - file
type Course struct {
	CourseName  string
	CourseId    string
	CoursePrice int
	Author      *Author
}

type Author struct {
	Fullname string
	Website  string
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to api"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

// how to grabsingle course based on request id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	//whenever user require 1 course, means they are gonna provide me unique id
	//then we'll compare that id with course id which is in form of arrays
	//if found then we will return it

	fmt.Println("Get one course")
	w.Header().Set("Content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//1-loop through course, 2- find matching id, 3- return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}

	}
}

// create course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id, string
	//add or append new course to courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// update course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//1-loop through value then we will get id. 2- remove that id and then add with my id(we will get by params)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}
