package internal

import (
	"context"
	auth "duonglt.net/internal/auth/presentation"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	Port string
	Mux  *http.ServeMux
}

func NewRouter(
	port string,
	auth auth.HttpHandler,
	authenticated auth.AuthMiddleware,
) *Router {
	r := &Router{Mux: http.NewServeMux(), Port: port}
	auth.RegisterHandlers(r.Mux, authenticated.Handle)
	return r
}

func (r *Router) ServeHTTP() error {
	sv := &http.Server{
		Addr:    net.JoinHostPort("", r.Port),
		Handler: r.Mux,
	}
	go func() {
		if err := sv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server closed unexpectedly")
		}
	}()

	log.Printf("🚀 Server starting on %v\n", sv.Addr)
	return r.graceful(sv)
}

func (r *Router) graceful(sv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return sv.Shutdown(ctx)
}
