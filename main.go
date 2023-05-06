package main

import "fmt"

func merge(arrs ...[]int) []int {
	length := 0
	for _, a := range arrs {
		length += len(a)
	}
	res := make([]int, length)
	h := res[0:length]
	for _, arr := range arrs {
		copy(h, arr)
		h = h[len(arr):]
	}
	return res
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	c := 1
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func main() {
	fmt.Println(fib(8))
}
