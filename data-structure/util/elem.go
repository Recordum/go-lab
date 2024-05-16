package util

type Elem struct {
	Next  *Elem
	Prev  *Elem
	Value interface{}
	Key  interface{}
}
