package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	for i:=1;;i++ {
		j:=1
		for ;j<=20;j++{
			if(i % j != 0) {break}
		}
		if(j > 20) {
			fmt.Println(i)
			break
		}
	}
	t:=time.Now()
	fmt.Println( t.Sub(start))
}

