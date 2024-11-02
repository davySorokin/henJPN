package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read number of test cases
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	results := make([]int, n)

	// Process each test case
	for i := 0; i < n; i++ {
		scanner.Scan() // Read and ignore X
		// Read the space-separated integers
		scanner.Scan()
		numbers := strings.Fields(scanner.Text())
		// Calculate the sum of squares of non-negative integers
		sum := 0
		for _, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			if num >= 0 {
				sum += num * num
			}
		}
		results[i] = sum
	}

	// Print all results without blank lines in between
	for _, result := range results {
		fmt.Println(result)
	}
}
