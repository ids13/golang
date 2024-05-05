package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile pola regex
	re := regexp.MustCompile(`\d+`)

	// Compile pola regex dengan penanganan error
	reWithErr, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Match string
	fmt.Println("MatchString:", re.MatchString("123abc")) // Output: true

	// Find all matches
	fmt.Println("FindAllString:", re.FindAllString("123abc456def789", -1)) // Output: [123 456 789]

	// Find first match
	fmt.Println("FindString:", re.FindString("123abc456def"))        // Output: 123
	fmt.Println("FindString:", reWithErr.FindString("123abc456def")) // Output: 123

	// Replace matches
	fmt.Println("ReplaceAllString:", re.ReplaceAllString("123abc456def", "X")) // Output: XabcXdef
}
