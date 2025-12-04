package main

import (
	"fmt"
	"net/http"
	"os"
	"yp-gofinal/pkg/api"
	"yp-gofinal/pkg/db"

	"github.com/go-chi/chi/v5"
)

func main() {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	webDir := "web"

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	err := db.Init(dbFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Get().Close()

	api.Init(r)

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
