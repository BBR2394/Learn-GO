package main

import (
	"fmt"
	"time"
	"math/rand"
	)

func thread (id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(99) * 1000))
	fmt.Println("je suis une go routine n°", id)
	c <- id
}

func main () {
	channel := make(chan int)
	/*thread(-1)*/
	for  i := 0; i < 100 ; i += 1 {
		go thread(i, channel)
	}
	fmt.Println("bonjour")
	for i := 0; i < 100 ; i += 1 {
		fin := <-channel
		fin += 1
	}
}