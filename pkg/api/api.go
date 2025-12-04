package api

import (
	"encoding/json"
	"net/http"
	"yp-gofinal/pkg/db"

	"github.com/go-chi/chi/v5"
)

const DateFormat = "20060102"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Id       int64  `json:"id,omitempty"`
	Nextdate string `json:"nextdate,omitempty"`
}

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func Init(r *chi.Mux) {
	r.Get("/api/nextdate", nextDayHandler)
	r.Get("/api/task", getTaskHandler)
	r.Post("/api/task", addTaskHandler)
	r.Put("/api/task", updateTaskHandler)
	r.Delete("/api/task", deleteTaskHandler)
	r.Get("/api/tasks", tasksHandler)
	r.Post("/api/task/done", doneTaskHandler)
}

func writeJson(w http.ResponseWriter, statusCode int, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
