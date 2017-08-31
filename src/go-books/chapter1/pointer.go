package main

import "fmt"

func main(){
	x := 5
	zero_non_pointer(x)
	fmt.Println(x)
	zero_pointer(&x)
	fmt.Println(x)
}

func zero_pointer(x *int){
	*x = 0
}

func zero_non_pointer(x int){
	x = 0
}
