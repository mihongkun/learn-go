package main

import "fmt"

func main(){
	var x string = "Hello World"
	fmt.Println(x)
	var y string
	y = "MiHongkun"
	fmt.Println(y)
	var name string = "Max"
	fmt.Println("My name is ",name)
	const dogName string = "Tiger Gril"
	fmt.Println("My dog's name is ",dogName)
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println("a",a,",b",b,",c",c)
	fmt.Println("Enter a number:")
	var input float64
	fmt.Scanf("%f",&input)
	output := input * 2.0
	fmt.Println("output is %20.5f",output)
}

