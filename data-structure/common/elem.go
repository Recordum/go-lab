package common

type Elem struct {
	Next *Elem
	Prev *Elem
	Value interface{}
}
