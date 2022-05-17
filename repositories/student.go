package repositories

import (
	"assigment1-basic/models"
	"assigment1-basic/params"
)

func InsertStudent(req *params.StudentRequest) models.Student {
	student := models.Student{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Job:     req.Job,
		Goal:    req.Goal,
	}

	return student
}
