package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLists(path string) ([]int, map[int]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("The file was not found")
	}
	defer file.Close()

	l := []int{}
	r := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		l = append(l, a)
		r[b]++
	}

	return l, r, nil
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func main() {
	l, r, err := getLists("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := 0
	for _, n := range l {
		sum += intAbs(n * r[n])
	}

	fmt.Println("The sum is:", sum)
}
