package main

import "fmt"

func main() {
	// string
	var gretting string = "Hello, Go!, from a variable";
	fmt.Println(gretting); // Hello, Go!, from a variable
	// int
	var isActive bool = true;
	fmt.Println(isActive); // true 

	// slice
	var numer []int = []int{1, 2, 3, 4, 5};
	fmt.Println(numer); // [1 2 3 4 5]

	// map
	var user map[string]string = map[string]string{
		"name": "John",
		"age":  "30",
	}
	fmt.Println(user); // map[age:30 name:John]

	// struct
	type Person struct {
		name string 
		Age int 
		Job string
	}

	var person Person = Person {
		name: "John",
		Age: 30,
		Job: "Developer",
	}
	fmt.Println(person); // {John 30 Developer}
}