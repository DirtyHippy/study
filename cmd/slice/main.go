package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[0:3]
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	slice = append(slice, "arg1")
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	slice = append(slice, "arg2")
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	slice = append(slice, "arg3")
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
}
