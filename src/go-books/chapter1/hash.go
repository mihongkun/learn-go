package main


import (
	"fmt"
	
	"hash/crc64"
	
	"hash/crc32"
	
)
func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h1 := crc64.New(new *Table)
	h1.Write([]byte("test"))
	v1 := h1.Sum64()
	fmt.Println(v1)
}
