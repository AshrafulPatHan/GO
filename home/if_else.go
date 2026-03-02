package main

import "fmt"

func if_Else(){
	// if function
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
	// if else function
	time := 20
	if (time < 18) {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
	// else if function
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}
