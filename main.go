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
	masks map[byte]byte
	root  *Node
}

// NewTree ...
func NewTree() *Tree {
	masksMap := make(map[byte]byte, 0)
	for i := 0; i <= maxUint; i++ {
		masksMap[byte(i)] = masks[byte(i)%byte(len(masks))]
	}
	return &Tree{masks: masksMap, root: &Node{}}
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
	for depth := 1; depth <= maxUint; depth++ {
		pos := depth / 8
		node = lookUp(key[pos], tr.masks[byte(pos)], node)
		if node == nil {
			return nil
		}
	}
	return node.Value
}

// Insert ...
func (tr *Tree) Insert(v interface{}) []byte {
	key := []byte(uid.New(32))
	node := tr.root

	for depth := 1; depth <= maxUint; depth++ {
		pos := depth / 8
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
	tree := NewTree()

	key := tree.Insert("Hello world")
	fmt.Printf("%s\n", key)

	nv := tree.Get(key)
	fmt.Println(nv.value)
}
