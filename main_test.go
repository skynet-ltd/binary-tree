package main

import (
	"testing"

	"github.com/skynet-ltd/uid"
)

var tree *Tree
var key []byte
var mapTest map[string]interface{}
var keyMap string

func BenchmarkInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.Insert("hello world")
	}
}

func BenchmarkRecursiveInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.RecursiveInsert("hello world")
	}
}

func BenchmarkGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.Get(key)
	}
}

func BenchmarkRecursiveGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.RecursiveGet(key)
	}
}

func init() {
	tree = NewTree(256)
	key = tree.Insert("test")
	mapTest = make(map[string]interface{}, 0)
	keyMap = uid.New(32)
	mapTest[keyMap] = interface{}("test")
}
