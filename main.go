package main

import "fmt"

func main() {
	n := 50
	nodes.calcPatterns(n)
	fmt.Printf("n=%d\n", n)
	for _, node := range nodes {
		fmt.Printf("%d can be dial in %d patterns\n", node.key, node.patterns)
	}
	// fmt.Println(pc)
	fmt.Printf("Cache: length %d, hits %d\n", len(pc), cacheHit)
}

var nodes Nodes
var pc map[string]int
var cacheHit int

func init() {
	nodes = createUG()
	pc = make(map[string]int, 0)
}

type (
	Nodes []*Node
	Node  struct {
		key      int
		vectors  []int
		patterns int
	}
)

func createUG() Nodes {
	n := Nodes{
		{key: 0, vectors: []int{4, 6}},
		{key: 1, vectors: []int{6, 8}},
		{key: 2, vectors: []int{7, 9}},
		{key: 3, vectors: []int{4, 8}},
		{key: 4, vectors: []int{0, 3, 9}},
		{key: 5, vectors: []int{}},
		{key: 6, vectors: []int{0, 1, 7}},
		{key: 7, vectors: []int{2, 6}},
		{key: 8, vectors: []int{1, 3}},
		{key: 9, vectors: []int{2, 4}},
	}
	return n
}
func (nodes Nodes) calcPatterns(n int) {
	for _, node := range nodes {
		// fmt.Printf("&&&&& %d\n", node)
		node.patterns = node.calcHelper(n)
		if node.patterns == 0 {
			node.patterns = 1
		}
	}
}

func (node Node) calcHelper(n int) int {
	// fmt.Printf("###%d:%d\n", node.key, n)
	// fmt.Printf("Cache length: %d\n", len(pc))
	var p int
	if v, ok := pc[fmt.Sprintf("%d:%d", node.key, n)]; ok {
		// fmt.Printf("$$$ Cache hit! %d:%d --- %d\n", node.key, n, v)
		cacheHit++
		return v
	}
	if n > 1 {
		tmpN := n - 1
		for _, v := range node.vectors {
			p += nodes[v].calcHelper(tmpN)
		}
	} else if n == 1 {
		p += len(node.vectors)
	}
	pc[fmt.Sprintf("%d:%d", node.key, n)] = p
	if _, ok := pc[fmt.Sprintf("%d:%d", node.key, n-2)]; ok {
		delete(pc, fmt.Sprintf("%d:%d", node.key, n-2))
	}
	// fmt.Printf("@@@@%d:%d --- %d\n", node.key, n, p)
	return p
}
