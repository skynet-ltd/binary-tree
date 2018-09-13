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
	masks map[byte]byte
	root  *Node
}

// NewTree ...
func NewTree(depth int) *Tree {
	if depth%len(masks) != 0 {
		panic("depth must be a multiple of 8")
	}
	masksMap := make(map[byte]byte, depth)
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
	node := tr.root
	for depth := 1; depth < tr.depth; depth++ {
		pos := depth / len(masks)
		node = lookUp(key[pos], tr.masks[byte(pos)], node)
		if node == nil {
			return nil
		}
	}
	return node.Value
}

// Insert ...
func (tr *Tree) Insert(v interface{}) []byte {
	key := []byte(uid.New(tr.depth / len(masks)))
	node := tr.root

	for depth := 1; depth < tr.depth; depth++ {
		pos := depth / len(masks)
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

func main() {
	tree := NewTree(256)

	key := tree.Insert("Hello world")
	fmt.Printf("%s\n", key)

	nv := tree.Get(key)
	fmt.Println(nv.value)
}
