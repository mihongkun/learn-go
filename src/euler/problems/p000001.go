package main

import (
	"fmt"
)

func main(){
	var result uint64 = 0
	for i:=0;i<1000;i++ {
		if(i % 3 == 0 || i % 5 ==0) {
			result += uint64(i)
		}
	}
	fmt.Println("Result:",result)
}
