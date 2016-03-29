package main

import (
	"fmt"
	"time"
	"sync"
	)

var mx sync.Mutex

func philosophe(name string, ch chan int) {
	fmt.Println("nbonjour ", name)
	ch <- 1
}

func thread(nb int) {
	fmt.Println("je vais lock nb", nb)
	mx.Lock()
	fmt.Println("je vais dormir nb", nb)

	time.Sleep(time.Duration(nb) * time.Second)
	mx.Unlock()
	fmt.Println("je vais unlock nb", nb)
}

func threadBis(mx chan int, nb int) {
	fmt.Println("je vais lock nb", nb)
	mx <- nb
	fmt.Println("je vais dormir nb", nb)
	time.Sleep(time.Duration(nb) * time.Second)
	fmt.Println("je vais unlock nb", nb)
}

func main() {
	philoName := []string{"Socrate", "Platon", "Pascal", "Descartes", "Marx"}
	nb := len(philoName)
	ch := make(chan int, nb)
	for i := 0; i < nb; i++ {
		go philosophe(philoName[i], ch)
	}
	fmt.Println(philoName)
	//time.Sleep(time.Second)
	//time.Sleep(time.Second)
	<-ch

}