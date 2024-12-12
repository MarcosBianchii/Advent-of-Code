package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getTable(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][]byte{}, fmt.Errorf("The input file was not found")
	}
	defer file.Close()

	var table [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\n")
		table = append(table, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return table, nil
}

func inRange(i, j, n, m int) bool {
	return 0 <= i && 0 <= j && i < n && j < m
}

func checkMatch(i, j, p, q int, table [][]byte, match string) uint {
	n := len(table)
	m := len(table[0])
	for k := 0; k < len(match); k++ {
		if !inRange(i, j, n, m) || match[k] != table[i][j] {
			return 0
		}

		i += p
		j += q
	}

	return 1
}

func find(i, j int, table [][]byte, match string) uint {
	n := len(table)
	m := len(table[0])
	if j == m {
		i++
		j = 0
	}

	if !inRange(i, j, n, m) {
		return 0
	}

	var count uint
	count += checkMatch(i, j, -1, -1, table, match)
	count += checkMatch(i, j, -1, 0, table, match)
	count += checkMatch(i, j, -1, 1, table, match)
	count += checkMatch(i, j, 0, -1, table, match)
	count += checkMatch(i, j, 0, 1, table, match)
	count += checkMatch(i, j, 1, -1, table, match)
	count += checkMatch(i, j, 1, 1, table, match)
	count += checkMatch(i, j, 1, 0, table, match)
	return find(i, j+1, table, match) + count
}

func main() {
	table, err := getTable("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	match := "XMAS"
	count := find(0, 0, table, match)
	fmt.Println("The word", match, "was found", count, "times")
}
