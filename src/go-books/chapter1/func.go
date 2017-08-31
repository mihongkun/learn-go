package main

import "fmt"



func main(){
	xs :=[]float64{1,2,3,4,5,6,7,8,9,}
	fmt.Println("The average of",xs,"is",averages(xs))
}


func averages(xs []float64) float64{
	total := float64(0)
	for _,v:= range xs{
		total += v
	} 
	
	return total/float64(len(xs))
}

