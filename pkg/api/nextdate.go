package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	var nowTime time.Time
	var err error

	if len(now) == 0 {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(DateFormat, now)
		if err != nil {
			writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
	}

	nextDate, err := NextDate(nowTime, date, repeat)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}

	repeatPeriod, repeatValue, err := parseRepeat(repeat)
	if err != nil {
		return "", err
	}

	switch repeatPeriod {
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if AfterNow(date, now) {
				break
			}
		}
	case "d":
		for {
			date = date.AddDate(0, 0, repeatValue)
			if AfterNow(date, now) {
				break
			}
		}
	}

	return date.Format(DateFormat), nil

}

func AfterNow(date, now time.Time) bool {
	dateDay := date.Truncate(24 * time.Hour)
	nowDay := now.Truncate(24 * time.Hour)
	return dateDay.After(nowDay)
}

func parseRepeat(repeat string) (string, int, error) {
	var repeatValue int
	var err error

	if len(repeat) == 0 {
		return "", 0, errors.New(`repeat rule is empty`)
	}

	params := strings.Split(repeat, " ")

	repeatPeriod := params[0]

	if repeatPeriod != "d" && repeatPeriod != "y" {
		return "", 0, errors.New("invalid repeat period format")
	}

	switch repeatPeriod {
	case "d":

		if len(params) != 2 {
			return "", 0, errors.New("invalid repeat period format")
		}

		repeatValue, err = strconv.Atoi(params[1])

		if err != nil {
			return "", 0, errors.New("invalid repeat value format")
		}

		if repeatValue < 0 || repeatValue > 400 {
			return "", 0, errors.New("invalid repeat value format")
		}
		return repeatPeriod, repeatValue, nil

	case "y":
		return repeatPeriod, 1, nil
	}

	return "", 0, errors.New("invalid repeat period format")
}
