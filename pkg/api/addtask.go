package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"yp-gofinal/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if len(task.Title) == 0 {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "invalid title format"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	writeJson(w, http.StatusOK, SuccessResponse{Id: id})
}

func checkDate(task *db.Task) error {
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(DateFormat)
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	if AfterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(DateFormat)
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}

	return nil
}
