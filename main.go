package main

import (
	"fmt"
	"os"
	"snoopy/asciiart"
)

func main() {
	fmt.Print("\033[H\033[2J")
	pick := 0
	fmt.Println("Hello, welcome to CLI Snoopy\nPlease, pick your Snoopy.")
	for {
		fmt.Println("\n1. Classic Snoopy\n2. Loving Snoopy")
		fmt.Println("For exiting, type 7")
		fmt.Scan(&pick)
		switch pick {
		case 1:
			asciiart.Snoopy()
		case 2:
			asciiart.Loving()
		case 7:
			os.Exit(0)
		default:
			fmt.Println("Not supported number. You won't get your Snoopy.")
		}
	}
}
