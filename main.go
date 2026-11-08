package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/n7x3n/cli-snoopy/asciiart"
)

func main() {
	fmt.Print("\033[H\033[2J")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, welcome to CLI Snoopy\nPlease, pick your Snoopy.")
	for {
		pick := 0
		for i, art := range asciiart.Gallery {
			fmt.Printf("%d. %s\n", i+1, art.Name)
		}
		fmt.Println("For exiting, type q")
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
		} else {
			fmt.Println("Exiting...")
			return
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
		if pick > 0 && pick <= len(asciiart.Gallery) {
			artNumber := pick - 1
			asciiart.ArtPrint(asciiart.Gallery[artNumber].Art)
		} else {
			fmt.Println("Sorry, you must've picked a wrong number. I don't have that Snoopy.")
		}
	}
}
