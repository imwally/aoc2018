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
	i := 0
	for i < len(b)-1 {
		r1, r2 := b[i], b[i+1]
		if oppositePolarity(r1, r2) {
			b = append(b[:i], b[i+2:]...)
			if i != 0 {
				i -= 1
			}
		} else {
			i++
		}
	}

	return string(b)
}

func matchChar(r1, r2 byte) bool {
	return r1 == r2 || r1+32 == r2
}

func removeUnit(s string, r byte) string {
	b := []byte(s)
	i := 0
	for i < len(b) {
		if matchChar(b[i], r) {
			b = append(b[:i], b[i+1:]...)
		} else {
			i++
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

	shortest := len(s)
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for _, l := range alpha {
		noUnits := removeUnit(s, byte(l))
		noOpposites := removeOpposites(noUnits)
		if len(noOpposites) < shortest {
			shortest = len(noOpposites)
		}
	}

	fmt.Println(shortest)
}
