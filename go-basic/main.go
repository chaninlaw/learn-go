package main

import (
	"fmt"

	"github.com/chaninlaw/control"
)

func main() {
	fullName := "John"
	const lastName = "Doe"
	age := 25

	fmt.Printf("Hello, my name is %s %s and I am %d years old.\n", fullName, lastName, age)

	sisterName := "Jane" + " " + lastName
	age = 26

	fmt.Printf("My sister name is %s and I am %d years old.\n", sisterName, age)
	fmt.Println(!false)

	if !true {
		fmt.Println("Hello")
	} else {
		fmt.Println("Goodbye")
	}

	control.Loop()
	control.Structure()
	whoName := control.Function("Ninja")
	fmt.Printf("Hello, %s again!\n", whoName)

	student := control.Student{
		FirstName: "Jim",
		LastName:  "Doe",
	}

	fullName = student.FullName()
	fmt.Println("FullName:", fullName)

	numbers := control.IntSlice{1, 2, 3, 4, 5}

	fmt.Println("Sum: ", numbers.Sum())
	fmt.Println("Average: ", numbers.Average())

	man := control.Man{
		Name: "John",
	}
	dog := control.Dog{
		Name: "Fido",
	}
	control.MakeSound(man)
	control.MakeSound(dog)

	// Pointer Mutability
	x := 20
	control.ChangeValue(&x)
	fmt.Println("New X: ", x)

	// Efficient large structs
	// ส่ง Struct หรือ Config ที่ใหญ่ๆเข้าไปทำให้ประหยัด Memory
	emp := control.Employee{
		Name:   "John",
		Salary: 1000,
	}
	control.GiveRaise(&emp, 5000)
	fmt.Println("After raise: ", emp)

	// Link list
	var head *control.ListNode
	control.Prepend(&head, 1)
	control.Prepend(&head, 2)
	control.Prepend(&head, 3)

	fmt.Println("List: ", head.Next.Next.Value)

	// Error
	control.ErrorHandling()
}
