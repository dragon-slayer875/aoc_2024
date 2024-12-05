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
	validExp := regexp.MustCompile(`MAS|SAM`)

	for scan.Scan() {
		input = append(input, scan.Text())
	}

	for i := 0; i < len(input)-2; i++ {
		for j := 0; j < len(input)-2; j++ {
			str := string(input[i][j]) + string(input[i+1][j+1]) + string(input[i+2][j+2])
			str2 := string(input[i+2][j]) + string(input[i+1][j+1]) + string(input[i][j+2])
			if validExp.Match([]byte(str)) && validExp.Match([]byte(str2)) {
				xmas_instances++
			}
		}
	}

	fmt.Println(xmas_instances)
}
