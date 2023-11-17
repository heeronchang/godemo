package graph_adjacency_matrix

import "fmt"

type GraphAdjMatrix struct {
	vertices  []int
	adjMatrix [][]int
}

func NewGraphAdjMatrix(vertices []int, edges [][]int) *GraphAdjMatrix {
	n := len(vertices)
	adjMatrix := make([][]int, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, n)
	}

	g := &GraphAdjMatrix{
		vertices:  vertices,
		adjMatrix: adjMatrix,
	}

	// 添加边
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func (g *GraphAdjMatrix) size() int {
	return len(g.vertices)
}

func (g *GraphAdjMatrix) AddVertex(val int) {
	n := g.size()
	g.vertices = append(g.vertices, val)

	newRow := make([]int, n)
	g.adjMatrix = append(g.adjMatrix, newRow)
	for i := range g.adjMatrix {
		g.adjMatrix[i] = append(g.adjMatrix[i], 0)
	}
}

func (g *GraphAdjMatrix) RemoveVertex(index int) {
	if index >= g.size() {
		return
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMatrix = append(g.adjMatrix[:index], g.adjMatrix[index+1:]...)
	for i := range g.adjMatrix {
		g.adjMatrix[i] = append(g.adjMatrix[i][:index], g.adjMatrix[i][i+1:]...)
	}
}

func (g *GraphAdjMatrix) AddEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Println(fmt.Errorf("%s\n", "Index out of Bounds Exception"))
		return
	}
	g.adjMatrix[i][j] = 1
	g.adjMatrix[j][i] = 1
}

func (g *GraphAdjMatrix) RemoveEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Println(fmt.Errorf("%s\n", "Index out of Bounds Exception"))
		return
	}
	g.adjMatrix[i][j] = 0
	g.adjMatrix[j][i] = 0
}

func (g *GraphAdjMatrix) Print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t邻接矩阵 = \n")
	for i := range g.adjMatrix {
		fmt.Printf("\t\t\t%v\n", g.adjMatrix[i])
	}
}
