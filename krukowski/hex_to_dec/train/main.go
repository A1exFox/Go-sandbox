package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Please bring (0x or 0o or 0b)  number or 'stop' to exit:")
	var input string
	for {
		fmt.Scanln(&input)
		input = strings.Trim(input, " ")
		input = strings.ToLower(input)
		if input == "stop" {
			break
		}
		if len(input) < 3 {
			fmt.Println("Input number has to be more than 3 characters. Try again:")
			continue
		}
		input_base, input_number := splitInput(input)
		base, err := parseBase(input_base)
		if err != nil {
			fmt.Println("Input number has not any base. Try again:")
			continue
		}
		i := new(big.Int)
		if _, ok := i.SetString(input_number, base); !ok {
			fmt.Println("Invalid convertation. Try again:")
			continue
		}
		fmt.Println(i)
	}
}

func splitInput(input string) (string, string) {
	return input[0:2], input[2:]
}
func parseBase(base string) (int, error) {
	switch base {
	case "0x":
		return 16, nil
	case "0o":
		return 8, nil
	case "0b":
		return 2, nil
	default:
		return 0, errors.New("unkown base")
	}
}
