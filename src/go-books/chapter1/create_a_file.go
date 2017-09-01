package main

import (
	"fmt"
	"os"
)

func main(){
	file,err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Create file failed!")
		return
	}

	defer file.Close()
	file.WriteString("Hello,World!")
	fmt.Println("Create file finish!")

}
