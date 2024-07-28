package repository

import (
	"context"
	"crud_go/model"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded .env file")
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Error while connecting to MongoDB: ", err)
	}
	log.Println("Connected to MongoDB")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error while pinging to MongoDB: ", err)
	}
	log.Println("Ping to MongoDB successful")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	client := newMongoClient()
	defer client.Disconnect(context.Background())

	//dummy data
	stud1 := uuid.New().String()
	stud2 := uuid.New().String()

	// connect to the database
	db := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	studRepo := StudentRepo{MongoCollection: db}

	// insert a student
	t.Run("InsertStudent1", func(t *testing.T) {
		stud := model.Student{
			StudentID:  stud1,
			Name:       "John Doe",
			Department: "CSE",
		}
		result, err := studRepo.InsertStudent(&stud)
		if err != nil {
			t.Fatal("Insert 1 operation failed", err)
		}
		t.Log("Insert 1 successful", result)
	})
	t.Run("InsertStudent2", func(t *testing.T) {
		stud := model.Student{
			StudentID:  stud2,
			Name:       "Swarada Bhosale",
			Department: "IT",
		}
		result, err := studRepo.InsertStudent(&stud)
		if err != nil {
			t.Fatal("Insert 2 operation failed", err)
		}
		t.Log("Insert 2 successful", result)
	})

	// get student by ID
	t.Run("GetStudent1", func(t *testing.T) {
		result, err := studRepo.FindStudentByID(stud1)
		if err != nil {
			t.Fatal("Get operation failed", err)
		}
		t.Log("Student 1", result.Name)
	})

	// get all students
	t.Run("GetAllStudents", func(t *testing.T) {
		result, err := studRepo.FindAllStudent()
		if err != nil {
			t.Fatal("Get all operation failed", err)
		}
		t.Log("All students", result)
	})

	// update student 1
	t.Run("UpdateStudent1", func(t *testing.T) {
		stud := model.Student{
			StudentID:  stud1,
			Name:       "Vaishnavi Ladda",
			Department: "IT",
		}
		res, err := studRepo.UpdateStudent(stud1, &stud)
		if err != nil {
			t.Fatal("Update operation failed", err)
		}
		t.Log("Update 1 successful", res)
	})

	// get student by ID again
	t.Run("GetStudent1AfterUpdate", func(t *testing.T) {
		result, err := studRepo.FindStudentByID(stud1)
		if err != nil {
			t.Fatal("Get operation failed", err)
		}
		t.Log("Student 1 after update", result.Name)
	})

	// delete student 1
	t.Run("DeleteStudent1", func(t *testing.T) {
		res, err := studRepo.DeleteStudent(stud1)
		if err != nil {
			t.Fatal("Delete operation failed", err)
		}
		t.Log("Delete 1 successful", res)
	})

	// get all students again
	t.Run("GetAllStudentsAfterDelete", func(t *testing.T) {
		result, err := studRepo.FindAllStudent()
		if err != nil {
			t.Fatal("Get all operation failed", err)
		}
		t.Log("All students after delete", result)
	})

	// delete all students
	t.Run("DeleteAllStudents", func(t *testing.T) {
		res, err := studRepo.DeleteAllStudent()
		if err != nil {
			t.Fatal("Delete all operation failed", err)
		}
		t.Log("Delete all successful", res)
	})
}
