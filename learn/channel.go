package main

import (
	"fmt"
	"time"
)

 func print( ch chan string) {
	
	 defer func (){ ch <- "Goroutine is done "}() // send data

		fmt.Println("Processing...")
		time.Sleep(2 * time.Second)
	}

func main() {
	ch := make(chan string) // create channel

	go print(ch)

	
	msg := <-ch // receive data
	fmt.Println(msg)
}