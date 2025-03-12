package utilities

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadFile reads text from stdin and returns a list of words.
func ReadFile() []string {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "Error: no input detected, provide input via a file or pipe")
		os.Exit(1)
	}

	words := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
			os.Exit(1)
		}
		if line == "" && err == io.EOF {
			break
		}

		words = append(words, strings.Fields(line)...) // Split the line into words and append to list

		if err == io.EOF {
			break
		}
	}

	return words
}

// ShowHelp prints usage instructions for the program.
func ShowHelp() {
	fmt.Println("Markov Chain text generator.\n")
	fmt.Println("Usage:")
	fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
	fmt.Println("  markovchain --help\n")
	fmt.Println("Options:")
	fmt.Println("  --help  Show this screen.")
	fmt.Println("  -w N    Number of maximum words")
	fmt.Println("  -p S    Starting prefix")
	fmt.Println("  -l N    Prefix length")
}

// CheckHelpFlag checks if help flag is passed and displays help text if found.
func CheckHelpFlag(args []string) {
	for _, arg := range args {
		if arg == "--help" || arg == "-help" {
			ShowHelp()
			os.Exit(0)
		}
	}
}

// CheckTextLength ensures that the input text has enough words to proceed.
func CheckTextLength(words []string, prefix_length int) {
	if len(words) < prefix_length {
		fmt.Fprintln(os.Stderr, "Error: not enough words to generate the text")
		os.Exit(1)
	}
}

// BuildStartingPrefix constructs a starting prefix of given length from input words.
func BuildStartingPrefix(n int, words []string) string {
	var starting_prefix string
	for i := 0; i < n; i++ {
		starting_prefix += words[i]
		if i != n-1 {
			starting_prefix += " "
		}
	}
	return starting_prefix
}

// CheckForFlags checks if the given string is a valid command-line flag.
func CheckForFlags(s string) bool {
	return s == "-w" || s == "-p" || s == "-l" || s == "--help" || s == "--w" || s == "--p" || s == "--l" || s == "-help"
}

// CheckLastArgument ensures that the last argument is not a flag without a value.
func CheckLastArgument(args []string) {
	if len(args) == 0 {
		return
	}
	last_argument := args[len(args)-1]
	if last_argument == "-w" || last_argument == "-p" || last_argument == "-l" {
		fmt.Fprintln(os.Stderr, "Error: a flag cannot be placed as the last argument")
		os.Exit(1)
	}
}

// CheckStartingPrefixLength ensures the starting prefix length is valid.
func CheckStartingPrefixLength(starting_prefix string, prefix_length int) {
	if len(strings.Fields(starting_prefix)) < prefix_length {
		fmt.Fprintln(os.Stderr, "Error: starting prefix length must be equal or more than the prefix length (at least 2 if you didn't specify the prefix length)")
		os.Exit(1)
	}
}

// CheckPrefixLength validates that the prefix length is appropriate for the word count.
func CheckPrefixLength(prefix_length int, word_count int) {
	if prefix_length >= word_count {
		fmt.Fprintln(os.Stderr, "Error: prefix length should be less than word count (more than 2)")
		os.Exit(1)
	}
}

// BuildMap constructs a Markov chain mapping prefixes to possible suffixes.
func BuildMap(words []string, n int) map[string][]string {
	words_map := make(map[string][]string)
	for i := 0; i < len(words)-n; i++ {
		current_prefix := ""
		for k := 0; k < n; k++ {
			current_prefix += words[k+i] + " "
		}
		next_suffix := words[i+n] // Get the next word following the prefix
		words_map[current_prefix] = append(words_map[current_prefix], next_suffix)
	}

	return words_map
}
