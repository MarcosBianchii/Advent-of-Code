package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Input file does not exist")
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

	sum := 0
	for i := range len(l) {
		sum += int(math.Abs(float64(l[i] * r[l[i]])))
	}

	fmt.Println("The sum is:", sum)
}
