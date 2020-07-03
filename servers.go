package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	servers := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers{
		checkServer(server)
	}
	finish := time.Since(start)
	fmt.Printf("Execution time %s\n", finish)
}

func checkServer(server string){
	_, err := http.Get(server)
	if err != nil{
		fmt.Println(server, "not available")
	} else {
		fmt.Println(server, "is ok")
	}
}