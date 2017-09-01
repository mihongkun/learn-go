package main

import "fmt"

func main(){
	a := new(Android)
	a.Person.Name = "Hello World!"
	a.Person.Talk()
	a.Talk()
	p := Person{Name:"Person"}
	p.Talk()
}


type Person struct{
	Name string
}

func (p *Person) Talk(){
	fmt.Println("Hi,My name is",p.Name)
}
type Android struct {
	Person 
	Model string
}



