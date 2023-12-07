package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := string(file)

	sum := 0
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		digits := make([]int, 0)
		for _, char := range line {
			if val, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, val)
			}
		}
		if len(digits) == 1 {
			digits = append(digits, digits[0])
		}
		result := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
		value, _ := strconv.Atoi(result)
		sum += value
	}

	fmt.Println(sum)
}
