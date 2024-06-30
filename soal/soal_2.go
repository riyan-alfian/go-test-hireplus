package soal

import (
	"fmt"
	"math"
)

func RunSoal2() {
	fmt.Println("\n## SOAL 2 ##")
	total1, paid1 := 700649, 800000
	total2, paid2 := 657650, 600000
	total3, paid3 := 575650, 580000

	// Output results
	fmt.Println("Change for first customer:", calculateChange(total1, paid1))
	fmt.Println("Change for second customer:", calculateChange(total2, paid2))
	fmt.Println("Change for third customer:", calculateChange(total3, paid3))
}

func calculateChange(total, paid int) interface{} {
	if paid < total {
		return false
	}

	var denominations = []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	result := struct {
		Change              int
		ChangeDenominations map[int]int
	}{}

	change := paid - total
	roundedChange := int(math.Floor(float64(change)/100) * 100)
	result.Change = roundedChange

	changeDenominations := make(map[int]int)
	for _, denom := range denominations {
		if roundedChange >= denom {
			count := roundedChange / denom
			changeDenominations[denom] = count
			roundedChange -= count * denom
		}
	}
	result.ChangeDenominations = changeDenominations

	return result
}
