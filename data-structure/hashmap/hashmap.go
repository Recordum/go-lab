package hashmap

import (
	"crypto/sha256"
	"encoding/binary"
	"go-lab/data-structure/util"
)

type HashMap struct {
	//사용중인 버킷 크기
	size int

	//해시맵의 총 용량(버킷의 크기)
	capacity int

	//해시맵의 버킷
	buckets []*util.Elem
}

var hash = sha256.New()

const (
	loadFactor   = 0.75
	initCapacity = 16
)

func New() *HashMap {
	return &HashMap{capacity: initCapacity, buckets: make([]*util.Elem, initCapacity)}
}

func hashCode(key interface{}) int {
	hash.Reset()
	var hashValue uint64
	switch key := key.(type) {
	case string:
		hash.Write([]byte(key))
		hashBytes := hash.Sum(nil)
		hashValue = binary.BigEndian.Uint64(hashBytes)
	case int:
		hashValue = uint64(key)
	case int64:
		hashValue = uint64(key)
	case float64:
		hashValue = uint64(key)
	default:
		panic("Unsupported key type")
	}
	hashInt := int(hashValue)
	if hashInt < 0 {
		hashInt = -hashInt
	}
	return hashInt
}

func (h *HashMap) Set(key, value interface{}) {
	if h.size > int(float64(h.capacity)*loadFactor) {
		//resize
	}
	index := hashCode(key) % h.capacity
	h.insert(key, value, index)
}

func (h *HashMap) Get(key interface{}) interface{} {
	index := hashCode(key) % h.capacity
	for e := h.buckets[index]; e != nil; e = e.Next {
		if e.Key == key {
			return e.Value
		}
	}
	return nil
}

func (h *HashMap) insert(key, value interface{}, index int) {
	if h.buckets[index] == nil {
		h.buckets[index] = &util.Elem{Key: key, Value: value}
		h.size++
		return 
	}
	for e := h.buckets[index]; e != nil; e = e.Next {
		//key가 같은 경우 value를 덮어쓴다.
		if e.Key == key {
			e.Value = value
			return
		}
		if e.Next == nil {
			e.Next = &util.Elem{Key: key, Value: value}
			e.Next.Prev = e
			h.size++
			return
		}
	}
}

func (h *HashMap) Size() int {
	return h.size
}

func (h *HashMap) Capacity() int {
	return h.capacity
}

// func (h *HashMap) resize() {
// 	c := h.capacity * 2
// 	newBuckets := make([]util.Elem, c)

// 	for b, i range h.buckets {

// 	}

// }
