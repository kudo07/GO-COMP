package main

import (
	"fmt"
	"net/http"
)

func importFunc() {
	fmt.Println("Hello, Go std. library")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Reponse", resp.Status)
}
