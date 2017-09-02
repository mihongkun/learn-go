package main

import (
	"fmt"
	"time"
)

func main(){
	for i := 0;i < 10;i ++ {
		go func(i int) {
			fmt.Println(i,"start")
			time.Sleep(time.Second)
			fmt.Println(i,"end")
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}
