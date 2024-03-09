package main

import (
	"fmt"
	"sync"
	"time"
)

func Hello() {
	fmt.Println("Hello world!")
}

func halo() {
	fmt.Println("Shinnamon")

	time.Sleep(3 * time.Second)

	var wg sync.WaitGroup

	msgchan := make(chan string, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()

		msgchan <- "Channel"

		Hello()
	}()
	go func() {
		defer wg.Done()

		msg := <-msgchan

		fmt.Println(msg)
	}()

	fmt.Println("Good bye")
}
