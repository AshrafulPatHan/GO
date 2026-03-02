package main

import "fmt"

func array(){
	// var array_name = [length]datatype{values} # here length is defined
	// var array_name = [...]datatype{values} # here length is inferred

	var arr1 = [3]int{1,2,3} // set variable
	arr2 := [...]int{4,5,6,7,8}
	arr3 := [...]string{"Go", "array", "Are", "Powerful"}

	fmt.Println(arr1,arr2,arr3)
	

}