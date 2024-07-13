package xgraph

import (
	"fmt"
)

/* 基于邻接矩阵实现的无向图类 */
type graphAdjMat struct {
	// 顶点列表，元素代表“顶点值”，索引代表“顶点索引”
	vertices []int
	// 邻接矩阵，行列索引对应“顶点索引”
	adjMat [][]int
}

// 初始化
func newGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := 0; i < n; i++ {
		adjMat[i] = make([]int, n)
	}
	// 初始化图
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}
	// 添加边
	// 请注意，edges 元素代表顶点索引，即对应 vertices 元素索引
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g
}

/* 获取顶点数量 */
func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

// 参数 i, j 对应 vertices 元素索引
func (g *graphAdjMat) addEdge(i, j int) {
	// 索引越界与相等处理
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index Out Of Bounds Exception")
	}

	// 在无向图中，邻接矩阵关于主对角线对称，即满足 (i, j) == (j, i)
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

// 添加顶点
func (g *graphAdjMat) addVertex(val int) {
	n := g.size()
	// 向顶点列表中添加新顶点的值
	g.vertices = append(g.vertices, val)
	// 在邻接矩阵中添加一行
	newrow := make([]int, n)
	g.adjMat = append(g.adjMat, newrow)
	// 在邻接矩阵中添加一列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}
