# Graph - DFS (Depth-First Search) 구현

깊이 우선 탐색(DFS) 알고리즘의 Go 구현입니다.

## 특징

- **인접 리스트 기반 그래프**: 메모리 효율적인 그래프 표현
- **재귀 DFS**: 간결하고 이해하기 쉬운 재귀적 구현
- **반복 DFS**: 스택을 사용한 명시적 반복 구현
- **경로 찾기**: 두 정점 사이의 경로 존재 여부 확인
- **연결되지 않은 그래프 지원**: 모든 컴포넌트를 탐색하는 DFSAll
- **방향/무방향 그래프 지원**: 유연한 그래프 구조

## 사용법

### 기본 사용 예제

```go
package main

import (
    "fmt"
    "go-lab/algorithm/graph"
)

func main() {
    // 새 그래프 생성
    g := graph.NewGraph()

    // 간선 추가 (방향 그래프)
    g.AddEdge(1, 2)
    g.AddEdge(1, 3)
    g.AddEdge(2, 4)
    g.AddEdge(3, 5)

    // 재귀 DFS 실행
    result := g.DFS(1)
    fmt.Println("DFS 탐색 결과:", result)
    // 출력: DFS 탐색 결과: [1 2 4 3 5]

    // 반복적 DFS 실행
    resultIterative := g.DFSIterative(1)
    fmt.Println("DFS Iterative 결과:", resultIterative)
    // 출력: DFS Iterative 결과: [1 2 4 3 5]
}
```

### 무방향 그래프

```go
g := graph.NewGraph()

// 무방향 간선 추가
g.AddUndirectedEdge(1, 2)
g.AddUndirectedEdge(2, 3)
g.AddUndirectedEdge(3, 4)

result := g.DFS(1)
fmt.Println(result) // [1 2 3 4]
```

### 경로 찾기

```go
g := graph.NewGraph()
g.AddEdge(1, 2)
g.AddEdge(2, 3)
g.AddEdge(3, 4)

// 1에서 4로 가는 경로가 있는지 확인
hasPath := g.HasPath(1, 4)
fmt.Println("경로 존재:", hasPath) // true

// 4에서 1로 가는 경로가 있는지 확인 (방향 그래프)
hasPath = g.HasPath(4, 1)
fmt.Println("경로 존재:", hasPath) // false
```

### 연결되지 않은 그래프

```go
g := graph.NewGraph()

// 두 개의 분리된 컴포넌트
g.AddEdge(1, 2)
g.AddEdge(3, 4)

// 모든 정점 탐색
result := g.DFSAll()
fmt.Println(result) // [1 2 3 4] (순서는 다를 수 있음)
```

## API 문서

### Graph 구조체

```go
type Graph struct {
    adjacencyList map[int][]int
    vertices      map[int]bool
}
```

### 주요 메서드

#### NewGraph()
```go
func NewGraph() *Graph
```
새로운 그래프 인스턴스를 생성합니다.

#### AddVertex(v int)
```go
func (g *Graph) AddVertex(v int)
```
그래프에 정점을 추가합니다.

#### AddEdge(from, to int)
```go
func (g *Graph) AddEdge(from, to int)
```
방향 간선을 추가합니다.

#### AddUndirectedEdge(v1, v2 int)
```go
func (g *Graph) AddUndirectedEdge(v1, v2 int)
```
무방향 간선을 추가합니다 (양방향).

#### DFS(start int) []int
```go
func (g *Graph) DFS(start int) []int
```
재귀를 사용한 깊이 우선 탐색을 수행합니다.
- **매개변수**: start - 시작 정점
- **반환값**: 방문 순서대로 정점들의 슬라이스

#### DFSIterative(start int) []int
```go
func (g *Graph) DFSIterative(start int) []int
```
스택을 사용한 반복적 깊이 우선 탐색을 수행합니다.
- **매개변수**: start - 시작 정점
- **반환값**: 방문 순서대로 정점들의 슬라이스

#### DFSAll() []int
```go
func (g *Graph) DFSAll() []int
```
모든 정점을 방문하는 DFS를 수행합니다 (연결되지 않은 그래프 처리).
- **반환값**: 방문 순서대로 정점들의 슬라이스

#### HasPath(start, end int) bool
```go
func (g *Graph) HasPath(start, end int) bool
```
두 정점 사이에 경로가 존재하는지 확인합니다.
- **매개변수**:
  - start - 시작 정점
  - end - 목적지 정점
- **반환값**: 경로가 존재하면 true

## 시간 복잡도

- **DFS**: O(V + E)
  - V: 정점의 수
  - E: 간선의 수
- **공간 복잡도**: O(V) - 방문 배열 및 재귀 스택

## 테스트

```bash
# 모든 테스트 실행
go test ./algorithm/graph/... -v

# 벤치마크 실행
go test ./algorithm/graph/... -bench=. -benchmem

# 커버리지 확인
go test ./algorithm/graph/... -cover
```

## 테스트 커버리지

구현은 다음을 포함한 포괄적인 테스트로 검증되었습니다:
- 선형 그래프 탐색
- 분기가 있는 그래프 탐색
- 사이클이 있는 그래프 처리
- 단일 노드 처리
- 복잡한 그래프 구조
- 연결되지 않은 그래프 처리
- 재귀 vs 반복 구현 비교
- 경로 찾기 알고리즘
- 성능 벤치마크

## 실제 활용 사례

1. **미로 찾기**: 미로의 출구를 찾는 알고리즘
2. **네트워크 연결성**: 네트워크에서 두 노드가 연결되어 있는지 확인
3. **위상 정렬**: 작업 스케줄링 및 의존성 해결
4. **사이클 감지**: 그래프에서 순환 의존성 확인
5. **경로 찾기**: 두 지점 사이의 경로 탐색
