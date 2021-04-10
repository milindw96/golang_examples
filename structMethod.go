package main

import "fmt"

type Student struct {
	name  string
	grade []int
	age   int
}

// get method
func (s Student) getAge() int {
	return s.age
}
func (s Student) getName() string {
	return s.name
}

func (s Student) getGrade() []int {
	return s.grade
}

func (s Student) getAvgGrade() float64 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float64(sum) / float64(len(s.grade))
}

// Set method
// if *Student(pointer not given it will only chnage the local copy)
func (s *Student) setAge(age int) {
	s.age = age
}
func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) setGrade(grade []int) {
	s.grade = grade
}

func main() {
	s1 := Student{"Milind", []int{57, 90, 75, 90}, 22}
	// fmt.Println(s1.name, s1.grade, s1.age)
	fmt.Println("Name: ", s1.getName())
	fmt.Println("Age:", s1.getAge())
	fmt.Println("Grade:", s1.getGrade())
	// Setting new values
	s1.setName("Milind1")
	s1.setAge(25)
	fmt.Println(s1.name, s1.grade, s1.age)

	fmt.Println("Avg: ", s1.getAvgGrade())
}
