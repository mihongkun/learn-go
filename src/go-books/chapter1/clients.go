package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	go client()
	var input string
	fmt.Scanln(&input)
}
func client() {
  // connect to the server
  c, err := net.Dial("tcp", "127.0.0.1:9999")
  if err != nil {
    fmt.Println(err)
    return
  }

  // send the message
  msg := "Hello World"
  fmt.Println("Sent:", msg)
  err = gob.NewEncoder(c).Encode(msg)
  if err != nil {
    fmt.Println(err)
  }

  c.Close()
}

