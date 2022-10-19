package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func computeWordScore(midTerm float32, final float32, attending float32) string {
	if midTerm < 0 || midTerm > 10 || final < 0 || final > 10 || attending < 0 || attending > 10 {
		return "Invalid input"
	}

	if attending < 8 {
		return "F"
	}

	finalScore := (attending * 10 / 100) + (midTerm * 30 / 100) + (final * 60 / 100)
	if finalScore < 4 {
		return "F"
	} else if finalScore < 5 {
		return "D"
	} else if finalScore < 5.5 {
		return "D+"
	} else if finalScore < 6.5 {
		return "C"
	} else if finalScore < 7 {
		return "C+"
	} else if finalScore < 8 {
		return "B"
	} else if finalScore < 8.5 {
		return "B+"
	} else if finalScore < 9 {
		return "A"
	}

	return "A+"
}
