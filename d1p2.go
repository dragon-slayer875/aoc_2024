package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var idx int

	var locations1 []int
	var locations2 []int
	var instances map[int]int
	instances = make(map[int]int)
	var similarity_score int

	for scan.Scan() {
		var (
			location1 int
			location2 int
		)
		_, err := fmt.Sscanf(scan.Text(), "%d   %d", &location1, &location2)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		locations1 = append(locations1, location1)
		locations2 = append(locations2, location2)
		idx++
	}

	for i := 0; i < idx; i++ {
		instances[locations1[i]] = 0
		for j := 0; j < idx; j++ {
			if locations1[i] == locations2[j] {
				instances[locations1[i]]++
			}
		}
	}

	for key, value := range instances {
		similarity_score = similarity_score + key*value
	}

	fmt.Println(similarity_score)
}
