package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var input []string
	var xmas_instances int
	validExp := regexp.MustCompile(`XMAS|SAMX`)

	for scan.Scan() {
		input = append(input, scan.Text())
	}

	for _, line := range input {
		expInstances := validExp.FindAll([]byte(line), -1)
		xmas_instances += len(expInstances)
	}

	// vertical lines

	var vertical_input []string

	for i := 0; i < len(input); i++ {
		var vertical_line string
		for j := 0; j < len(input); j++ {
			vertical_line += string(input[j][i])
		}
		vertical_input = append(vertical_input, vertical_line)
	}

	for _, line := range vertical_input {
		expInstances := validExp.FindAll([]byte(line), -1)
		xmas_instances += len(expInstances)
	}

	// diagonal lines

	var diagonal_input []string

	for i := 0; i < len(input); i++ {
		var diagonal_line string
		for j := 0; j <= i; j++ {
			diagonal_line += string(input[i-j][j])
		}
		diagonal_input = append(diagonal_input, diagonal_line)
	}

	for i := 1; i < len(input); i++ {
		var diagonal_line string
		for j := i; j < len(input); j++ {
			diagonal_line += string(input[len(input)-j+i-1][j])
		}
		diagonal_input = append(diagonal_input, diagonal_line)
	}

	for _, line := range diagonal_input {
		expInstances := validExp.FindAll([]byte(line), -1)
		xmas_instances += len(expInstances)
	}

	// opposite diagonal lines

	var opposite_diagonal_input []string

	for i := 0; i < len(input); i++ {
		var diagonal_line string
		for j := 0; j <= i; j++ {
			diagonal_line += string(input[len(input)-1-i+j][j])
		}
		opposite_diagonal_input = append(opposite_diagonal_input, diagonal_line)
	}

	for i := 1; i < len(input); i++ {
		var diagonal_line string
		for j := i; j < len(input); j++ {
			diagonal_line += string(input[0-i+j][j])
		}
		fmt.Println(diagonal_line)
		opposite_diagonal_input = append(opposite_diagonal_input, diagonal_line)
	}

	for _, line := range opposite_diagonal_input {
		expInstances := validExp.FindAll([]byte(line), -1)
		xmas_instances += len(expInstances)
	}

	fmt.Println(xmas_instances)
}
