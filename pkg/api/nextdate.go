package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", fmt.Errorf("Invalid date format: %v", err)
	}
	firstDate := afterNow(date, now)
	if !firstDate {
		date = now
	}

	partsRepeat := strings.Split(repeat, " ")
	switch partsRepeat[0] {
	case " ":
		fmt.Println("Task is complete and will be deleted")
	case "d":
		if len(partsRepeat) < 2 || len(partsRepeat) > 3 {
			return "", fmt.Errorf("Invalid repeat format: %v", partsRepeat)
		}

	}
}

func afterNow(date, now time.Time) bool {
	if date.After(now) {
		return true
	}
	return false
}

func countDays(partsRepeat []string, date time.Time) (time.Time, error) {
	var nextdate time.Time
	for _, part := range partsRepeat[1:] {
		days, err := strconv.Atoi(part)
		if err != nil {
			return time.Time{}, fmt.Errorf("Error parsing repeat part: %v", err)
		}
		nextdate = date.AddDate(0, 0, days)
	}
	return nextdate, nil
}
