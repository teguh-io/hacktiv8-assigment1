package models

import "fmt"

type Student struct {
	Name    string
	Email   string
	Address string
	Job     string
	Goal    string
}

func (s *Student) GetStudent() {
	fmt.Println("Student Name:", s.Name)
	fmt.Println("Student Email:", s.Email)
	fmt.Println("Student Address:", s.Address)
	fmt.Println("Student Job:", s.Job)
	fmt.Println("Student Goals:", s.Goal)

}
