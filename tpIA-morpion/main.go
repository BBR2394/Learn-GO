package main

import "fmt"
import "os"
import "bufio"

var tab [10][10]string
var decide [10][10]int
var firstplay bool

func playerPlay() {

	bio := bufio.NewReader(os.Stdin)
	read, hasMoreInLine, err := bio.ReadLine()
	input := string(read)
	fmt.Println(input)
	fmt.Println("fin")
	_ = hasMoreInLine
	_ = err
}

func initDecideTable() {
	for c := 0; c < 10; c++{
			for i := 0; i < 10; i++ {
				decide[c][i] = 10
				fmt.Printf("%d-", decide[c][i])
			}
			fmt.Printf("\n")
		}
	decide[4][4] = 1
	decide[5][5] = 1
	decide[4][5] = 1
	decide[5][4] = 1
}


func theGame() {
	if firstplay == true {
		initDecideTable()
	}

}

func main() {
	firstplay = true	
	fmt.Printf("bonjour ca marche\n")
	for c := 0; c < 10; c++{
		for i := 0; i < 10; i++ {
			tab[c][i] = "-"
			fmt.Printf(tab[c][i])
		}
		fmt.Printf("\n")
	}
	theGame()
}