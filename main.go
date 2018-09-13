package main

import (
	"fmt"

	"github.com/skynet-ltd/uid"
)

const (
	b8 = 1
	b7 = 1 << iota
	b6
	b5
	b4
	b3
	b2
	b1

	maxUint = 1<<8 - 1
)

var masks = []byte{b1, b2, b3, b4, b5, b6, b7, b8}

// NodeValue ...
type NodeValue struct {
	value interface{}
}

// Node ...
type Node struct {
	Hash  []byte
	Left  *Node
	Right *Node
	Value *NodeValue
}

// Tree ...
type Tree struct {
	depth int
	masks []byte
	root  *Node
}

// NewTree ...
func NewTree(depth int) *Tree {
	if depth%len(masks) != 0 {
		panic("depth must be a multiple of 8")
	}
	masksMap := make([]byte, depth)
	for i := 0; i <= depth; i++ {
		masksMap[byte(i)] = masks[byte(i)%byte(len(masks))]
	}
	return &Tree{masks: masksMap, root: &Node{}, depth: depth}
}

func lookUp(b, m byte, parent *Node) *Node {
	if isSet(b, m) {
		return parent.Right
	}
	return parent.Left
}

func isSet(val, mask byte) bool {
	return (val & mask) != 0
}

// Get ...
func (tr *Tree) Get(key []byte) *NodeValue {
	pos := 0
	node := tr.root
	for depth := 0; depth < tr.depth; depth++ {
		pos = depth / len(masks)
		node = lookUp(key[pos], tr.masks[byte(pos)], node)
		if node == nil {
			return nil
		}
	}
	return node.Value
}

// RecursiveGet ...
func (tr *Tree) RecursiveGet(key []byte) *NodeValue {
	var get func(n *Node) *Node
	var pos int
	var depth int
	get = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if depth == tr.depth {
			return n
		}
		pos = depth / 8 //bit
		depth++
		if isSet(key[pos], tr.masks[byte(pos)]) {
			return get(n.Right)
		}
		return get(n.Left)
	}
	return get(tr.root).Value
}

// Insert ...
func (tr *Tree) Insert(v interface{}) []byte {
	var pos int
	key := []byte(uid.New(tr.depth / len(masks)))
	node := tr.root

	for depth := 0; depth < tr.depth; depth++ {
		pos = depth / 8 //bit
		if isSet(key[pos], tr.masks[byte(pos)]) {
			if node.Right == nil {
				node.Right = &Node{}
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = &Node{}
			}
			node = node.Left
		}
	}
	node.Hash = key
	node.Value = &NodeValue{v}
	return key
}

// RecursiveInsert ...
func (tr *Tree) RecursiveInsert(v interface{}) []byte {
	key := []byte(uid.New(tr.depth / len(masks)))
	pos := 0
	depth := 0
	var insert func(n **Node) **Node
	insert = func(n **Node) **Node {
		if depth == tr.depth {
			*n = &Node{
				Hash:  key,
				Value: &NodeValue{v},
			}
			return nil
		}

		if *n == nil {
			*n = &Node{}
		}
		pos = depth / 8 //bit
		depth++

		if isSet(key[pos], tr.masks[byte(pos)]) {
			return insert(&(*n).Right)
		}
		return insert(&(*n).Left)
	}
	insert(&tr.root)
	return key
}

func main() {
	tree := NewTree(256)

	key := tree.Insert("Hello world")
	fmt.Printf("%s\n", key)

	nv := tree.RecursiveGet(key)
	fmt.Println(nv.value)
	nv = tree.Get(key)
	fmt.Println(nv.value)
}
