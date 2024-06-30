package soal

import (
	"fmt"
	"time"
)

func RunSoal4() {
	fmt.Println("\n## SOAL 4 ##")
	inputs := []struct {
		totalJointLeave  int
		joinDate         string
		plannedLeaveDate string
		leaveDuration    int
	}{
		{7, "2021-05-01", "2021-07-05", 1},
		{7, "2021-04-01", "2021-11-05", 3},
		{7, "2021-01-05", "2021-12-18", 1},
		{7, "2021-01-05", "2021-12-18", 3},
	}

	for _, input := range inputs {
		result, reason := canTakePersonalLeave(input.totalJointLeave, input.joinDate, input.plannedLeaveDate, input.leaveDuration)
		fmt.Printf("Input: Jml cuti: %d, Tgl join: %s, Tgl cuti: %s, Durasi cuti: %d\n", input.totalJointLeave, input.joinDate, input.plannedLeaveDate, input.leaveDuration)
		fmt.Printf("Output: %t\nAlasan: %s\n\n", result, reason)
	}
}

func canTakePersonalLeave(totalJointLeave int, joinDate, plannedLeaveDate string, leaveDuration int) (bool, string) {
	const totalOfficeLeave = 14
	const maxPersonalLeave = 3

	// Parse dates
	join, err := time.Parse("2006-01-02", joinDate)
	if err != nil {
		return false, "Invalid join date format"
	}
	plannedLeave, err := time.Parse("2006-01-02", plannedLeaveDate)
	if err != nil {
		return false, "Invalid planned leave date format"
	}

	// Calculate the 180th day from the join date
	day180 := join.AddDate(0, 0, 180)

	// Calculate the total number of days from day 180 to the end of the year
	endOfYear := time.Date(join.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	remainingDays := endOfYear.Sub(day180).Hours()/24 + 1

	// Calculate the allowable personal leave days
	personalLeaveDays := int((remainingDays / 365) * float64(totalOfficeLeave-totalJointLeave))

	// Check if the planned leave is within the allowable period and duration
	if plannedLeave.Before(day180) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	} else if leaveDuration > maxPersonalLeave {
		return false, fmt.Sprintf("Karena maksimal adalah %d hari cuti", maxPersonalLeave)
	} else if leaveDuration > personalLeaveDays {
		return false, fmt.Sprintf("Karena jumlah cuti pribadi yang bisa diambil adalah %d hari", personalLeaveDays)
	}
	return true, "Cuti pribadi diperbolehkan"
}
