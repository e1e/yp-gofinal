package api

import (
	"net/http"
	"strconv"
	"yp-gofinal/pkg/db"
)

type TaskResponse struct {
	Task *db.Task `json:"task"`
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if len(idStr) == 0 {
		// !!не проходит тест
		// writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "Не указан идентификатор"})
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

	// !!не проходит тест
	// writeJson(w, http.StatusOK, TaskResponse{
	// 	Task: task,
	// })
	successMap := map[string]string{
		"id":      task.Id,
		"date":    task.Date,
		"title":   task.Title,
		"comment": task.Comment,
		"repeat":  task.Repeat,
	}
	writeJson(w, http.StatusOK, successMap)
}
