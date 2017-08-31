package main

import "fmt"

func main(){
	x := new(int)
	one(x)
	fmt.Println(*x)	
}

func one(x *int){
	*x = 1
}
