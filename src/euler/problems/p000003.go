package main 

import (
	"fmt"
	"math"
)

func isPrime(x uint64) bool{
	for i:=uint64(2);i< uint64(math.Sqrt(float64(x)));i++{
		if x % i ==0 {
			return false
		}	
	}
	return true
}

func main() {
	fmt.Println(maxfactor(600851475143))	
}

func maxfactor(x uint64) uint64{
	i:=x
	j:=uint64(1)
	for {
		if isPrime(i) {
			return i
		}
		j++
		for {	
			if(x % j == 0) {break}
			j++
		}
		i = x / j
	}
}

