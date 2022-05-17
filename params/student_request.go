package params

type StudentRequest struct {
	Name    string
	Email   string
	Address string
	Job     string
	Goal    string
}

func CreateStudent(name, email, address, job, goal string) *StudentRequest {
	return &StudentRequest{
		Name:    name,
		Email:   email,
		Address: address,
		Job:     job,
		Goal:    goal,
	}
}
