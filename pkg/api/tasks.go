package api

import (
	"net/http"
	"yp-gofinal/pkg/db"
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)

	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	writeJson(w, http.StatusOK, TasksResponse{
		Tasks: tasks,
	})
}
