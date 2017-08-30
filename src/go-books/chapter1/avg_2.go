package main

import "fmt"

func main(){
	x:=[5]float64{100,100,99,99,98}
	total:=float64(0)
	for i:=0;i<len(x);i++{
	
		total += x[i]
	}
	fmt.Println(total/float64(len(x)))
}
