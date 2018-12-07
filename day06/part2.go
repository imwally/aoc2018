package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	ID   int
	X, Y int
}

func taxiDistance(a, b Point) int {
	return int(math.Abs(float64(a.X-b.Y)) + math.Abs(float64(b.X-a.Y)))
}

func totalDistance(a Point, points []Point) int {
	total := 0
	for _, p := range points {
		total += taxiDistance(p, a)
	}

	return total
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

	size := 500
	max := 10000

	region := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			total := totalDistance(Point{i, i, j}, points)
			if total < max {
				region += 1
			}
		}
	}

	fmt.Println(region)
}
