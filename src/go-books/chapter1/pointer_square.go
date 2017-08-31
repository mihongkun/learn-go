package main

import "fmt" 

func main(){
	x := 1.5
	area := square(&x)
	fmt.Println(area)
}

func square(x *float64) float64{
	return *x * *x
}
