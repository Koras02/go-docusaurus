package main

import "fmt"

func main() {
    var age int = 27

	if age >= 18 {
		fmt.Println("성인입니다.")
	} else {
		fmt.Println("청소년 입니다.")
	}

	// if-elif-else 조건 체크
	var grade int = 90
	if grade >= 90 {
		fmt.Printf("A")
	} else if grade >= 80 {
		fmt.Printf("B")
	} else if grade >= 70 {
		fmt.Printf("C")
	} else if grade >= 60 {
		fmt.Printf("D")
	} else {
		fmt.Printf("F")
	}

}		


