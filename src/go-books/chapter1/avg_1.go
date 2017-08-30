package main

import "fmt"

func main(){
	var x [5]float64
	x[0] = 99
	x[1] = 77
	x[2] = 88
	x[3] = 100
	x[4] = 66
	var total float64 = 0 
	for i:=0;i<len(x);i++{
		total +=x[i]
	}
	fmt.Println(total/float64(len(x)))

}
