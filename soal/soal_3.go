package soal

import "fmt"

func RunSoal3() {
	fmt.Println("\n## SOAL 3 ##")
	validInputs := []string{
		"{[<><[{}]]]}",
		"<<[[{{(([]<>))}}]]>>",
		"{{[[<<[[<<[[<<[<><>]]>>]]>>]]}}",
	}

	invalidInputs := []string{
		"]",
		"[>",
		"<[>]",
		"<>",
		"<<[<<{[{([]<>)}]}>]>",
	}

	fmt.Println("Valid inputs:")
	for _, input := range validInputs {
		fmt.Printf("%s: %v\n", input, isValidCharacter(input))
	}

	fmt.Println("Invalid inputs:")
	for _, input := range invalidInputs {
		fmt.Printf("%s: %v\n", input, isValidCharacter(input))
	}
}

func isValidCharacter(s string) bool {
	if len(s) < 1 || len(s) > 4096 {
		return false
	}

	stack := []int32{}
	for _, char := range s {

		switch char {
		case '<', '(', '[', '{':
			stack = append(stack, char)
		case '>':
			if len(stack) == 0 || stack[len(stack)-1] != '<' {
				return false
			}
			stack = stack[:len(stack)-1]
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

		default:
			return false
		}
	}

	return len(stack) == 0
}
