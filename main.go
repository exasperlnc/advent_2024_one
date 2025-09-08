package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2, _ := pull_list()
	sort1, sort2 := sort_lists(list1, list2)
	fmt.Println("Sorted List 1:", sort1)
	fmt.Println("Sorted List 2:", sort2)
}

// pull the lists from the file
func pull_list() ([]int32, []int32, error) {
	// open the file
	file, err := os.Open("list.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	// scan it for lines
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	var list1, list2 []int32

	// append the first half to one list
	for _, line := range lines {
		parts := strings.Fields(line)
		first, second := parts[0], parts[1]
		num1, err := strconv.Atoi(first)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse first number: %w", err)
		} else {
			list1 = append(list1, int32(num1))
		}
		// append the second half to the other list
		num2, err := strconv.Atoi(second)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse second number: %w", err)
		} else {
			list2 = append(list2, int32(num2))
		}
	}
	return list1, list2, nil
}

// sort the lists
func sort_lists(list1, list2 []int32) ([]int32, []int32) {
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})
	return list1, list2
}

// go through the lists and compare the indexes, tracking the difference total 

// return that total 

