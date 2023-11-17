package graph_adjacency_matrix

import "testing"

func TestName(t *testing.T) {
	g := NewGraphAdjMatrix([]int{1, 3, 5, 6, 8, 7}, [][]int{[]int{0, 1}, []int{2, 3}})
	g.Print()
}
