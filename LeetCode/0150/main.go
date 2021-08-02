package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	queue := make([]int, len(tokens)*2)
	h := -1
	for _, v := range tokens {
		switch v {
		case "*":
			queue[h-1] = queue[h] * queue[h-1]
			h -= 1
		case "/":
			queue[h-1] = queue[h-1] / queue[h]
			h -= 1
		case "+":
			queue[h-1] = queue[h] + queue[h-1]
			h -= 1
		case "-":
			queue[h-1] = queue[h-1] - queue[h]
			h -= 1
		default:
			x, _ := strconv.Atoi(v)
			queue[h+1] = x
			h += 1
		}
	}
	return queue[0]
}

func main() {
	fmt.Println(
		evalRPN(
			[]string{
				"4", "13", "5", "/", "+",
			},
		),
	)
}
