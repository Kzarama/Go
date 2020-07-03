package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 10 {
			break
		}
		for _, server := range servers{
			go checkServer(server, channel)
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-channel)
		i++
	}

	finish := time.Since(start)
	fmt.Printf("Execution time %s\n", finish)
}

func checkServer(server string, channel chan string){
	_, err := http.Get(server)
	if err != nil{
		channel <- server + " not available"
	} else {
		channel <- server + " is ok"
	}
}