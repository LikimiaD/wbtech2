package main

import "time"

var (
	database map[int]map[time.Time]struct{}
)

func InitDatabase() {
	database = make(map[int]map[time.Time]struct{})
}

func AddDatabase(userID int, date time.Time) bool {
	if _, exists := database[userID]; !exists {
		database[userID] = make(map[time.Time]struct{})
	}
	database[userID][date] = struct{}{}
	return true
}

func RemoveDatabase(userID int, date time.Time) bool {
	if _, exists := database[userID]; exists {
		delete(database[userID], date)
	}
	return true
}

func GetEventsForDay(userID int, date time.Time) []time.Time {
	if events, exists := database[userID]; exists {
		if _, exists := events[date]; exists {
			return []time.Time{date}
		}
	}
	return nil
}

func GetEventsForWeek(userID int, date time.Time) []time.Time {
	var events []time.Time
	if userEvents, exists := database[userID]; exists {
		for eventDate := range userEvents {
			if isSameWeek(eventDate, date) {
				events = append(events, eventDate)
			}
		}
	}
	return events
}

func GetEventsForMonth(userID int, date time.Time) []time.Time {
	var events []time.Time
	if userEvents, exists := database[userID]; exists {
		for eventDate := range userEvents {
			if eventDate.Month() == date.Month() && eventDate.Year() == date.Year() {
				events = append(events, eventDate)
			}
		}
	}
	return events
}

func isSameWeek(a, b time.Time) bool {
	yearA, weekA := a.ISOWeek()
	yearB, weekB := b.ISOWeek()
	return yearA == yearB && weekA == weekB
}
