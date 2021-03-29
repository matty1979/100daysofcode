package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(" Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(" OS X. ")
	case "linux":
		fmt.Println("Linux.")
	default:
		// free bsd, openbsd,
		// plan9, windows
		fmt.Printf("%s. \n", os)
	}
}
