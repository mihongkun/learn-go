package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32,error) {
	bs,err := ioutil.ReadFile(filename)
	if err != nil {
		return 0,err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(),nil

}




func main(){
	var filename1 string = "test1.txt"
	var filename2 string = "test2.txt"
	h1,err := getHash(filename1)
	if err!= nil{
		fmt.Println("Open file failed,filename:",filename1)
		return
	}
	
	h2,err := getHash(filename2)
	if err != nil {
		fmt.Println("Open file failed,filename:",filename2)
		return
	}
	fmt.Println(h1,h2,h1==h2)	

}


