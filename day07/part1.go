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

	dependencies := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		dependencies[s[7]] = append(dependencies[s[7]], s[1])

		// Add any steps without dependencies
		if _, ok := dependencies[s[1]]; !ok {
			dependencies[s[1]] = []string{}
		}
	}

	d := dependencies

	var steps []string
	for len(d) > 0 {
		// Create a new slice to store newly available steps
		// (also clears available slice from last iteration)
		var avail []string
		for step, deps := range d {
			// No more dependencies, make step available
			if len(deps) == 0 {
				avail = append(avail, step)
			}
		}

		// Sort available steps
		sort.Strings(avail)
		step := avail[0]

		// Add new step
		steps = append(steps, step)

		// Remove step from dependencies
		delete(d, step)

		// Remove newly added step from dependencies of remaining steps
		for step, deps := range d {
			for _, dep := range deps {
				if in(steps, dep) {
					d[step] = remove(deps, dep)
				}
			}
		}
	}

	fmt.Println(strings.Join(steps, ""))
}
