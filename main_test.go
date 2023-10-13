package main

import (
	s "root/tree/slice"
	w "root/tree/word"
	"testing"
)

func BenchmarkTree(b *testing.B) {
	structure := s.NewStruct(1000)

	for i := 0; i < b.N; i++ {
		structure.Inset(5)
	}
}

func BenchmarkTreeSearch(b *testing.B) {
	structure := w.WordCollectorInit(10)

	for i := 0; i < b.N; i++ {
		structure.Insert("44 React Frontend Interview Questions")
	}
}

func BenchmarkReadAllNodes(b *testing.B) {
	structure := w.WordCollectorInit(10)

	for _, value := range []string{"What color is grass?", "A. Green", "B. Red", "C. Pink", "A",
		"Whats the first month called?", "A. December", "B. January", "C. March", "B",
		"What shape is a soccer ball?", "A. square", "B. flat", "C. round", "C"} {
		structure.Insert(value)
	}

	// for i := 0; i < b.N; i++ {
	// 	structure.ReadAllNodes("a")
	// }
}

func BenchmarkIsExist(b *testing.B) {
	structure := w.WordCollectorInit(100)

	for _, value := range []string{"What color is grass?", "A. Green", "B. Red", "C. Pink", "A",
		"Whats the first month called?", "A. December", "B. January", "C. March", "B",
		"What shape is a soccer ball?", "A. square", "B. flat", "C. round", "C"} {
		structure.Insert(value)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		structure.IsExist("Whats the first month called?")
	}
	b.StopTimer()
}

func BenchmarkRead(b *testing.B) {
	structure := w.WordCollectorInit(100)

	for _, value := range []string{"What color is grass?", "A. Green", "B. Red", "C. Pink", "A",
		"Whats the first month called?", "A. December", "B. January", "C. March", "B",
		"What shape is a soccer ball?", "A. square", "B. flat", "C. round", "C", "Was", "Ward"} {
		structure.Insert(value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		structure.ReadAllNodes("W")
	}
	b.StopTimer()
}
