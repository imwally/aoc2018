package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func in(s []string, a string) bool {
	for _, b := range s {
		if a == b {
			return true
		}
	}

	return false
}

func remove(s []string, a string) []string {
	var new []string
	for _, b := range s {
		if a != b {
			new = append(new, b)
		}
	}

	return new
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

	var dkeys []string
	dependencies := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if !in(dkeys, s[7]) {
			dkeys = append(dkeys, s[7])
		}
		dependencies[s[7]] = append(dependencies[s[7]], s[1])

	}

	d := dependencies

	// Add keys that have no depencies
	for _, key := range dkeys {
		for _, dep := range d[key] {
			if !in(dkeys, dep) {
				dkeys = append(dkeys, dep)
			}
		}
	}

	// Sort keys alphabetically
	sort.Slice(dkeys, func(i, j int) bool {
		return dkeys[i] < dkeys[j]
	})

	var steps []string
	for len(d) > 0 {
		for _, key := range dkeys {
			for _, dep := range d[key] {
				if in(steps, dep) {
					d[key] = remove(d[key], dep)
				}
			}
			if d[key] == nil {
				steps = append(steps, key)
				dkeys = remove(dkeys, key)
				delete(d, key)
				break
			}

		}
	}

	fmt.Println(strings.Join(steps, ""))
}
