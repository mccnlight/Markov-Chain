package generator

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// GenerateText generates text using the Markov Chain model.
// It selects words based on prefix mapping and prints the generated output.
func GenerateText(words []string, words_map map[string][]string, word_count int, prefix_length int, starting_prefix string) {
	starting_prefix_slice := strings.Fields(starting_prefix)
	count := len(starting_prefix_slice)
	current_prefix := SetFirstPrefix(starting_prefix_slice, prefix_length)
	_, suffix_found := words_map[current_prefix]
	if suffix_found {
		fmt.Print(starting_prefix + " ")
	} else {
		fmt.Fprintln(os.Stderr, "Error: suffix for prefix \""+strings.TrimSpace(current_prefix)+"\" not found")
		os.Exit(1)
	}
	// Generate text until reaching the desired word count
	for count < word_count {
		suffixes, ok := words_map[current_prefix]
		if !ok {
			fmt.Println()
			fmt.Fprintln(os.Stderr, "Error: suffix for prefix \""+strings.TrimSpace(current_prefix)+"\" not found")
			os.Exit(1)
		}
		suffix := suffixes[rand.Intn(len(suffixes))] // Randomly select a suffix
		if count != word_count-1 {
			suffix += " "
		}
		fmt.Print(suffix)
		current_prefix = SetNextPrefix(current_prefix, suffix) // Update prefix
		count++
	}
	fmt.Println()
	os.Exit(0)
}

// SetFirstPrefix constructs the initial prefix from the starting words.
func SetFirstPrefix(starting_prefix_slice []string, n int) string {
	result_prefix := ""
	for i := len(starting_prefix_slice) - n; i < len(starting_prefix_slice); i++ {
		result_prefix += starting_prefix_slice[i] + " "
	}
	return result_prefix
}

// SetNextPrefix updates the prefix by removing the first word and adding the new suffix.
func SetNextPrefix(current_prefix string, suffix string) string {
	idx := strings.Index(current_prefix, " ")
	if idx == -1 {
		return suffix
	}
	return current_prefix[idx+1:] + suffix
}
