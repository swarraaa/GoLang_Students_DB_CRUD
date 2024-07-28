package main

import (
	"context"
	"crud_go/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded .env file")

	// Initialize the mongoClient properly
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping failed: ", err)
	}
	log.Println("Connected to MongoDB!")
}

func main() {
	// Ensure that mongoClient is properly disconnected at the end
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting from MongoDB: ", err)
		}
	}()

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	// Create student service
	studService := usecase.StudentService{MongoCollection: coll}

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/student", studService.CreateStudent).Methods(http.MethodPost)
	r.HandleFunc("/student/{id}", studService.GetStudentByID).Methods(http.MethodGet)
	r.HandleFunc("/student", studService.GetAllStudent).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", studService.UpdateStudentByID).Methods(http.MethodPut)
	r.HandleFunc("/student/{id}", studService.DeleteStudentByID).Methods(http.MethodDelete)
	r.HandleFunc("/student", studService.DeleteAllStudent).Methods(http.MethodDelete)

	log.Println("Server is running on port 4444")
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running..."))
}
