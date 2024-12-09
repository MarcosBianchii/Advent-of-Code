package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func tryMakingSafe(report []int, retries int) bool {
	for i := range report {
		reportCopy := make([]int, len(report))
		copy(reportCopy, report)
		if isSafe(append(reportCopy[:i], reportCopy[i+1:]...), retries+1) {
			return true
		}
	}

	return false
}

func isSafe(report []int, retries int) bool {
	if retries > 1 {
		return false
	}

	initDif := report[1] - report[0]

	for i := range len(report) - 1 {
		dif := report[i+1] - report[i]
		if initDif*dif <= 0 || intAbs(dif) > 3 {
			if !tryMakingSafe(report, retries) {
				return false
			}
		}
	}

	return true
}

func getReports(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("The file was not found")
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportStr := strings.Fields(scanner.Text())
		report := make([]int, 0, len(reportStr))

		for i := range reportStr {
			n, _ := strconv.Atoi(reportStr[i])
			report = append(report, n)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func main() {
	reports, err := getReports("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report, 0) {
			safeCount++
		}
	}

	fmt.Println("There are", safeCount, "safe reports")
}
