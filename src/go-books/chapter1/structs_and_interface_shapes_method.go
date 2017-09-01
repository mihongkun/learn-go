package main

import ("fmt";"math")

func main(){
	r := Rectangle{0,0,10,20}
	fmt.Println(r.area())	
}

type Rectangle struct{
	x,y,x1,y1 float64
}

func (r *Rectangle) area() float64{
	return distance(r.x,r.y,r.x,r.y1) * distance(r.x,r.y,r.x1,r.y)
}

func distance(x,y,x1,y1 float64) float64{
	return math.Sqrt(math.Pow((x1-x),2) + math.Pow(y1-y,2))
} 


