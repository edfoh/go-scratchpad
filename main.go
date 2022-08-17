package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].

Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

"3+2*2"
*/
func Calculate(s string) int {
	stk := make([]int, 0)
	num := 0
	sign := '+'
	for i, r := range s {
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')

			if i != len(s)-1 {
				continue
			}
		}

		if r == ' ' && i != len(s)-1 {
			continue
		}

		switch sign {
		case '+':
			stk = append(stk, num)
		case '-':
			stk = append(stk, -num)
		case '*':
			newNum := stk[len(stk)-1] * num
			stk[len(stk)-1] = newNum
		case '/':
			newNum := stk[len(stk)-1] / num
			stk[len(stk)-1] = newNum
		}

		num = 0
		sign = r
	}

	res := 0
	for _, el := range stk {
		res += el
	}
	return res
}
