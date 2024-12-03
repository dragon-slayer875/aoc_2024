package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var idx int

	var locations1 []int
	var locations2 []int
	var distances []int
	var distanceSum int

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

	slices.Sort(locations1)
	slices.Sort(locations2)

	for i := 0; i < idx; i++ {
		distance := locations2[i] - locations1[i]
		if distance < 0 {
			distance = -distance
		}
		distances = append(distances, distance)
	}

	for i := 0; i < idx; i++ {
		distanceSum += distances[i]
	}

	fmt.Println(distanceSum)
}
