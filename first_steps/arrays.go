package main

import "fmt"

func main() {
	var array_1 = [3]int{} // not initialized
	array_2 := [3]float32{1.1,2.2} // partially inicialized 
	var array_3 = [...]string{"string","muito","foda"} 
	var array_4 = [4]uint8{1:20,3:40} // [0,20,0,40]
	array_3[2] = "linda"

	fmt.Printf("ARRAY 1: %v LENGTH:%v \n",array_1,len(array_1))
	fmt.Println(array_2)
	fmt.Printf("TYPE : %T ALL VALUES: %v, VALUE 1: %v \n",array_3,array_3,array_3[0])
	fmt.Println(array_4)
}	