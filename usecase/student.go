package usecase

import (
	"crud_go/model"
	"crud_go/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error string 	`json:"error,omitempty"`
}

func (svc *StudentService) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var stud model.Student
	err := json.NewDecoder(r.Body).Decode(&stud)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid request payload", err)
		res.Error = err.Error()
		return
	}

	stud.StudentID = uuid.New().String()
	repo := repository.StudentRepo{MongoCollection: svc.MongoCollection}
	insertID, err := repo.InsertStudent(&stud)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while inserting student", err)
		res.Error = err.Error()
		return
	}
	res.Data = stud.StudentID
	w.WriteHeader(http.StatusOK)
	log.Println("Student inserted successfully with ID", insertID)

}

func (svc *StudentService) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	studId := mux.Vars(r)["id"]
	log.Println("Student ID", studId)

	repo := repository.StudentRepo{MongoCollection: svc.MongoCollection}
	stud, err := repo.FindStudentByID(studId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Student not found", err)
		res.Error = err.Error()
		return
	}
	res.Data = stud
	w.WriteHeader(http.StatusOK)
	log.Println("Student found successfully", stud)
}

func (svc *StudentService) GetAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.StudentRepo{MongoCollection: svc.MongoCollection}
	stud, err := repo.FindAllStudent()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Student not found", err)
		res.Error = err.Error()
		return
	}
	res.Data = stud
	w.WriteHeader(http.StatusOK)
}

func (svc *StudentService) UpdateStudentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	studId := mux.Vars(r)["id"]
	log.Println("Student ID", studId)

	if studId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Student ID is missing")
		res.Error = "Student ID is missing"
		return
	}

	var stud model.Student
	err := json.NewDecoder(r.Body).Decode(&stud)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid request payload", err)
		res.Error = err.Error()
		return
	}

	repo := repository.StudentRepo{MongoCollection: svc.MongoCollection}
	stud.StudentID = studId

	count, err := repo.UpdateStudent(studId, &stud)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while updating student", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)

}

func (svc *StudentService) DeleteStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	studId := mux.Vars(r)["id"]
	log.Println("Student ID", studId)

	if studId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Student ID is missing")
		res.Error = "Student ID is missing"
		return
	}

	repo := repository.StudentRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.DeleteStudent(studId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while deleting student", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)

}

func (svc *StudentService) DeleteAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.StudentRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.DeleteAllStudent()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while deleting student", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}