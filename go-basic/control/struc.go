package control

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Shape interface {
	Area() float64
}

func Structure() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	fmt.Printf("Hello, my name is %s %s and I am %d years old.\n", p.FirstName, p.LastName, p.Age)

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println("Array", myArray)

	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice)
	fmt.Println("Length Slice:", len(mySlice))
	fmt.Println("Cap Slice:", cap(mySlice))
	fmt.Println("Append Slice:", append(mySlice, 60, 70, 80))

	// Slicing a slice
	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))

	myNewSlice := myArray[:]
	fmt.Println("chain slice with array", append(myNewSlice, mySlice...))

	myMap := make(map[string]int)

	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	fmt.Println("Apple: ", myMap["apple"])

	delete(myMap, "apple")

	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	val, ok := myMap["apple"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found in map")
	}

	var person1 [1]Person
	person1[0].FirstName = "John"
	person1[0].LastName = "Doe"
	person1[0].Age = 25

	fmt.Println("Person1 name is", person1[0].FirstName)
}
