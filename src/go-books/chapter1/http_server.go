package main

import (
	"fmt"
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/hello",hello)	
	fmt.Println("Http server listening :9000")	
	http.ListenAndServe(":9000",nil)
}

func hello(res http.ResponseWriter,req *http.Request){
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				<h1 style="color:red;">Hello World!</h1>
			</body>
		</html>
			

		`,
	)

	


}

