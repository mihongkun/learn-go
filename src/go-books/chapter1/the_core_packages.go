package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(
		strings.Contains("test","es"),
		strings.Count("test","t"),
		strings.HasPrefix("test","te"),
		strings.HasSuffix("test","st"),
		strings.Index("test","e"),
		strings.Join([]string{"a","b"},"-"),
		strings.Repeat("a",10),
		strings.Replace("baba","a","b",2),
		strings.Split("a-b-c-d-e","-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)
	
}
