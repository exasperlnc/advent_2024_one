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

	// insert each number in sorted order
	for _, line := range lines {
		parts := strings.Fields(line)
		first, second := parts[0], parts[1]
		num1, err := strconv.Atoi(first)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse first number: %w", err)
		} else {
			idx := sort.Search(len(list1), func(i int) bool { return list1[i] >= int32(num1) })
			list1 = append(list1, 0)
			copy(list1[idx+1:], list1[idx:])
			list1[idx] = int32(num1)
		}
		num2, err := strconv.Atoi(second)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse second number: %w", err)
		} else {
			idx := sort.Search(len(list2), func(i int) bool { return list2[i] >= int32(num2) })
			list2 = append(list2, 0)
			copy(list2[idx+1:], list2[idx:])
			list2[idx] = int32(num2)
		}
	}
	return list1, list2, nil
}

// go through the lists and compare the indexes, tracking the difference total
func compare_lists(list1, list2 []int32) int32 {
	var total int32
	for i := 0; i < len(list1); i++ {
		total += absInt32(list1[i] - list2[i])
	}
	return total
}

// absInt32 returns the absolute value of an int32
func absInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
