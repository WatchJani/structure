package main

import (
	"fmt"
	w "root/tree/word"
)

func main() {
	// structure := s.NewStruct(1000)
	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// start := time.Now()
	// for _, number := range numbers {
	// 	structure.Inset(number)
	// }

	// fmt.Println(time.Since(start))

	// structure.Read()

	//=======================WORD================================

	structure := w.WordCollectorInit(100)

	structure.Insert("arb")
	structure.Insert("arr")

	fmt.Println(structure.Memory)
}
