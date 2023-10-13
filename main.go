package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	w "root/tree/word"
	"strings"
	"time"
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

	structure := w.WordCollectorInit(2629735)

	file, err := os.Open("./1.1million word list.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}

		for _, word := range strings.Fields(scanner.Text()) {
			structure.Insert(word)
		}
	}

	start := time.Now()

	structure.ReadAllNodes("la")

	fmt.Println(time.Since(start))
}
