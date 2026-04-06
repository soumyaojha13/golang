package main

import (
	"errors"
	"fmt"
)

// Subject struct
type Subject struct {
	Name  string
	Marks int
}

// GradeSystem struct
type GradeSystem struct {
	students map[string][]Subject
}

// Constructor
func NewGradeSystem() *GradeSystem {
	return &GradeSystem{
		students: make(map[string][]Subject),
	}
}

// Add student (FIXED: pointer receiver)
func (g *GradeSystem) AddStudent(name string) error {

	if _, ok := g.students[name]; ok {
		return errors.New("student already exists")
	}

	g.students[name] = []Subject{}
	return nil
}

// Add grade
func (g *GradeSystem) AddGrade(name, subject string, marks int) error {

	if _, exists := g.students[name]; !exists {
		return errors.New("student not found")
	}

	g.students[name] = append(g.students[name], Subject{
		Name:  subject,
		Marks: marks,
	})

	return nil
}

// Get all grades
func (g *GradeSystem) GetGrades(name string) ([]Subject, error) {

	grades, ok := g.students[name]
	if !ok {
		return nil, errors.New("student not found")
	}

	return grades, nil
}

// Get specific subject grade
func (g *GradeSystem) GetGrade(name, subject string) (int, error) {

	grades, ok := g.students[name]
	if !ok {
		return 0, errors.New("student not found")
	}

	for _, sub := range grades {
		if sub.Name == subject {
			return sub.Marks, nil
		}
	}

	return 0, errors.New("subject not found")
}

func main() {

	system := NewGradeSystem()

	// Add students
	system.AddStudent("Soumya")
	system.AddStudent("Rahul")

	// Add grades
	system.AddGrade("Soumya", "Math", 90)
	system.AddGrade("Soumya", "Science", 85)
	system.AddGrade("Rahul", "English", 78)

	// Get all grades
	grades, err := system.GetGrades("Soumya")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Soumya Grades:", grades)

	// Get specific subject
	marks, err := system.GetGrade("Soumya", "Math")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Math Marks:", marks)
}
