package soal

import (
	"fmt"
	"strings"
)

func RunSoal1() {
	fmt.Println("\n## SOAL 1 ##")
	input1 := []string{"abcd", "acbd", "aaab", "acbd", "acbd"}
	input2 := []string{"pisang", "goreng", "enak", "sekali", "rasanya"}
	input3 := []string{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"}

	// Output results
	fmt.Println(findFirstDuplicate(5, input1))  // Should output [2 4]
	fmt.Println(findFirstDuplicate(5, input2))  // Should output false
	fmt.Println(findFirstDuplicate(11, input3)) // Should output [3 5 10]
}

func findFirstDuplicate(n int, stringList []string) interface{} {
	occurences := make(map[string]int)

	var indexDuplicate []int
	var flag bool
	var firstFoundStr string
	for i, s := range stringList {
		lowerCaseStr := strings.ToLower(s)

		if index, found := occurences[lowerCaseStr]; found {
			if firstFoundStr == lowerCaseStr || firstFoundStr == "" {
				firstFoundStr = lowerCaseStr

				if !flag {
					indexDuplicate = append(indexDuplicate, index+1, i+1)
				} else {
					indexDuplicate = append(indexDuplicate, i+1)
				}
			}

			flag = true
		}

		occurences[lowerCaseStr] = i
	}

	if len(indexDuplicate) > 0 {
		return indexDuplicate
	}
	return false
}
