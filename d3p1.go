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
	validExp := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	for scan.Scan() {
		fmt.Println("Program started")
		expInstances := validExp.FindAll([]byte(scan.Text()), -1)
		fmt.Printf("%q\n", expInstances)

		for _, exp := range expInstances {
			exprtkObj.SetExpression(string(exp))
			exprtkObj.CompileExpression()
			sum += int(exprtkObj.GetEvaluatedValue())
			fmt.Println(sum)
		}
	}

	fmt.Println(sum)
}
