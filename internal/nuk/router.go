package nuk

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() Router {
	mux := http.NewServeMux()
	return Router{Mux: mux}
}

func (r Router) ServeHTTP() error {
	sv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r.Mux,
	}
	go func() {
		if err := sv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to listen and serve: %v\n", err)
		}
	}()

	log.Printf("🚀 Server starting on %v\n", sv.Addr)
	return r.graceful(sv)
}

func (r Router) graceful(sv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("⚠️ Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return sv.Shutdown(ctx)
}
