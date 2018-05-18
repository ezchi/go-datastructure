package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {

	end := node{val: 100}

	start := node{val: 20, next: &end}

	fmt.Printf("Start node: %v\n", start)
	fmt.Printf("End node: %v\n", end)
}
