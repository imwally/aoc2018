package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func matrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	return m
}

// Longest common subsequence using dynamic programming ALGORITS!
func lcs(s1, s2 string) int {
	l := len(s1) + 1
	m := matrix(l)

	f, k := 0, 0
	for i := 1; i < l; i++ {
		for j := 1; j < l; j++ {
			if s1[i-1] == s2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
			} else {
				m[i][j] = max(m[i][j-1], m[i-1][j])
			}
			k = j
		}
		f = i
	}

	return m[f][k]
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

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	greatest := 0
	var box1, box2 string
	for i := 0; i < len(lines); i++ {
		a := lines[i]
		for j := 0; j < len(lines); j++ {
			if i != j {
				b := lines[j]
				current := lcs(a, b)
				if current > greatest {
					greatest = current
					box1, box2 = a, b
				}
			}
		}
	}

	fmt.Println(box1)
	fmt.Println(box2)
}
