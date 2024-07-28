package repository

import (
	"context"
	"crud_go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepo struct {
	MongoCollection *mongo.Collection
}

func (r StudentRepo) InsertStudent (stud *model.Student) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), stud)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *StudentRepo) FindStudentByID (studID string) (*model.Student, error) { 
	var stud model.Student
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "student_id", Value: studID}}).Decode(&stud)
	if err !=nil { 
		return nil, err
	}
	return &stud, nil
}

func (r *StudentRepo) FindAllStudent () ([]model.Student, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var students []model.Student
	err = results.All(context.Background(), &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepo) UpdateStudent (studID string, stud *model.Student) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "student_id", Value: studID}}, bson.D{{Key: "$set", Value: stud}})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (r *StudentRepo) DeleteStudent (studID string) (int64, error) {	
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "student_id", Value: studID}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}	


func (r *StudentRepo) DeleteAllStudent () (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
