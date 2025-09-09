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
	total := compare_lists(list1, list2)
	fmt.Println(total)
	fmt.Println(get_match_score(list1, list2))
}

// pull the lists from the file
func pull_list() ([]int, []int, error) {
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

	var list1, list2 []int

	// insert each number in sorted order
	for _, line := range lines {
		parts := strings.Fields(line)
		first, second := parts[0], parts[1]
		num1, err := strconv.Atoi(first)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse first number: %w", err)
		} else {
			idx := sort.Search(len(list1), func(i int) bool { return list1[i] >= int(num1) })
			list1 = append(list1, 0)
			copy(list1[idx+1:], list1[idx:])
			list1[idx] = int(num1)
		}
		num2, err := strconv.Atoi(second)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse second number: %w", err)
		} else {
			idx := sort.Search(len(list2), func(i int) bool { return list2[i] >= int(num2) })
			list2 = append(list2, 0)
			copy(list2[idx+1:], list2[idx:])
			list2[idx] = int(num2)
		}
	}
	return list1, list2, nil
}

// go through the lists and compare the indexes, tracking the difference total
func compare_lists(list1, list2 []int) int {
	var total int
	for i := 0; i < len(list1); i++ {
		total += abs(list1[i] - list2[i])
	}
	return total
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func get_match_score(list1, list2 []int) int {
	total := int(0)
	for _, left := range list1 {
		matches := 0
		// use bsearch to find the first instance of left in list2
		idx := sort.Search(len(list2), func(i int) bool { return list2[i] >= left })
		// check the left and right of that val
		for i := idx; i < len(list2) && list2[i] == left; i++ {
			matches++
		}
		matches *= left
		total += matches
	}
	return total
}
