package graph

import (
	"reflect"
	"sort"
	"testing"
)

// TestNewGraph는 새 그래프 생성을 테스트
func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g == nil {
		t.Error("NewGraph should return non-nil graph")
	}
	if g.adjacencyList == nil {
		t.Error("adjacencyList should be initialized")
	}
	if g.vertices == nil {
		t.Error("vertices should be initialized")
	}
}

// TestAddVertex는 정점 추가를 테스트
func TestAddVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	vertices := g.GetVertices()
	if len(vertices) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(vertices))
	}

	// 중복 추가 테스트
	g.AddVertex(1)
	vertices = g.GetVertices()
	if len(vertices) != 3 {
		t.Errorf("Duplicate vertex should not be added, expected 3, got %d", len(vertices))
	}
}

// TestAddEdge는 간선 추가를 테스트
func TestAddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	neighbors := g.GetNeighbors(1)
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors for vertex 1, got %d", len(neighbors))
	}

	expectedNeighbors := []int{2, 3}
	if !reflect.DeepEqual(neighbors, expectedNeighbors) {
		t.Errorf("Expected neighbors %v, got %v", expectedNeighbors, neighbors)
	}
}

// TestDFS는 재귀 DFS를 테스트
func TestDFS(t *testing.T) {
	// 테스트 케이스 1: 간단한 선형 그래프
	// 1 -> 2 -> 3 -> 4
	t.Run("Linear Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		result := g.DFS(1)
		expected := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFS failed: got %v, want %v", result, expected)
		}
	})

	// 테스트 케이스 2: 분기가 있는 그래프
	//     1
	//    / \
	//   2   3
	//  /
	// 4
	t.Run("Branching Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)

		result := g.DFS(1)
		expected := []int{1, 2, 4, 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFS failed: got %v, want %v", result, expected)
		}
	})

	// 테스트 케이스 3: 사이클이 있는 그래프
	// 1 -> 2 -> 3
	// ^         |
	// |_________|
	t.Run("Cyclic Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 1)

		result := g.DFS(1)
		expected := []int{1, 2, 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFS failed: got %v, want %v", result, expected)
		}
	})

	// 테스트 케이스 4: 단일 노드
	t.Run("Single Node", func(t *testing.T) {
		g := NewGraph()
		g.AddVertex(1)

		result := g.DFS(1)
		expected := []int{1}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFS failed: got %v, want %v", result, expected)
		}
	})

	// 테스트 케이스 5: 복잡한 그래프
	//     1
	//    / \
	//   2   3
	//  / \ / \
	// 4   5   6
	t.Run("Complex Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)
		g.AddEdge(2, 5)
		g.AddEdge(3, 5)
		g.AddEdge(3, 6)

		result := g.DFS(1)
		expected := []int{1, 2, 4, 5, 3, 6}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFS failed: got %v, want %v", result, expected)
		}
	})
}

// TestDFSIterative는 반복적 DFS를 테스트
func TestDFSIterative(t *testing.T) {
	// 테스트 케이스 1: 간단한 선형 그래프
	t.Run("Linear Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		result := g.DFSIterative(1)
		expected := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFSIterative failed: got %v, want %v", result, expected)
		}
	})

	// 테스트 케이스 2: 분기가 있는 그래프
	t.Run("Branching Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)

		result := g.DFSIterative(1)
		expected := []int{1, 2, 4, 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFSIterative failed: got %v, want %v", result, expected)
		}
	})

	// 재귀와 반복 버전이 같은 결과를 내는지 확인
	t.Run("Recursive vs Iterative", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)
		g.AddEdge(2, 5)
		g.AddEdge(3, 6)

		recursive := g.DFS(1)
		iterative := g.DFSIterative(1)

		if !reflect.DeepEqual(recursive, iterative) {
			t.Errorf("Recursive and Iterative DFS should produce same result: recursive=%v, iterative=%v", recursive, iterative)
		}
	})
}

// TestDFSAll은 연결되지 않은 그래프에서 모든 정점 방문을 테스트
func TestDFSAll(t *testing.T) {
	// 연결되지 않은 두 개의 컴포넌트
	// 1 -> 2    3 -> 4
	t.Run("Disconnected Graph", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(3, 4)

		result := g.DFSAll()

		// 순서는 보장되지 않으므로 정렬해서 비교
		sort.Ints(result)
		expected := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFSAll failed: got %v, want %v", result, expected)
		}
	})

	// 세 개의 컴포넌트
	t.Run("Three Components", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(1, 2)
		g.AddEdge(3, 4)
		g.AddVertex(5)

		result := g.DFSAll()
		sort.Ints(result)
		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFSAll failed: got %v, want %v", result, expected)
		}
	})
}

// TestHasPath는 두 정점 사이의 경로 존재 여부를 테스트
func TestHasPath(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)

	tests := []struct {
		name     string
		start    int
		end      int
		expected bool
	}{
		{"Direct path exists", 1, 2, true},
		{"Indirect path exists", 1, 4, true},
		{"Path to another branch", 1, 5, true},
		{"No path exists", 4, 1, false},
		{"No path to isolated node", 1, 6, false},
		{"Same node", 1, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.HasPath(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("HasPath(%d, %d) = %v, want %v", tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

// TestAddUndirectedEdge는 무방향 간선 추가를 테스트
func TestAddUndirectedEdge(t *testing.T) {
	g := NewGraph()
	g.AddUndirectedEdge(1, 2)

	// 양방향으로 간선이 추가되었는지 확인
	neighbors1 := g.GetNeighbors(1)
	neighbors2 := g.GetNeighbors(2)

	if len(neighbors1) != 1 || neighbors1[0] != 2 {
		t.Errorf("Expected neighbor 2 for vertex 1, got %v", neighbors1)
	}

	if len(neighbors2) != 1 || neighbors2[0] != 1 {
		t.Errorf("Expected neighbor 1 for vertex 2, got %v", neighbors2)
	}
}

// BenchmarkDFS는 DFS 성능을 벤치마크
func BenchmarkDFS(b *testing.B) {
	g := NewGraph()
	// 큰 그래프 생성
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.DFS(0)
	}
}

// BenchmarkDFSIterative는 반복적 DFS 성능을 벤치마크
func BenchmarkDFSIterative(b *testing.B) {
	g := NewGraph()
	// 큰 그래프 생성
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.DFSIterative(0)
	}
}
