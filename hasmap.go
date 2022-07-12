package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	key   int
	value int
}

type HashMap struct {
	capacity int
	list     []list.List
}

func NewHashMap(capacity int) HashMap {
	return HashMap{
		capacity: capacity,
		list:     make([]list.List, capacity),
	}
}

func (h HashMap) getHashValue(key int) int {
	return key % h.capacity
}

func (h HashMap) insert(key, value int) {
	hash := h.getHashValue(key)
	h.list[hash].PushFront(Node{
		key:   key,
		value: value,
	})
}

func (h HashMap) getValue(key int) int {
	hash := h.getHashValue(key)
	currList := h.list[hash]

	for ele := currList.Front(); ele != nil; ele = ele.Next() {
		node := ele.Value.(Node)
		if key == node.key {
			return node.value
		}

	}

	return -1
}

func main() {
	h := NewHashMap(5)

	h.insert(1, 10)
	h.insert(2, 20)
	h.insert(3, 30)
	h.insert(11, 40)
	h.insert(21, 50)
	h.insert(22, 100)
	h.insert(32, 23)

	fmt.Println(h.getValue(1))
	fmt.Println(h.getValue(11))
	fmt.Println(h.getValue(21))
	fmt.Println(h.getValue(22))
	fmt.Println(h.getValue(32))
}
