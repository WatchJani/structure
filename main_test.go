package main

import (
	s "root/tree/slice"
	"testing"
)

func BenchmarkTree(b *testing.B) {
	structure := s.NewStruct(1000)

	for i := 0; i < b.N; i++ {
		structure.Inset(5)
	}
}
