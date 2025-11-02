package graph

// Graph는 인접 리스트를 사용한 그래프 구조체
type Graph struct {
	// adjacencyList는 각 정점에서 연결된 정점들의 목록을 저장
	adjacencyList map[int][]int
	// vertices는 그래프의 모든 정점을 저장
	vertices map[int]bool
}

// NewGraph는 새로운 그래프를 생성
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
		vertices:      make(map[int]bool),
	}
}

// AddVertex는 그래프에 정점을 추가
func (g *Graph) AddVertex(v int) {
	if !g.vertices[v] {
		g.vertices[v] = true
		g.adjacencyList[v] = []int{}
	}
}

// AddEdge는 그래프에 간선을 추가 (방향 그래프)
func (g *Graph) AddEdge(from, to int) {
	// 정점이 없으면 추가
	g.AddVertex(from)
	g.AddVertex(to)

	// 간선 추가
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
}

// AddUndirectedEdge는 무방향 간선을 추가
func (g *Graph) AddUndirectedEdge(v1, v2 int) {
	g.AddEdge(v1, v2)
	g.AddEdge(v2, v1)
}

// DFS는 재귀를 사용한 깊이 우선 탐색
// start: 시작 정점
// 반환값: 방문 순서대로 정점들의 슬라이스
func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{}

	g.dfsRecursive(start, visited, &result)

	return result
}

// dfsRecursive는 DFS의 재귀 헬퍼 함수
func (g *Graph) dfsRecursive(vertex int, visited map[int]bool, result *[]int) {
	// 이미 방문한 정점이면 리턴
	if visited[vertex] {
		return
	}

	// 현재 정점 방문 처리
	visited[vertex] = true
	*result = append(*result, vertex)

	// 인접한 정점들을 재귀적으로 방문
	for _, neighbor := range g.adjacencyList[vertex] {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited, result)
		}
	}
}

// DFSIterative는 스택을 사용한 반복적 깊이 우선 탐색
// start: 시작 정점
// 반환값: 방문 순서대로 정점들의 슬라이스
func (g *Graph) DFSIterative(start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	stack := []int{start}

	for len(stack) > 0 {
		// 스택에서 정점 꺼내기
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 이미 방문한 정점이면 스킵
		if visited[vertex] {
			continue
		}

		// 현재 정점 방문 처리
		visited[vertex] = true
		result = append(result, vertex)

		// 인접한 정점들을 스택에 추가 (역순으로 추가하여 작은 번호부터 방문)
		neighbors := g.adjacencyList[vertex]
		for i := len(neighbors) - 1; i >= 0; i-- {
			if !visited[neighbors[i]] {
				stack = append(stack, neighbors[i])
			}
		}
	}

	return result
}

// DFSAll은 모든 정점을 방문하는 DFS (연결되지 않은 그래프 처리)
// 반환값: 방문 순서대로 정점들의 슬라이스
func (g *Graph) DFSAll() []int {
	visited := make(map[int]bool)
	result := []int{}

	// 모든 정점에 대해 DFS 수행
	for vertex := range g.vertices {
		if !visited[vertex] {
			g.dfsRecursive(vertex, visited, &result)
		}
	}

	return result
}

// HasPath는 두 정점 사이에 경로가 존재하는지 확인
func (g *Graph) HasPath(start, end int) bool {
	visited := make(map[int]bool)
	return g.hasPathRecursive(start, end, visited)
}

// hasPathRecursive는 경로 존재 여부를 확인하는 재귀 헬퍼 함수
func (g *Graph) hasPathRecursive(current, end int, visited map[int]bool) bool {
	// 목적지에 도달하면 true
	if current == end {
		return true
	}

	// 이미 방문한 정점이면 false
	if visited[current] {
		return false
	}

	// 현재 정점 방문 처리
	visited[current] = true

	// 인접한 정점들을 재귀적으로 탐색
	for _, neighbor := range g.adjacencyList[current] {
		if g.hasPathRecursive(neighbor, end, visited) {
			return true
		}
	}

	return false
}

// GetVertices는 그래프의 모든 정점을 반환
func (g *Graph) GetVertices() []int {
	vertices := make([]int, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

// GetNeighbors는 특정 정점의 인접 정점들을 반환
func (g *Graph) GetNeighbors(v int) []int {
	return g.adjacencyList[v]
}
