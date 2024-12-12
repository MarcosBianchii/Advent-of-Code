package main

import (
	"bufio"
	"day5/digraph"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getOrdersAndPaths(path string) (map[int][]int, [][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("the input file was not found")
	}
	defer file.Close()

	orders := map[int][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\n")
		if len(line) == 0 {
			break
		}

		nums := strings.Split(line, "|")
		v, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])

		list, _ := orders[v]
		orders[v] = append(list, w)
	}

	var paths [][]int
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\n")
		nums := strings.Split(line, ",")
		var path []int
		for _, n := range nums {
			v, _ := strconv.Atoi(n)
			path = append(path, v)
		}

		paths = append(paths, path)
	}

	return orders, paths, nil
}

func topologicalSort(orders map[int][]int, path []int) ([]int, bool) {
	d := digraph.NewDigraph[int]()
	indegs := map[int]int{}
	for _, v := range path {
		indegs[v] = 0
		d.Vertex(v)
	}

	for _, v := range path {
		for _, w := range orders[v] {
			if d.ContainsVertex(w) {
				d.Arc(v, w)
				indegs[w]++
			}
		}
	}

	var zero []int
	for v, indeg := range indegs {
		if indeg == 0 {
			zero = append(zero, v)
		}
	}

	var newPath []int
	for len(zero) > 0 {
		v := zero[0]
		zero = zero[1:]
		newPath = append(newPath, v)

		for _, w := range d.Adjacents(v) {
			indegs[w]--
			if indegs[w] == 0 {
				zero = append(zero, w)
			}
		}
	}

	return newPath, len(newPath) == d.Len()
}

func isValidPath(orders map[int][]int, path []int) ([]int, bool) {
	for i := range path {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(orders[path[i]], path[j]) {
				return topologicalSort(orders, path)
			}
		}
	}

	return path, false
}

func main() {
	orders, paths, err := getOrdersAndPaths("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sum int
	for _, path := range paths {
		if newPath, ok := isValidPath(orders, path); ok {
			mid := len(newPath) / 2
			sum += newPath[mid]
		}
	}

	fmt.Println("The sum of all the middle values in each valid path is", sum)
}
