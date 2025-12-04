package api

import (
	"net/http"
	"strconv"
	"yp-gofinal/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if len(idStr) == 0 {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "Не указан идентификатор"})
		return
	}

	_, err := strconv.Atoi(idStr)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.DeleteTask(idStr)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	writeJson(w, http.StatusOK, SuccessResponse{})
}
