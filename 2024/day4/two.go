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

func isXmas(table [][]byte, i, j int) bool {
	if table[i][j] != 'A' {
		return false
	}

	if table[i-1][j-1] == 'M' && table[i-1][j+1] == 'M' && table[i+1][j-1] == 'S' && table[i+1][j+1] == 'S' {
		return true
	}

	if table[i-1][j+1] == 'M' && table[i+1][j+1] == 'M' && table[i-1][j-1] == 'S' && table[i+1][j-1] == 'S' {
		return true
	}

	if table[i+1][j-1] == 'M' && table[i+1][j+1] == 'M' && table[i-1][j-1] == 'S' && table[i-1][j+1] == 'S' {
		return true
	}

	if table[i-1][j-1] == 'M' && table[i+1][j-1] == 'M' && table[i-1][j+1] == 'S' && table[i+1][j+1] == 'S' {
		return true
	}

	return false
}

func main() {
	table, err := getTable("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n := len(table)
	m := len(table[0])

	var count uint
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if isXmas(table, i, j) {
				count++
			}
		}
	}

	fmt.Println("There are", count, "X-mas")
}
