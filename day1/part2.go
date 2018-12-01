package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Specify file to read
	if len(os.Args) != 2 {
		fmt.Println("specify a file, ya dummy")
		return
	}

	// Read input data
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// May need to reread the list of frequences multiple times so
	// save to a slice for easy re-indexing
	var changes []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, _ := strconv.Atoi(scanner.Text())
		changes = append(changes, change)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Use a map to increment the amount of times a frequency is
	// seen. The program prints the first time a frequency is seen
	// twice and then exits.
	i := 0
	frequency := 0
	seen := make(map[int]int)
	for {
		frequency += changes[i]
		seen[frequency] += 1
		if seen[frequency] > 1 {
			fmt.Println(frequency)
			return
		}

		i++
		if i == len(changes) {
			i = 0
		}
	}
}
