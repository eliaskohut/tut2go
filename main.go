package main

import (
	"fmt"

	"github.com/google/uuid"
	Appdispatcher "github.com/grupawp/appdispatcher"
)

type Student struct {
	FirstName, LastName string
	applicationID       uuid.UUID
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
}

func (s Student) FullName() string {
	return fmt.Sprintf("%s %s", s.FirstName, s.LastName)
}

func main() {
	var a Student
	a.FirstName = "Marek"
	a.LastName = "Dobrzański"
	fmt.Printf(a.FullName())
	Appdispatcher.Submit(a)
}
