package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	ID     string
	Top    int
	Left   int
	Width  int
	Height int
}

func toClaim(s string) Claim {
	fields := strings.Fields(s)
	id := fields[0]
	left, _ := strconv.Atoi(strings.Split(fields[2], ",")[0])
	top, _ := strconv.Atoi(strings.Trim(strings.Split(fields[2], ",")[1], ":"))
	width, _ := strconv.Atoi(strings.Split(fields[3], "x")[0])
	height, _ := strconv.Atoi(strings.Split(fields[3], "x")[1])

	return Claim{
		ID:     id,
		Top:    top,
		Left:   left,
		Width:  width,
		Height: height,
	}
}

func matrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	return m
}

func printMatrix(m [][]int) {
	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
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

	var fabrics []Claim

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		c := toClaim(s)
		fabrics = append(fabrics, c)
	}

	overlap := 0
	m := matrix(1024)
	for _, fabric := range fabrics {
		for i := fabric.Left; i < fabric.Left+fabric.Width; i++ {
			for j := fabric.Top; j < fabric.Top+fabric.Height; j++ {
				m[i][j] += 1
				if m[i][j] == 2 {
					overlap += 1
				}
			}
		}
	}

	//printMatrix(m)
	fmt.Println(overlap)
}
