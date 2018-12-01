package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	frequency := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, _ := strconv.Atoi(scanner.Text())
		frequency += change
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(frequency)
}
