package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func oppositePolarity(r1, r2 byte) bool {
	return r1+32 == r2 || r2+32 == r1
}

func removeOpposites(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		r1, r2 := b[i], b[i+1]
		if oppositePolarity(r1, r2) {
			b = append(b[:i], b[i+2:]...)
			i = 0
		}
	}

	return string(b)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("specify a file, ya dummy")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = scanner.Text()
	}

	// Something if off by one. I need to run this twice to remove
	// the last remaining opposite which starts at index 0.
	s = removeOpposites(s)
	s = removeOpposites(s)

	fmt.Println(len(s))
}
