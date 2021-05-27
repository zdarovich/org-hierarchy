package lca

type UpwardTree struct {
	adjacent map[int][]int
	parent   map[int]int
	visited  map[int]bool
}

func New() *UpwardTree {
	t := UpwardTree{
		adjacent: make(map[int][]int),
		parent:   make(map[int]int),
		visited:  make(map[int]bool),
	}
	return &t
}

// Adds edge of 'src' and 'dst' nodes to adjacent map
func (t *UpwardTree) AddEdge(src, dst int) {
	if t.adjacent[src] == nil {
		t.adjacent[src] = []int{}
	}
	t.adjacent[src] = append(t.adjacent[src], dst)

	if t.adjacent[dst] == nil {
		t.adjacent[dst] = []int{}
	}
	t.adjacent[dst] = append(t.adjacent[dst], src)
}

// Get parent of every node
func (t *UpwardTree) GetParent(node, parent int) {
	for _, v := range t.adjacent[node] {
		if v == parent {
			continue
		}
		t.parent[v] = node
		t.GetParent(v, node)
	}
}

// finds nearest common manager between 'a' and 'b' employees by traversing to 'root' node
func (t *UpwardTree) FindCommonManager(root, a, b int) *int {
	lca := root

	//Traverse from a upto the root
	for {
		t.visited[a] = true
		if a == root {
			break
		}
		a = t.parent[a]
	}
	//Traverse from 'b' up to the root node. If along the path a visited parent is found, it is the LCA of (a, b)
	for {
		if t.visited[b] {
			lca = b
			break
		}
		b = t.parent[b]
	}

	return &lca
}
