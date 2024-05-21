package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		slog.Error("error loading config", "err", err)
		os.Exit(1)
	}

	InitDatabase()

	http.HandleFunc("/create_event", LoggerMiddleware(CreateEventHandler))
	http.HandleFunc("/update_event", LoggerMiddleware(UpdateEventHandler))
	http.HandleFunc("/delete_event", LoggerMiddleware(DeleteEventHandler))
	http.HandleFunc("/events_for_day", LoggerMiddleware(EventsForDayHandler))
	http.HandleFunc("/events_for_week", LoggerMiddleware(EventsForWeekHandler))
	http.HandleFunc("/events_for_month", LoggerMiddleware(EventsForMonthHandler))

	slog.Info(fmt.Sprintf("server starting work: http://localhost%s", config.Port))
	if err := http.ListenAndServe(config.Port, nil); err != nil {
		slog.Error("error while starting server", "err", err)
		os.Exit(1)
	}
}
