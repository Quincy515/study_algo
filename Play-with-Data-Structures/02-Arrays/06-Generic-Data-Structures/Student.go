package _6_Generic_Data_Structures

import "fmt"

type Student struct {
	name  string
	score int
}

func NewStudent(studentName string, studentScore int) *Student {
	return &Student{name: studentName, score: studentScore}
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(name: %s, score: %d)", s.name, s.score)
}
