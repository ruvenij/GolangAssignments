package main

import "fmt"

func PublishMessages(ch chan string) {
	ch <- "Hello GoRoutine"
}

func ReadMessages(ch chan string) {
	message := <-ch
	fmt.Println("Consumed " + message)
}
