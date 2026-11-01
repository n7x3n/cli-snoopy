package main

import (
	"bufio"
	"fmt"
	"os"
	"snoopy/asciiart"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("\033[H\033[2J")
	pick := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, welcome to CLI Snoopy\nPlease, pick your Snoopy.")
	for {
		fmt.Println("\n1. Classic Snoopy\n2. Loving Snoopy")
		fmt.Println("For exiting, type q")
		/*_, err := fmt.Scan(&pick)
		if err != nil {
			fmt.Println("\nExiting...")
			return
		}
		*/
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if input == "q" {
				fmt.Println("\033[H\033[2J")
				fmt.Println("Goodbye...\nSee you soon...")
				return
			}
			var err error
			pick, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Please, enter correct number, otherwise no Snoopy")
				continue
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}

		switch pick {
		case 1:
			asciiart.Snoopy()
		case 2:
			asciiart.Loving()
		default:
			fmt.Println("Not supported number. You won't get your Snoopy.")
		}
	}
}
