package main

import (
	"fmt"
	"runtime"
)

func prendqqchetrendqqch (x, y int) (r, s int) {
	r = x + 1
	s = y + 1	
	return 
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Println(prendqqchetrendqqch(5, 10))
}