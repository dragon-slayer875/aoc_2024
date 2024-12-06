package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// "slices"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var input []string
	rules := make(map[string][]string)
	var mid_sum int
	validExp := regexp.MustCompile(`[0-9]{2}\|[0-9]{2}`)

	for scan.Scan() {
		str := scan.Text()
		if validExp.Match([]byte(str)) {
			nums := strings.Split(str, "|")
			rules[nums[1]] = append(rules[nums[1]], nums[0])
		} else if scan.Text() == "" {
		} else {
			input = append(input, scan.Text())
		}
	}

	for _, updates := range input {
		safe := true
		updates_slice := strings.Split(updates, ",")
		for i := 0; i < len(updates_slice); i++ {
			check_slice := updates_slice[i:]
			for _, rule := range rules[updates_slice[i]] {
				for _, page := range check_slice {
					if rule == page {
						safe = false
					}
				}
			}
		}
		if safe {
			mid_num, _ := strconv.Atoi(updates_slice[len(updates_slice)/2])
			mid_sum += mid_num
		}
	}

	fmt.Println(mid_sum)
}
