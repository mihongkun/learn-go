package main

import ("fmt";"math")

func main(){
	c := Circle{1,2,10}
	fmt.Println(c.area())
}

type Circle struct {
	cx,xy,r float64
}


func (c *Circle) area() float64{
	return math.Pow(c.r,2) * math.Pi
} 


