package main

import (
	"fmt"
	"snoopy/asciiart"
)

func main() {
	pick := 0
	fmt.Println("Hello, welcome to CLI Snoopy\nPlease, pick your Snoopy.")
	for true {
		fmt.Println("\n1. Classic Snoopy\n2. Loving Snoopy")
		fmt.Println("For exiting, type 7")
		fmt.Scan(&pick)
		switch pick {
		case 1:
			asciiart.Snoopy()
		case 2:
			asciiart.Loving()
		case 7:
			return
		default:
			fmt.Println("Not supported number. You won't get your Snoopy.")
		}
	}

}
