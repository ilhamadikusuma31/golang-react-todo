package router

import (
	"GOLANG-REACT-TODO/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("api/task", middleware.GetAllTasks).Methods("GET","OPTIONS")
	router.HandleFunc("api/tasks",middleware.CreateTask).Methods("POST","OPTIONS")
	router.HandleFunc("api/tasks/{id}",middleware.TaskComplete).Methods("PUT","OPTIONS")
	router.HandleFunc("api/undoTasks/{id}",middleware.UndoTask).Methods("PUT","OPTIONS")
	router.HandleFunc("api/deleteTask/{id}",middleware.DeleteTask).Methods("DELETE","OPTIONS")
	router.HandleFunc("api/deleteAllTasks/{id}",middleware.DeleteAllTask).Methods("DELETE","OPTIONS")
}