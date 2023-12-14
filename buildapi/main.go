package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//we are creating online courses like udemy, coursera etc similar kind of backend api
//where user get feature all the courses like they create new courses, read, update
//and delete courses

// Models for course - file
type Course struct {
	CourseName  string
	CourseId    string
	CoursePrice string
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
	return c.CourseId == "" && c.CourseName == ""
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
