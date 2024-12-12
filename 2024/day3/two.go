package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatch(reader *bufio.Reader, match string) int {
	var i int
	for i < len(match) {
		b, err := reader.ReadByte()
		if err != nil {
			return -1
		}

		if match[i] == b {
			i++
		} else {
			reader.UnreadByte()
			return 1
		}
	}

	return 0
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

func readMatches(reader *bufio.Reader, matches []string) int {
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return -1
		}

		for i, match := range matches {
			if b == match[0] {
				reader.UnreadByte()
				status := readMatch(reader, match)
				if status == -1 {
					return -1
				} else if status == 0 {
					return i
				}
			}
		}
	}
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

	for {
		// do
		matched := readMatches(reader, []string{"mul(", "don't()"})
		if matched == -1 {
			break
		}

		if matched == 1 {
			// don't
			matched = readMatches(reader, []string{"do()"})
			if matched == -1 {
				break
			}
		} else {
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
	}

	fmt.Println("The sum is:", sum)
}
