package main

import (
	"math"
	"fmt"
)

func main(){
	sum_s :=0
	s_sum :=0
	for i:=1;i<=100;i++ {
		sum_s += int(math.Pow(float64(i),2))
		s_sum += i
	}
	fmt.Println(s_sum * s_sum - sum_s)
}
