package main

import "fmt"

func truc(ch chan int, x int) {
	ch <- x
	fmt.Println("toto", x)
}

func main() {
	ch := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go truc(ch, i)	
	}
	fmt.Println(<-ch)
}