package asciiart

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Ascii struct {
	Name string
	Art  string
}

//go:embed assets/classicSnoopy.txt
var classicSnoopy string

//go:embed assets/lovingSnoopy.txt
var lovingSnoopy string

//go:embed assets/joeCool.txt
var joeCool string

//go:embed assets/smilingSnoopy.txt
var smilingSnoopy string

var Gallery = []Ascii{
	{Name: "Classic Snoopy",
		Art: classicSnoopy,
	}, {
		Name: "Loving Snoopy",
		Art:  lovingSnoopy,
	}, {
		Name: "Joe Cool",
		Art:  joeCool,
	}, {
		Name: "Smiling Snoopy",
		Art:  smilingSnoopy,
	},
}

func ArtPrint(art string) {
	lines := strings.Split(art, "\n")
	fmt.Print("\033[H\033[2J")
	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(80 * time.Millisecond)
	}
}
