package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"markov-chain/generator"
	"markov-chain/utilities"
)

func main() {
	args := os.Args[1:] // Retrieve command-line arguments

	utilities.CheckHelpFlag(args) // Check if help flag is present and display help if needed

	words := utilities.ReadFile() // Read input text from stdin

	// Default parameters
	word_count := 100
	prefix_length := 2
	starting_prefix := ""
	custom_starting_prefix := false

	utilities.CheckTextLength(words, prefix_length) // Ensure input has enough words

	utilities.CheckLastArgument(args) // Ensure the last argument is valid

	// Parse command-line arguments
	for i := 0; i < len(args); i++ {
		if (args[i] == "-w" || args[i] == "--w") && i != len(args)-1 {
			// Parse word count argument
			val, err := strconv.Atoi(args[i+1])
			if err == nil {
				if val <= 0 || val > 10000 {
					fmt.Fprintln(os.Stderr, "Error: number of words must be in between 1 and 10000")
					os.Exit(1)
				}
				word_count = val
			} else {
				fmt.Fprintln(os.Stderr, "Error: invalid value for the number of maximum words")
				os.Exit(1)
			}
			i++
		} else if (args[i] == "-p" || args[i] == "--p") && i != len(args)-1 {
			// Parse starting prefix argument
			if !utilities.CheckForFlags(args[i+1]) {
				starting_prefix = args[i+1]
				custom_starting_prefix = true
			} else {
				fmt.Fprintln(os.Stderr, "Error: invalid value for the starting prefix")
				os.Exit(1)
			}
			i++
		} else if (args[i] == "-l" || args[i] == "--l") && i != len(args)-1 {
			// Parse prefix length argument
			val, err := strconv.Atoi(args[i+1])
			if err == nil {
				if val <= 0 || val > 5 {
					fmt.Fprintln(os.Stderr, "Error: prefix length must be in between 1 and 5")
					os.Exit(1)
				}
				prefix_length = val
			} else {
				fmt.Fprintln(os.Stderr, "Error: invalid value for the prefix length")
				os.Exit(1)
			}
			i++
		} else {
			fmt.Fprintln(os.Stderr, "Error: unknown argument \""+args[i]+"\"")
			os.Exit(1)
		}
	}

	// Validate and set the starting prefix
	if custom_starting_prefix {
		if strings.ReplaceAll(starting_prefix, " ", "") == "" {
			fmt.Fprintln(os.Stderr, "Error: starting prefix is empty")
			os.Exit(1)
		}
	} else {
		starting_prefix = utilities.BuildStartingPrefix(prefix_length, words)
	}

	// Validate prefix and word count constraints
	utilities.CheckPrefixLength(prefix_length, word_count)
	utilities.CheckStartingPrefixLength(starting_prefix, prefix_length)

	// Build Markov chain model
	words_map := utilities.BuildMap(words, prefix_length)

	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	// Generate text based on the model
	generator.GenerateText(words, words_map, word_count, prefix_length, starting_prefix)
}
