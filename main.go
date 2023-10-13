package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	structure := w.WordCollectorInit(27395)

	file, err := os.Open("./test")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
		structure.Insert(scanner.Text())
	}

	// fmt.Println(structure.IsExistPrefix("b"))

	fmt.Println(structure.ReadAllNodes("ba"))

	fmt.Println(structure.Counter)
}
