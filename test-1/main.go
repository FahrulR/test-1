package main

import (
	"bufio"
	f "fmt"
	"os"
	"sort"
	"strings"
)

var output string

func main() {
	var input string
	f.Printf("Input any word: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	splitChar(input)
	f.Printf("Output: %s\n", output)
}

func splitChar(input string) string {
	var vowels []string
	var consonants []string
	var prints []string
	input = strings.ToLower(input)
	input = strings.Replace(input, " ", "", -1)
	word := strings.Split(input, "")
	for _, ch := range word {
		switch ch {
		case "a", "e", "i", "o", "u":
			vowels = append(vowels, ch)
		default:
			consonants = append(consonants, ch)
		}
	}
	sort.Strings(vowels)
	sort.Strings(consonants)
	prints = append(vowels, consonants...)
	output = strings.Join(prints, "")
	return output
}
