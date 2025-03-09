// A simple program to demonstrate how to use Go slices and maps
package main

import "fmt"

func main() {
	var students = []string{"Alice", "Bob", "Charlie"}
	var scores = map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Charlie": 70,
	}

	fmt.Println("Students:", students)
	fmt.Println("Scores:", scores)

	// Get the score for a student
	student := "Alice"
	score, ok := scores[student]
	if ok {
		fmt.Printf("%s's score is %d\n", student, score)
	} else {
		fmt.Println(student, "does not have a score")
	}

	// Update the score for a student
	scores["Alice"] = 95
	fmt.Println("Updated scores:", scores)
}
