package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func main() {
	slog.Info("Check if program starts")

	handler := http.DefaultServeMux

	handler.HandleFunc("GET /quiz/{id}", getQuiz)
	handler.HandleFunc("POST /quiz/{id}", answerQuiz)

	api := http.Server{
		Addr:    "0.0.0.0:8822",
		Handler: handler,
	}

	serverError := make(chan error, 1)
	go func() {
		slog.Info("server start", slog.String("addr", api.Addr))
		serverError <- api.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		slog.Error("server error", slog.String("errorMessage", err.Error()))
	}

}

func getQuiz(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("Invalid param: %s", r.PathValue("id"))))
		return
	}

	// Read file and save as byte
	b, err := os.ReadFile(fmt.Sprintf("quiz/%d.go", id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func answerQuiz(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("Invalid param: %s", r.PathValue("id"))))
		return
	}

	// Read file and save as byte
	b, err := os.ReadFile(fmt.Sprintf("quiz/%d.md", id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
