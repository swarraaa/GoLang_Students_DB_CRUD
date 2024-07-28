package model

type Student struct {
	StudentID  string `json:"student_id,omitempty" bson:"student_id"` 
	Name       string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}
