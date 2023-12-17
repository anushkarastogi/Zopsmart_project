package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	CourseID   primitive.ObjectID `json: "_id,omitempty" bson: "_id,omitempty"`
	CourseName string             `json: "coursename,omitempty"`
	Verified   bool               `json: "watched,omitempty"`
}
