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

func BenchmarkGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.Get(key)
	}
}

func BenchmarkMapInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := uid.New(32)
		mapTest[key] = interface{}("hello world")
	}
}

func BenchmarkMapGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mapTest[keyMap]
	}
}

func init() {
	tree = NewTree()
	key = tree.Insert("test")
	mapTest = make(map[string]interface{}, 0)
	keyMap = uid.New(32)
	mapTest[keyMap] = interface{}("test")
}
