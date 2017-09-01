package main

import (
	"fmt"
	"os"
)


func main(){
	file,err := os.Open("test.properties")
	if err != nil {
		fmt.Println("Open file fail")	
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	
	if err != nil {
		fmt.Println("Get file stat fail")
		return
	}
	
	bs := make([]byte,stat.Size())
	_,err = file.Read(bs)
	if err != nil{
		fmt.Println("Read file fail")
		return
	}
	
	str := string(bs)
	fmt.Println(str)
	
}
