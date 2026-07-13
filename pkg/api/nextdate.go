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
	partsRepeat := strings.Fields(repeat)
	if len(partsRepeat) == 0 {
		return "", fmt.Errorf("Invalid repeat format: %v", partsRepeat)
	}
	switch partsRepeat[0] {
	case "d":
		nextDate, err := countDays(partsRepeat, date, now)
		if err != nil {
			return "", fmt.Errorf("Error calculating next date for daily repeat: %v", err)
		}
		return nextDate.Format("20060102"), nil
	case "y":
		if len(partsRepeat) != 1 {
			return "", fmt.Errorf("Invalid repeat format for yearly: %v", partsRepeat)
		}
		nextDate := date.AddDate(1, 0, 0)
		for !nextDate.After(now) {
			nextDate = nextDate.AddDate(1, 0, 0)

		}
		return nextDate.Format("20060102"), nil
	case "w":
		nextDate, err := countWeeks(partsRepeat, date, now)
		if err != nil {
			return "", fmt.Errorf("Error calculating next date for weekly repeat: %v", err)
		}
		return nextDate.Format("20060102"), nil
	case "m":
		nextDate, err := countMonths(partsRepeat, date, now)
		if err != nil {
			return "", fmt.Errorf("Error calculating next date for monthly repeat: %v", err)
		}
		return nextDate.Format("20060102"), nil
	default:
		return "", nil
	}
}

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func countDays(partsRepeat []string, date, now time.Time) (time.Time, error) {
	var nextDate time.Time
	if len(partsRepeat) != 2 {
		return time.Time{}, fmt.Errorf("Invalid repeat format: %v", partsRepeat)
	}

	days, err := strconv.Atoi(partsRepeat[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Error parsing repeat part: %v", err)
	}
	if days < 1 || days > 400 {
		return time.Time{}, fmt.Errorf("Invalid number of days: %d", days)
	}
	nextDate = date.AddDate(0, 0, days)
	for !afterNow(nextDate, now) {
		nextDate = nextDate.AddDate(0, 0, days)
	}

	return nextDate, nil
}

func countWeeks(partsRepeat []string, date, now time.Time) (time.Time, error) {
	var (
		nextDate time.Time
		daysInt  []int
	)

	if len(partsRepeat) != 2 {
		return time.Time{}, fmt.Errorf("Invalid repeat format: %v", partsRepeat)
	}

	days := strings.Split(partsRepeat[1], ",")
	for _, day := range days {
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return time.Time{}, fmt.Errorf("Error parsing repeat part: %v", err)
		}
		if dayInt < 1 || dayInt > 7 {
			return time.Time{}, fmt.Errorf("Invalid day of the week: %d", dayInt)
		}
		if dayInt == 7 {
			dayInt = 0
		}
		daysInt = append(daysInt, dayInt)
	}

	nextDate = date
	for {
		if afterNow(nextDate, now) {
			weekDay := int(nextDate.Weekday())

			for _, day := range daysInt {
				if weekDay == day {
					return nextDate, nil
				}
			}
		}
		nextDate = nextDate.AddDate(0, 0, 1)
	}
}

func countMonths(partsRepeat []string, date, now time.Time) (time.Time, error) {
	var (
		nextDate  time.Time
		daysInt   []int
		monthsInt []int
	)
	if len(partsRepeat) != 2 && len(partsRepeat) != 3 {
		return time.Time{}, fmt.Errorf("Invalid repeat format: %v", partsRepeat)
	}

	days := strings.Split(partsRepeat[1], ",")
	for _, day := range days {
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return time.Time{}, fmt.Errorf("Error parsing repeat part: %v", err)
		}
		if dayInt != -1 && dayInt != -2 && (dayInt < 1 || dayInt > 31) {
			return time.Time{}, fmt.Errorf("Invalid day of the month: %d", dayInt)
		}

		daysInt = append(daysInt, dayInt)
	}

	if len(partsRepeat) == 2 {
		monthsInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	}
	if len(partsRepeat) == 3 {
		months := strings.Split(partsRepeat[2], ",")

		for _, month := range months {
			monthInt, err := strconv.Atoi(month)
			if err != nil {
				return time.Time{}, fmt.Errorf("Error parsing repeat part: %v", err)
			}

			if monthInt < 1 || monthInt > 12 {
				return time.Time{}, fmt.Errorf("Invalid month: %d", monthInt)
			}

			monthsInt = append(monthsInt, monthInt)
		}
	}

	nextDate = date

	for {
		if nextDate.After(now) {
			currentMonth := int(nextDate.Month())

			monthOK := false
			for _, allowedMonth := range monthsInt {
				if currentMonth == allowedMonth {
					monthOK = true
					break
				}
			}

			if monthOK {
				currentDay := nextDate.Day()

				lastDay := time.Date(nextDate.Year(), nextDate.Month()+1, 0, 0, 0, 0, 0, nextDate.Location()).Day()

				for _, day := range daysInt {
					if currentDay == day {
						return nextDate, nil
					}
					if day == -1 && currentDay == lastDay {
						return nextDate, nil
					}
					if day == -2 && currentDay == (lastDay-1) {
						return nextDate, nil
					}
				}
			}
		}
		nextDate = nextDate.AddDate(0, 0, 1)
	}
}
