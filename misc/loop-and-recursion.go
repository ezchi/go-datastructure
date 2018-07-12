package main

import "fmt"

func loopPrint(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func recursionPrint(n int) {
	if n > 0 {
		recursionPrint(n - 1)
		fmt.Println(n)
	}
}

func main() {

	fmt.Println("Loop print:")
	loopPrint(5)

	fmt.Println("Recursion print:")
	recursionPrint(5)
}
