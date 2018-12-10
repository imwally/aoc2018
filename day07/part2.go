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

	const MAX_WORKERS = 5
	workers := make(map[string]int)

	time := 0

	var steps []string
	for len(d) > 0 {
		// Time it takes for workers to finish
		time += 1

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

		for _, step := range avail {
			// Do we have a worker available?
			if len(workers) < MAX_WORKERS && workers[step] == 0 {
				// Add new worker
				seconds := int([]byte(step)[0] - 4)
				workers[step] = seconds
			}
		}

		// DEBUGGING
		//fmt.Println(workers)

		for worker, _ := range workers {
			workers[worker] -= 1

		}

		for worker, _ := range workers {
			// Worker finished, add step and remove from
			// dependencies
			if workers[worker] == 0 {
				// Add new step
				steps = append(steps, worker)

				// Remove step from dependencies
				delete(d, worker)

				// Remove newly added step from
				// dependencies of remaining steps
				for step, deps := range d {
					for _, dep := range deps {
						if in(steps, dep) {
							d[step] = remove(deps, dep)
						}
					}
				}

				delete(workers, worker)
			}
		}
	}

	fmt.Println(strings.Join(steps, ""), time)
}
