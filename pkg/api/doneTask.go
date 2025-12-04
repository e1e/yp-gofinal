package api

import (
	"net/http"
	"strconv"
	"time"
	"yp-gofinal/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if len(idStr) == 0 {
		responseMap := map[string]string{"error": "Не указан идентификатор"}
		writeJson(w, http.StatusBadRequest, responseMap)
		return
	}

	_, err := strconv.Atoi(idStr)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	task, err := db.GetTask(idStr)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if len(task.Repeat) == 0 {
		err = db.DeleteTask(task.Id)
		if err != nil {
			writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		writeJson(w, http.StatusInternalServerError, SuccessResponse{})
		return
	}

	nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	task.Date = nextDate
	err = db.UpdateDateTask(task.Date, task.Id)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	writeJson(w, http.StatusOK, SuccessResponse{})
}
