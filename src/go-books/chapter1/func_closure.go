package main

import "fmt"

func main(){
	add:=func(x,y int) int{
		return x + y;
	}
	fmt.Println(add(4,5))	
}


