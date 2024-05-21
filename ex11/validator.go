package main

import (
	"strconv"
	"time"
)

func CheckPostTwoParams(userID, date string) (int, time.Time, error) {
	if userID == "" || date == "" {
		return 0, time.Time{}, ErrorInputValues
	}

	userIDInteger, err := strconv.Atoi(userID)
	if err != nil {
		return 0, time.Time{}, ErrorInputValues
	}

	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return 0, time.Time{}, ErrorInputValues
	}

	return userIDInteger, dateTime, nil
}

func CheckPostThreeParams(userID, oldDate, newDate string) (int, time.Time, time.Time, error) {
	if userID == "" || oldDate == "" || newDate == "" {
		return 0, time.Time{}, time.Time{}, ErrorInputValues
	}

	userIDInteger, err := strconv.Atoi(userID)
	if err != nil {
		return 0, time.Time{}, time.Time{}, ErrorInputValues
	}

	oldDateTime, err := time.Parse(layout, oldDate)
	if err != nil {
		return 0, time.Time{}, time.Time{}, ErrorInputValues
	}

	newDateTime, err := time.Parse(layout, newDate)
	if err != nil {
		return 0, time.Time{}, time.Time{}, ErrorInputValues
	}

	return userIDInteger, oldDateTime, newDateTime, nil
}

func CheckDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, ErrorInputValues
	}

	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, ErrorInputValues
	}

	return dateTime, nil
}
