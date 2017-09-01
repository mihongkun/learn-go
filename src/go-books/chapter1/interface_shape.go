package main

import "fmt"

func main(){
	
}

type Shape interface {
	area() float64
}

func totalArea(shapes...Shape) float64{
	var area float64
	for _,s := range shapes{
		area += s.area()
	}
	return area
}



