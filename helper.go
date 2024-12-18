package aoc2024helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ParseArgs(args []string) (string, error) {
	if len(args) < 2 {
		fmt.Println("No input file provided")
		fmt.Printf("Usage: go run %s `inputfile`\n", args[0])
		return "", errors.New("no input file provided")
	}
	if len(args) != 2 {
		fmt.Println("Illegal number of arguments provided")
		fmt.Printf("Usage: go run %s `inputfile`\n", args[1])
		return "", errors.New("illegal number of arguments provided")

	}
	return args[1], nil
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func StripLine(line string) []int {
	var output []int
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		var word strings.Builder
		for !unicode.IsSpace(char) {
			word.WriteString(string(char))
			i++
			if i >= len(line) {
				break
			}
			char = rune(line[i])
		}
		if word.Len() > 0 {
			val, err := strconv.Atoi(word.String())
			CheckError(err)
			output = append(output, val)
		}
	}
	return output
}

func AbsoluteVal(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}
