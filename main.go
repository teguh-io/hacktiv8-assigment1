package main

import (
	"assigment1-basic/models"
	"assigment1-basic/params"
	"assigment1-basic/repositories"
	"fmt"
	"os"
)

func main() {
	students := make(map[string]models.Student)
	req1 := params.CreateStudent("Teguh Ainul Darajat", "teguh@email.com", "Indonesia", "Junior Backend Engineer", "Pro Backend Engineer")
	req2 := params.CreateStudent("Ujang", "ujang@email.com", "Bandung", "Student", "Pro Backend Engineer")

	students["1"] = repositories.InsertStudent(req1)
	students["2"] = repositories.InsertStudent(req2)

	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Please input student ID")
		return
	}

	for _, argument := range arguments {
		if _, ok := students[argument]; !ok {
			fmt.Printf("Student with ID %s not found\n", argument)
			continue
		}

		fmt.Println("Name:", students[argument].Name)
		fmt.Println("Email:", students[argument].Email)
		fmt.Println("Address:", students[argument].Address)
		fmt.Println("Job:", students[argument].Job)
		fmt.Println("Goal:", students[argument].Goal)
		fmt.Println()
	}
}
