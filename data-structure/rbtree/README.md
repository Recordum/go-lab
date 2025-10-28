# Red-Black Tree

Red-Black Tree는 자가 균형 이진 탐색 트리(self-balancing binary search tree)입니다.

## 특징

### Red-Black Tree의 속성

1. 모든 노드는 빨강(Red) 또는 검정(Black)이다
2. 루트 노드는 검정이다
3. 모든 리프(NIL) 노드는 검정이다
4. 빨강 노드의 자식은 모두 검정이다 (빨강 노드가 연속으로 나올 수 없다)
5. 루트에서 리프까지의 모든 경로는 같은 수의 검정 노드를 포함한다

### 시간 복잡도

- 삽입(Insert): O(log n)
- 삭제(Delete): O(log n)
- 검색(Search): O(log n)
- 순회(Traversal): O(n)

### 공간 복잡도

- O(n)

## 사용 방법

```go
package main

import (
    "fmt"
    "github.com/Recordum/go-lab/data-structure/rbtree"
)

func main() {
    // 새로운 Red-Black Tree 생성
    tree := rbtree.NewRBTree()

    // 데이터 삽입
    tree.Insert(10, "Apple")
    tree.Insert(20, "Banana")
    tree.Insert(5, "Cherry")
    tree.Insert(15, "Date")
    tree.Insert(25, "Elderberry")

    // 데이터 검색
    value, found := tree.Search(20)
    if found {
        fmt.Printf("Found: %v\n", value) // Found: Banana
    }

    // In-order 순회 (정렬된 순서로 출력)
    fmt.Println("In-order traversal:")
    tree.InOrderTraversal(func(key int, value interface{}) {
        fmt.Printf("%d: %v\n", key, value)
    })

    // 데이터 삭제
    tree.Delete(15)

    // 트리 크기 확인
    fmt.Printf("Tree size: %d\n", tree.GetSize())

    // 트리 구조 출력
    fmt.Println("Tree structure:")
    fmt.Println(tree.String())

    // 트리 비우기
    tree.Clear()
    fmt.Printf("Is empty: %v\n", tree.IsEmpty())
}
```

## API

### 생성

- `NewRBTree() *RBTree`: 새로운 Red-Black Tree를 생성합니다.

### 삽입/수정

- `Insert(key int, value interface{})`: 키-값 쌍을 트리에 삽입합니다. 키가 이미 존재하면 값을 업데이트합니다.

### 검색

- `Search(key int) (interface{}, bool)`: 주어진 키에 해당하는 값을 검색합니다. 두 번째 반환값은 키의 존재 여부입니다.

### 삭제

- `Delete(key int) bool`: 주어진 키를 가진 노드를 삭제합니다. 삭제 성공 여부를 반환합니다.

### 순회

- `InOrderTraversal(visit func(key int, value interface{}))`: 중위 순회를 수행하며 각 노드에 대해 주어진 함수를 실행합니다.

### 유틸리티

- `GetSize() int`: 트리의 노드 개수를 반환합니다.
- `IsEmpty() bool`: 트리가 비어있는지 확인합니다.
- `Clear()`: 트리의 모든 노드를 제거합니다.
- `String() string`: 트리 구조를 시각화한 문자열을 반환합니다.

## 구현 세부사항

### 회전(Rotation)

트리의 균형을 유지하기 위해 좌회전(Left Rotation)과 우회전(Right Rotation)을 사용합니다.

### 삽입 후 수정(Fix Insert)

삽입 후 Red-Black Tree의 속성을 유지하기 위해 색상 변경과 회전을 수행합니다.

### 삭제 후 수정(Fix Delete)

삭제 후 Red-Black Tree의 속성을 유지하기 위해 색상 변경과 회전을 수행합니다.

## 장점

1. 균형잡힌 트리 구조로 최악의 경우에도 O(log n) 성능 보장
2. AVL Tree보다 삽입/삭제가 더 빠름 (회전 횟수가 적음)
3. 많은 언어의 표준 라이브러리에서 사용 (C++ STL map/set, Java TreeMap/TreeSet)

## 단점

1. 구현이 복잡함
2. 추가 메모리 필요 (색상 정보 저장)
3. AVL Tree보다 검색이 약간 느림 (덜 엄격한 균형)

## 테스트

```bash
# 테스트 실행
go test -v ./data-structure/rbtree

# 벤치마크 실행
go test -bench=. ./data-structure/rbtree

# 예제 실행
go test -run=Example ./data-structure/rbtree
```
