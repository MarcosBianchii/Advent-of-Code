package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMul(reader *bufio.Reader) bool {
	match := "mul("
	var i int
	for i < 4 {
		b, err := reader.ReadByte()
		if err != nil {
			return true
		}

		if match[i] == b {
			i++
		} else {
			i = 0
		}
	}

	return false
}

func readNumber(reader *bufio.Reader, terminator byte) (int, error) {
	x := strings.Builder{}
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return 0, fmt.Errorf("encountered EOF")
		}

		if b == terminator {
			break
		} else if x.Len() == 3 {
			return 0, fmt.Errorf("invalid token")
		}

		_, err = strconv.ParseInt(string(b), 10, 0)
		if err == nil {
			x.WriteByte(b)
		} else {
			return 0, fmt.Errorf("invalid token")
		}
	}

	if x.Len() == 0 {
		return 0, fmt.Errorf("empty number")
	}

	return strconv.Atoi(x.String())
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file was not found")
		os.Exit(1)
	}
	defer file.Close()

	sum := 0
	reader := bufio.NewReader(file)
	for !readMul(reader) {
		x, err := readNumber(reader, ',')
		if err != nil {
			continue
		}

		y, err := readNumber(reader, ')')
		if err != nil {
			continue
		}

		sum += x * y
	}

	fmt.Println("The sum is:", sum)
}
