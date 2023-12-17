package router

import (
	"github.com/anushkarastogi/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/courses", controller.GetMyAllCourse).Methods("GET")
	router.HandleFunc("/api/course", controller.CreateModule).Methods("POST")
	router.HandleFunc("/api/course/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/course/{id}", controller.DeleteACourse).Methods("DELETE")
	router.HandleFunc("/api/deleteallcourse", controller.DeleteMovieAll).Methods("DELETE")
	return router
}
