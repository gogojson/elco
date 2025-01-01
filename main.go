package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Check if program starts")

	handler := http.DefaultServeMux

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
