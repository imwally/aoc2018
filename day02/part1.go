package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	duplicates := make(map[string]map[rune]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		chars := make(map[rune]int)
		for _, c := range s {
			chars[c] += 1
		}

		for k, v := range chars {
			if v < 2 {
				delete(chars, k)
			}
		}
		duplicates[s] = chars
	}

	twos, threes := 0, 0
	for _, v := range duplicates {
		total, len := 0, 0
		for _, n := range v {
			total += n
			len += 1
		}

		if total%len == 0 {
			if total/len == 3 {
				threes += 1
			}
			if total/len == 2 {
				twos += 1
			}
		} else {
			twos += 1
			threes += 1
		}
	}

	fmt.Println("twos:", twos, "threes:", threes, "checksum:", twos*threes)
}
