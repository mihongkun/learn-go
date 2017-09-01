package main

import "fmt"

func main(){
	go f("Hello\t")
	go f("World\t")
	for i:=0;i<100;i++{
		fmt.Println("Myself\t :",i)
	}
	var input string
	fmt.Scanln(&input)
}


func f(s string){
	for i:=0;i<100;i++{
		fmt.Println(s,":",i)
	}
}



