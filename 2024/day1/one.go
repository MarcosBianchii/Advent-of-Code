package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Input file does not exist")
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

	sort.Ints(l)
	sort.Ints(r)

	sum := 0
	for i := range len(l) {
		sum += int(math.Abs(float64(l[i] - r[i])))
	}

	fmt.Println("The total sum is:", sum)
}
