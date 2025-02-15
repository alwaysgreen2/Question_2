package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// waveRearrange rearranges an array into a wave pattern on a per-block basis.
// Each block is of size 2*x+1. In each block:
//   - The element at index x is the maximum of that block.
//   - The elements from indices 0 to x-1 are in ascending order.
//   - The elements from indices x+1 to 2*x are in descending order.
func waveRearrange(arr []int, x int) []int {
	n := len(arr)
	blockSize := 2*x + 1
	result := make([]int, n)

	for i := 0; i < n; i += blockSize {
		// Copy current block.
		block := make([]int, blockSize)
		copy(block, arr[i:i+blockSize])

		// Sort the block in ascending order.
		sort.Ints(block)

		// The maximum element (for the wave peak) is at the end.
		maxVal := block[blockSize-1]
		rem := block[:blockSize-1] // remaining 2*x elements

		leftGroup := []int{}
		rightGroup := []int{}
		// Partition the remaining elements:
		// even-indexed elements go to leftGroup and odd-indexed to rightGroup.
		for j, v := range rem {
			if j%2 == 0 {
				leftGroup = append(leftGroup, v)
			} else {
				rightGroup = append(rightGroup, v)
			}
		}

		// Reverse rightGroup so it becomes non-increasing.
		for i, j := 0, len(rightGroup)-1; i < j; i, j = i+1, j-1 {
			rightGroup[i], rightGroup[j] = rightGroup[j], rightGroup[i]
		}

		// Combine leftGroup, the maximum element, and rightGroup.
		finalBlock := append(leftGroup, maxVal)
		finalBlock = append(finalBlock, rightGroup...)

		// Place the rearranged block in the result.
		copy(result[i:i+blockSize], finalBlock)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt for array input.
		fmt.Println("Enter array of integers separated by commas (or type 'q' to quit):")
		if !scanner.Scan() {
			break
		}
		arrayInput := strings.TrimSpace(scanner.Text())
		if arrayInput == "q" {
			break
		}
		parts := strings.Split(arrayInput, ",")
		var arr []int
		for _, part := range parts {
			numStr := strings.TrimSpace(part)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Invalid integer: %s\n", numStr)
				continue
			}
			arr = append(arr, num)
		}

		// Prompt for x value.
		fmt.Println("Enter integer x (x >= 1) (or type 'q' to quit):")
		if !scanner.Scan() {
			break
		}
		xInput := strings.TrimSpace(scanner.Text())
		if xInput == "q" {
			break
		}
		x, err := strconv.Atoi(xInput)
		if err != nil || x < 1 {
			fmt.Println("Invalid value for x. Must be an integer >= 1.")
			continue
		}

		blockSize := 2*x + 1
		if len(arr)%blockSize != 0 {
			fmt.Printf("Error: array length (%d) is not a multiple of (2*x+1) = %d.\n", len(arr), blockSize)
			continue
		}

		result := waveRearrange(arr, x)
		fmt.Println("Resulting array in wave pattern:")
		fmt.Println(result)
		fmt.Println()
	}
}
