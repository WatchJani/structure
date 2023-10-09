package main

import "testing"

func BenchmarkTree(b *testing.B) {
	structure := NewStruct(1000)

	for i := 0; i < b.N; i++ {
		structure.Inset(5)
	}
}
