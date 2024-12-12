package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"day5/digraph"
)

func createDigraph(path string) (digraph.Digraph[int], [][]int, error) {
	d := digraph.NewDigraph[int]()

	file, err := os.Open(path)
	if err != nil {
		return d, nil, fmt.Errorf("the input file was not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\n")
		if len(line) == 0 {
			break
		}

		nums := strings.Split(line, "|")
		v, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])

		if !d.ContainsVertex(v) {
			d.Vertex(v)
		}

		if !d.ContainsVertex(w) {
			d.Vertex(w)
		}

		if !d.ContainsArc(v, w) {
			d.Arc(v, w)
		}
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

	return d, paths, nil
}

func isValidPath(d digraph.Digraph[int], path []int) bool {
	for i := range path {
		for j := i - 1; j >= 0; j-- {
			if d.ContainsArc(path[i], path[j]) {
				return false
			}
		}
	}

	return true
}

func main() {
	d, paths, err := createDigraph("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sum int
	for _, path := range paths {
		if isValidPath(d, path) {
			mid := len(path) / 2
			sum += path[mid]
		}
	}

	fmt.Println("The sum of all the middle values in each valid path is", sum)
}
