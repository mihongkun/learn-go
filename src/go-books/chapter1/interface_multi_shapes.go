package main


import ("fmt";"math")

func main(){
	c := Circle{10,10,10}
	r := Rectangle{0,0,10,10}
	m := MultiShape{[]Shape{&c,&r}}
	fmt.Println(m.area())

}

type Rectangle struct {
	x,y,x1,y1 float64
}

type Circle struct{
	x,y,r float64
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (c *Circle) area() float64{
	return math.Pow(c.r,2) * math.Pi	
}

func (r *Rectangle) area() float64{
	return distance(r.x,r.y,r.x,r.y1) * distance(r.x,r.y,r.x1,r.y)
}

func distance(x,y,x1,y1 float64) float64{
	return math.Sqrt(math.Pow(x1-x,2) + math.Pow(y1-y,2))
}


func (m *MultiShape) area() float64 {
	var area float64
	for _,s := range m.shapes {
		area += s.area()
	}
	return area
} 
