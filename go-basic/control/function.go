package control

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
}

type IntSlice []int

type Speaker interface {
	Speak() string
	run() string
}
type Dog struct {
	Name string
}

type Man struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) run() string {
	return "90km/h"
}

func (p Man) Speak() string {
	return "Hello!"
}

func (p Man) run() string {
	return "50km/h"
}

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func (s IntSlice) Average() float64 {
	return float64(s.Sum()) / float64(len(s))
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func Function(name string) string {
	fmt.Printf("------- %s -------\n", name)
	return "Hello, " + name
}
