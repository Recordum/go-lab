package queue

import (
	. "go-lab/data-structure/common"
	"container/list"
)

type Queue struct {
	front *Elem
	rear  *Elem
	Len   int
}

func New() *Queue {
	list.New()
	return &Queue{}
}

func (q *Queue)Front() interface{} {
	if q.front == nil {
		return nil
	}
	return q.front.Value
}

func (q *Queue)Rear() interface{} {
	if q.rear == nil {
		return nil
	}
	return q.rear.Value
}

//가장 뒤에 요소를 추가한다.
func (q *Queue) Offer(value interface{}) {
	//만약 front가 nill 일 경우에는 front와 rear 모두 새로운 요소를 가리키게 한다.
	if q.front == nil {
		q.front = &Elem{Value: value}
		q.rear = q.front
	} else {
		newElem := &Elem{Value: value}
		newElem.Next = q.rear
		q.rear.Prev = newElem
		q.rear = newElem
	}
	q.Len++
}

//가장 앞에 있는 요소를 제거하고 반환한다.
func (q *Queue) Poll() interface{} {
	if (q.front == nil) {
		return nil
	}
	pollElem := q.front
	q.front = pollElem.Prev
	q.front.Next = nil
	q.Len--
	return pollElem.Value
}

func (q *Queue) IsEmpty() bool {
	return q.Len == 0
}

func (q *Queue) Clear() {
	q.front = nil
	q.rear = nil
	q.Len = 0
}
