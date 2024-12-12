package digraph

import (
	"fmt"
)

type Digraph[T comparable] struct {
	vs map[T]map[T]int
}

func NewDigraph[T comparable]() Digraph[T] {
	return Digraph[T]{
		vs: make(map[T]map[T]int),
	}
}

func NewWithVertices[T comparable](vs []T) Digraph[T] {
	table := make(map[T]map[T]int)

	for _, v := range vs {
		table[v] = make(map[T]int)
	}

	return Digraph[T]{
		vs: table,
	}
}

func (g *Digraph[T]) Vertex(v T) error {
	if g.ContainsVertex(v) {
		return fmt.Errorf("the vertex is already in the Digraph")
	}

	g.vs[v] = make(map[T]int)
	return nil
}

func (g *Digraph[T]) Arc(v, w T) error {
	if g.ContainsArc(v, w) {
		return fmt.Errorf("the arc is already in the Digraph")
	}

	g.vs[v][w] = 1
	return nil
}

func (g *Digraph[T]) RemoveVertex(v T) error {
	if !g.ContainsVertex(v) {
		return fmt.Errorf("the vertex is not in the Digraph")
	}

	for _, w := range g.Adjacents(v) {
		g.Removearc(v, w)
	}

	for w := range g.vs[v] {
		delete(g.vs[w], v)
	}

	delete(g.vs, v)
	return nil
}

func (g *Digraph[T]) Removearc(v, w T) error {
	if !g.ContainsArc(v, w) {
		return fmt.Errorf("the arc is not in the Digraph")
	}

	delete(g.vs[v], w)
	delete(g.vs[w], v)
	return nil
}

func (g Digraph[T]) Degree(v T) int {
	if !g.ContainsVertex(v) {
		return -1
	}

	return len(g.vs[v])
}

func (g Digraph[T]) Adjacents(v T) []T {
	vs := make([]T, 0, g.Degree(v))

	if !g.ContainsVertex(v) {
		return vs
	}

	for w := range g.vs[v] {
		vs = append(vs, w)
	}

	return vs
}

func (g Digraph[T]) Vertices() []T {
	vs := make([]T, 0, g.Len())

	for v := range g.vs {
		vs = append(vs, v)
	}

	return vs
}

func (g Digraph[T]) Weight(v, w T) (int, error) {
	if !g.ContainsArc(v, w) {
		return 0, fmt.Errorf("the arc is not in the Digraph")
	}

	return g.vs[v][w], nil
}

func (g Digraph[T]) Len() int {
	return len(g.vs)
}

func (g *Digraph[T]) ContainsVertex(v T) bool {
	_, ok := g.vs[v]
	return ok
}

func (g *Digraph[T]) ContainsArc(v, w T) bool {
	_, ok := g.vs[v][w]
	return ok
}
