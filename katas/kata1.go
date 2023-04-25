package main

import "fmt"

func ValidBraces(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, c)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	result1 := ValidBraces("(){}[]")   // result1 will be true
	result2 := ValidBraces("([{}])")   // result2 will be true
	result3 := ValidBraces("(}")       // result3 will be false
	result4 := ValidBraces("[(])")     // result4 will be false
	result5 := ValidBraces("[({})](]") // result5 will be false
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
}