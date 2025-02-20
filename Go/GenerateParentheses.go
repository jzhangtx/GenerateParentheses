package main

import "fmt"

func Parentheses(n, left, right int, s []byte, result *[]string) {
	if left == n && right == n {
		c := make([]byte, len(s))
		copy(c, s)
		*result = append(*result, string(c))
		return
	}

	if left < n {
		s = append(s, '(')
		Parentheses(n, left+1, right, s, result)
		s = s[:len(s)-1]
	}

	if right < left {
		s = append(s, ')')
		Parentheses(n, left, right+1, s, result)
		s = s[:len(s)-1]
	}
}

func GenerateParentheses(n int) []string {
	var result []string
	var s []byte
	Parentheses(n, 0, 0, s, &result)

	return result
}

func main() {
	for {
		fmt.Print("Please enter number n: ")
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		fmt.Print("The generated parentheses: ")
		v := GenerateParentheses(n)
		for i := 0; i < len(v); i++ {
			fmt.Print("\"", v[i], "\"  ")
		}
		fmt.Println("")
	}
}
