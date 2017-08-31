package main

import ("fmt";"math")

func main(){
	var x,y float64 = 0.0,0.0
	var x1,y1 float64 = 10.0,10.0
	var cx,cy,r float64 = 0.0,0.0,5.0	
	fmt.Println(rectangleArea(x,y,x1,y1))
	fmt.Println(distance(x,y,y1,y1))
	fmt.Println(circleArea(cx,cy,r))
}

func distance(x,y,x1,y1 float64) float64 {
	return math.Sqrt(math.Pow(x1-x,2) + math.Pow(y1-y,2))
}

func rectangleArea(x,y,x1,y1 float64) float64{
	return distance(x,y,x,y1) * distance(x,y,x1,y)
}

func circleArea(x,y,r float64) float64{
	return math.Pi * math.Pow(r,2)
}




