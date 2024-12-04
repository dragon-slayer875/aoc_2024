package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var safe_reports int

	for scan.Scan() {
		report := scan.Text()
		report_slice := strings.Split(report, " ")
		var report_slice_int []int

		for _, i := range report_slice {
			j, _ := strconv.Atoi(i)
			report_slice_int = append(report_slice_int, j)
		}

		if slices.IsSorted(report_slice_int) || slices.IsSortedFunc(report_slice_int, func(a, b int) int {
			if a < b {
				return 1
			}
			if a > b {
				return -1
			}
			return 0
		}) {

			safe := true
			for i := 0; i < len(report_slice_int)-1; i++ {
				diff := math.Abs(float64(report_slice_int[i] - report_slice_int[i+1]))
				if diff == 0 || diff > 3 {
					safe = false
				}
			}

			if safe {
				safe_reports++
			}

		}

	}
	fmt.Println(safe_reports)
}
