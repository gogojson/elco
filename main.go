package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func main() {
	slog.Info("Check if program starts")

	handler := http.DefaultServeMux

	handler.HandleFunc("GET /quiz/{id}", getQuiz)

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
	//TODO: Validate if id is type int
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("Invalid param: %s", r.PathValue("id"))))
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(fmt.Sprintf("This is the path value: %d", id)))

}
