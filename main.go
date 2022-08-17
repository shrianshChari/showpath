package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/gookit/color"
	)

// Returns the number of characters that the integer
// takes up when printed out.
// Examples: 10 -> 2, 123 -> 3
func num_chars(val int) int {
	return int(math.Floor(math.Log10(float64(val)))) + 1
}

func main() {
	rainbow_flag := flag.Bool("r", false,
		"Will color the output with the colors of the rainbow.")
	numbered_flag := flag.Bool("n", false,
		"Will number the paths in the output.")
	flag.Parse()

	// Gets the PATH variable
	path := os.Getenv("PATH")

	// Splits the path based on the : character
	path_strings := strings.Split(path, ":")
	if len(path_strings) == 0 {
		fmt.Println("There are no folders in your $PATH!")
	} else {
		
		// The best workaround for dealing with stupidly long $PATHs
		// Remember that "%%" evaluates to just a %
		str_fmt := fmt.Sprintf("%%%dd. ", num_chars(len(path_strings)))
		
		if *rainbow_flag {
			colors := map[int]color.Color {
				0: color.FgRed,
				// There isn't an orange color in the 16 colors, so we skip it
				1: color.FgYellow,
				2: color.FgGreen,
				3: color.FgBlue,
				4: color.FgMagenta,
			}

			num_colors := len(colors)

			for ind, val := range path_strings {
				if *numbered_flag {
					colors[ind % num_colors].Printf(str_fmt, ind + 1)
				}
				colors[ind % num_colors].Println(val)
				// Doing ind % num_colors allows me to cycle through the colors
			}
		} else {
			for ind, val := range path_strings {
				if *numbered_flag {
					fmt.Printf(str_fmt, ind + 1)
				}

				fmt.Println(val)
			}
		}
	}
} // Nice
