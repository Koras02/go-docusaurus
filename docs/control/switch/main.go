package main

import "fmt"

func main() {
     selector := 5

	 switch selector {
	 case 1:
		fmt.Println("Booting")
	 case 2:
		fmt.Println("Shut Down")
	 case 3:
		fmt.Println("Reboot")
	 default: 
	    fmt.Println("Running...")
	 }
}