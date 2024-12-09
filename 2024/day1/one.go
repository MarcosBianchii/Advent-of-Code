package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getLists(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("The file was not found")
	}
	defer file.Close()

	var l, r []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		l, r = append(l, a), append(r, b)
	}

	return l, r, nil
}

func intDif(a, b int) int {
	dif := a - b
	if dif < 0 {
		return -dif
	}

	return dif
}

func main() {
	l, r, err := getLists("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sort.Ints(l)
	sort.Ints(r)

	sum := 0
	for i := range len(l) {
		sum += intDif(l[i], r[i])
	}

	fmt.Println("The total sum is:", sum)
}
