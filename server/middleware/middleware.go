package middleware

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var koleksi *mongo.Collection

func init()  {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv()  {
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("error loading di file .env")
	}
}



func GetAllTasks(w http.ResponseWriter, r *http.Request)  {
	
}

func CreateTask()  {
	
}

func TaskComplete(){

}

func UndoTask()  {
	
}

func DeleteTask()  {
	
}

func deleteAllTasks()  {
	
}