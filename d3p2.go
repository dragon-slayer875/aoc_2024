package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/Pramod-Devireddy/go-exprtk"
)

func mul(a, b int) int {
	return a * b
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()
	var sum int
	enable := true
	validExp := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don't\(\)|do\(\)`)

	for scan.Scan() {
		expInstances := validExp.FindAll([]byte(scan.Text()), -1)

		for _, exp := range expInstances {
			if string(exp) == "don't()" {
				enable = false
			} else if string(exp) == "do()" {
				enable = true
			} else {
				if enable == true {
					exprtkObj.SetExpression(string(exp))
					exprtkObj.CompileExpression()
					sum += int(exprtkObj.GetEvaluatedValue())
				}
			}
		}
	}

	fmt.Println(sum)
}
