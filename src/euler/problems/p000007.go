package main

import (
	"fmt"
	"math"
)

func main(){
	prime := uint64(2)
	for i := 1;i <= 10001; {
		
		if isPrime(prime) {
			i++
		}
		prime ++
		
	}	
	fmt.Println(prime - 1)
}


func isPrime(n uint64) bool{
	for i:=uint64(2);i <= uint64(math.Sqrt(float64(n)));i++{
		if n % i == 0 {
			return false
		}
	}
	return true
}


