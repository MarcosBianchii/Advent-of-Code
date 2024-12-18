package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	i, j int
}

type Guard struct {
	n, m      int
	position  Point
	board     map[Point]struct{}
	path      map[Point]struct{}
	direction Point
}

func newGuard(startingPosition Point, n, m int, board map[Point]struct{}) *Guard {
	return &Guard{
		position:  startingPosition,
		n:         n,
		m:         m,
		board:     board,
		path:      map[Point]struct{}{startingPosition: {}},
		direction: Point{-1, 0},
	}
}

func (g Guard) isInRange() bool {
	i, j := g.position.i, g.position.j
	return 0 <= i && 0 <= j && i < g.n && j < g.m
}

func (g Guard) nextPosition() Point {
	pi, pj := g.position.i, g.position.j
	di, dj := g.direction.i, g.direction.j
	return Point{pi + di, pj + dj}
}

func (g Guard) nextPositionIsObstruction() bool {
	_, ok := g.board[g.nextPosition()]
	return ok
}

func (g *Guard) turnRight() {
	g.direction.i, g.direction.j = g.direction.j, -g.direction.i
}

func (g *Guard) advance() {
	g.path[g.position] = struct{}{}
	g.position = g.nextPosition()
}

func (g Guard) walk() {
	for g.isInRange() {
		if g.nextPositionIsObstruction() {
			g.turnRight()
		} else {
			g.advance()
		}
	}
}

func (g Guard) differentPositions() int {
	return len(g.path)
}

func readInput(path string) (*Guard, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	n, m := 0, 0
	startingRow, startingCol := 0, 0
	board := make(map[Point]struct{})
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, r := range line {
			switch r {
			case '#':
				board[Point{i, j}] = struct{}{}
			case '^':
				startingRow, startingCol = i, j
			}

			m = max(m, j)
		}

		n = max(n, i)
	}

	return newGuard(Point{startingRow, startingCol}, n, m, board), nil
}

func main() {
	guard, err := readInput("input.txt")
	if err != nil {
		log.Fatal("the input file was not found")
	}

	guard.walk()
	fmt.Println("The guard walked a total of", guard.differentPositions(), "different positions")
}
