package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Header struct {
	NumChildren int
	NumMeta     int
}

type Node struct {
	Header
	Children []*Node
	Metadata []int
}

func buildTree(n []int, ni int) (*Node, int) {

	nc := n[ni]
	nm := n[ni+1]

	ni += 2

	var children []*Node
	for i := 0; i < nc; i++ {
		child, nii := buildTree(n, ni)
		children = append(children, child)
		ni = nii
	}

	meta := n[ni : ni+nm]

	return &Node{Header{nc, nm}, children, meta}, ni + nm
}

func printTree(n *Node) {
	fmt.Println(n.Header, n.Metadata)
	for _, child := range n.Children {
		printTree(child)
	}
}

func sumMeta(n *Node) int {
	total := 0

	for _, meta := range n.Metadata {
		if meta <= n.NumChildren {
			total += sumMeta(n.Children[meta-1])
		}
	}

	if n.NumChildren == 0 {
		total += sum(n.Metadata)
	}

	return total
}

func sumTree(n *Node) int {
	total := 0
	for _, child := range n.Children {
		total += sumTree(child)
	}
	total += sum(n.Metadata)

	return total
}

func sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum
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

	var numbers []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	root, _ := buildTree(numbers, 0)
	//printTree(root)
	fmt.Println(sumMeta(root))
}
