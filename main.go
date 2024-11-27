package main

import "fmt"

func main() {
	n := 5
	target := 42
	findNumberWithSum("", n, target)
}

func findNumberWithSum(result string, n int, target int) {
	if n > 0 && target >= 0 {
		//Here n>1 and n<=9, so starting from 1
		startDigit := '0'
		if len(result) == 0 { // This is Special case so number cannot start from 0
			startDigit = '1'
		}

		for d := startDigit; d <= '9'; d++ {
			findNumberWithSum(result+string(d), n-1, target-int(d-'0'))
		}
	} else if n == 0 && target == 0 {
		// If the number becomes n-digit and its sum of digits equals the given sum, print it
		fmt.Print(result + " ")
	}
}
