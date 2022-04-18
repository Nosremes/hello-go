package main

import "fmt"

func main() {
	var array_1 = [5]int{1,2,3}

	slice_1 := []int{1,2,3}
	slice_2 := []int{}
	slice_3 := array_1[0:3] // slice made from array
	slice_4 := make([]int, 5, 10) //type, length, capacity
	slice_4 = append(slice_4, 12, 14, 15)

	fmt.Printf("TYPE: %T, LENGTH: %v CAPACITY: %v, CONTENT: %v\n",slice_1,len(slice_1),cap(slice_1),slice_1) 
	fmt.Printf("TYPE: %T, LENGTH: %v CAPACITY: %v, CONTENT: %v\n",slice_2,len(slice_2),cap(slice_2),slice_2) 
	fmt.Printf("TYPE: %T, LENGTH: %v CAPACITY: %v, CONTENT: %v\n",slice_3,len(slice_3),cap(slice_3),slice_3)
	fmt.Printf("TYPE: %T, LENGTH: %v, CAPACITY: %v, CONTENT: %v\n",slice_4,len(slice_4),cap(slice_4),slice_4)

	all_slices := append(slice_1, slice_4...)
	fmt.Println(all_slices)
}	
