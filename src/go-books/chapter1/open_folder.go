package main

import (
	"fmt"
	"os"
)

func main(){
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Open folder failed!")
		return
	}

	defer dir.Close()

	fileInfos,err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir failed!")
		return
	}
	for _,fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

