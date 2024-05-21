package main

import (
	"errors"
	"net/http"
)

var (
	layout = "2006-01-02"
)

var (
	ErrorWrongMethod = errors.New("wrong method for handler")
	ErrorParseForm   = errors.New("server can't parse form")
	ErrorInputValues = errors.New("bad request")
	ErrorServerSide  = errors.New("server can't complete your request")
)

func CreateEventHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		RespondWithError(w, http.StatusInternalServerError, ErrorWrongMethod)
		return
	}

	if err := req.ParseForm(); err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, ErrorParseForm)
		return
	}

	userID := req.Form.Get("user_id")
	date := req.Form.Get("date")
	if userIDInteger, dateTime, err := CheckPostTwoParams(userID, date); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
	} else {
		RespondPostRequest(w, AddDatabase(userIDInteger, dateTime))
	}
}

func UpdateEventHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		RespondWithError(w, http.StatusInternalServerError, ErrorWrongMethod)
		return
	}

	if err := req.ParseForm(); err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, ErrorParseForm)
		return
	}

	userID := req.Form.Get("user_id")
	oldDate := req.Form.Get("old_date")
	newDate := req.Form.Get("new_date")

	if userIDInteger, oldDateTime, newDateTime, err := CheckPostThreeParams(userID, oldDate, newDate); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
	} else {
		if RemoveDatabase(userIDInteger, oldDateTime) {
			RespondPostRequest(w, AddDatabase(userIDInteger, newDateTime))
		} else {
			RespondWithError(w, http.StatusInternalServerError, ErrorServerSide)
		}
	}
}

func DeleteEventHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		RespondWithError(w, http.StatusInternalServerError, ErrorWrongMethod)
		return
	}

	if err := req.ParseForm(); err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, ErrorParseForm)
		return
	}
	userID := req.Form.Get("user_id")
	date := req.Form.Get("date")
	if userIDInteger, dateTime, err := CheckPostTwoParams(userID, date); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
	} else {
		RespondPostRequest(w, RemoveDatabase(userIDInteger, dateTime))
	}
}

func EventsForDayHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		RespondWithError(w, http.StatusInternalServerError, ErrorWrongMethod)
		return
	}

	if err := req.ParseForm(); err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, ErrorParseForm)
		return
	}

	userID := req.Form.Get("user_id")
	date := req.Form.Get("date")
	if userIDInteger, dateTime, err := CheckPostTwoParams(userID, date); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
	} else {
		events := GetEventsForDay(userIDInteger, dateTime)
		RespondAny(w, http.StatusOK, events)
	}
}

func EventsForWeekHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		RespondWithError(w, http.StatusInternalServerError, ErrorWrongMethod)
		return
	}

	if err := req.ParseForm(); err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, ErrorParseForm)
		return
	}

	userID := req.Form.Get("user_id")
	date := req.Form.Get("date")
	if userIDInteger, dateTime, err := CheckPostTwoParams(userID, date); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
	} else {
		events := GetEventsForWeek(userIDInteger, dateTime)
		RespondAny(w, http.StatusOK, events)
	}
}

func EventsForMonthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		RespondWithError(w, http.StatusInternalServerError, ErrorWrongMethod)
		return
	}

	if err := req.ParseForm(); err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, ErrorParseForm)
		return
	}

	userID := req.Form.Get("user_id")
	date := req.Form.Get("date")
	if userIDInteger, dateTime, err := CheckPostTwoParams(userID, date); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
	} else {
		events := GetEventsForMonth(userIDInteger, dateTime)
		RespondAny(w, http.StatusOK, events)
	}
}
