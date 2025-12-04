package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"yp-gofinal/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	writeJson(w, http.StatusOK, SuccessResponse{})
}
