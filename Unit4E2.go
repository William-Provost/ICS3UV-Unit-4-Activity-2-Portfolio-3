// Author: William
// Version: 1.0.0
// Date: 2025-11-25
// Fileoverview: This program asks for a positive integer and counts down to 0.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var number int
	reader := bufio.NewReader(os.Stdin)

	// input loop (do-while equivalent)
	for {
		fmt.Print("Enter a positive integer: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, _ = strconv.Atoi(input)

		if number > 0 {
			break
		}

		fmt.Println("Invalid input. Please enter a positive integer:")
	}

	// countdown
	fmt.Printf("Counting down from %d:\n", number)

	for number >= 0 {
		fmt.Println(number)
		number--
	}

	fmt.Println("\nDone.")
}
