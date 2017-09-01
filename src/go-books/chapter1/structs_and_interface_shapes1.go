package main

import ("fmt";"math")

type Circle struct {
	x float64
	y float64
	r float64
}

func main(){
	c := Circle{x:0, y:0, r:5}
	fmt.Println(circleArea(c))	



}

func circleArea(c Circle) float64{
	return math.Pow(c.r,2) * math.Pi
}




