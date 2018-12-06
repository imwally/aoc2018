package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	ID   int
	X, Y int
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
			fmt.Print(m[i][j], "")
		}
		fmt.Println()
	}
}

func taxiDistance(a, b Point) int {
	return int(math.Abs(float64(b.X-a.X)) + math.Abs(float64(b.Y-a.Y)))
}

func closestNeighbor(a Point, points []Point) (bool, Point) {
	sd := 999999
	var sp Point

	var dist []int
	for _, p := range points {
		d := taxiDistance(a, p)
		if d <= sd {
			sd = d
			sp = p
			dist = append(dist, d)
		}
	}

	if len(dist) > 1 {
		if dist[len(dist)-1] == dist[len(dist)-2] {
			return true, Point{}
		}
	}

	return false, sp
}

func isInfinite(p Point, size int) bool {
	return p.X == 0 || p.Y == 0 || p.X == size-1 || p.Y == size-1

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

	var points []Point
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		coords := strings.Split(strings.Replace(scanner.Text(), " ", "", -1), ",")
		y, _ := strconv.Atoi(coords[0])
		x, _ := strconv.Atoi(coords[1])

		points = append(points, Point{i, x, y})
		i++
	}

	areas := make(map[int]int)
	var edges []int

	size := 1000
	m := matrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			dup, p := closestNeighbor(Point{i, i, j}, points)
			if dup {
				continue
			}

			m[i][j] = p.ID
			areas[p.ID] += 1

			if isInfinite(Point{i, i, j}, size) {
				edges = append(edges, p.ID)
			}
		}
	}

	for _, e := range edges {
		delete(areas, e)
	}

	var inner []int
	for _, a := range areas {
		inner = append(inner, a)
	}

	//printMatrix(m)
	sort.Ints(inner)
	fmt.Println(inner[len(inner)-1])
}
