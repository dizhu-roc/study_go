package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)
}

func swap(c, d *int) {
	*c, *d = *d, *c
}
