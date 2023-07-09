package main

import (
	"os"
)

func isNumerique(str string) int {
	for _, v := range str {

		if v < '0' || v > '9' {
			return 1
		}
	}
	return 0
}
func convert(str string) int {
	result := 0
	sign := 1
	start := 0
	if str[0] == '-' {
		sign = -1
		start = 1
	}
	for i := start; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			digit := int(str[i] - '0')
			result = result*10 + digit

		}
	}
	return result * sign
}

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		result := 0
		arg0 := convert(args[0])
		arg1 := args[1]
		arg2 := convert(args[2])
		if isNumerique(args[0]) == 1 || isNumerique(args[2]) == 1 {
			os.Exit(0)
		}
		switch arg1 {
		case "+":
			result = arg0 + arg2
			break
		case "-":
			result = arg0 - arg2
			break
		case "*":
			result = arg0 * arg2
			break
		case "/":
			if arg2 == 0 {
				s := []byte("No division by 0\n")
				os.Stdout.Write(s)
				os.Exit(0)
			}
			result = arg0 / arg2
			break
		case "%":
			if arg2 == 0 {
				s1 := []byte("No modulo by 0\n")
				os.Stdout.Write(s1)
				os.Exit(0)
			}
			result = arg0 % arg2
			break
		default:
			os.Stdout.Write("\n")

		}
		os.Stdout.Write(result)
	}

}
