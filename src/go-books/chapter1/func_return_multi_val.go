package main

import "fmt"

func main(){
	fmt.Println(f())	
}

func f()(r1 int,r2 int){
	r1 = 1
	r2 = 2
	return
}
