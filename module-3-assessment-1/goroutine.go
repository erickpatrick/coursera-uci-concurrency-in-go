package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Split(slice []string) [][]string {
	var split [][]string
	// var final [][]string
	midPoint := len(slice) / 2

	split = append(split, slice[0:midPoint])
	split = append(split, slice[midPoint:])

	return split
}

func SortList(slice []string, c chan []int) {
	fmt.Println(slice)
	var numbers []int

	for _, part := range slice {
		number, err := strconv.Atoi(part)

		if err == nil {
			numbers = append(numbers, number)
		}
	}
	sort.Ints(numbers)

	c <- numbers
}

func main() {
	fmt.Println("Please, type a series of numbers separated by space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	split := Split(input)

	var final [][]string

	for i := 0; i < len(split); i += 1 {
		final = append(final, Split(split[i])...)
	}

	// sort in different go routines
	c := make(chan []int, 4)
	go SortList(final[0], c)
	go SortList(final[1], c)
	go SortList(final[2], c)
	go SortList(final[3], c)

	sortedA := <-c
	sortedB := <-c
	sortedC := <-c
	sortedD := <-c

	sortedFirstHalf := append(sortedA, sortedB...)
	sortedSecondHalf := append(sortedC, sortedD...)
	sorted := append(sortedFirstHalf, sortedSecondHalf...)

	sort.Ints(sorted)
	fmt.Println(sorted)
}
