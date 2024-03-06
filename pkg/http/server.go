package http

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	presentation "duonglt.net/internal/presentation/auth"
	"github.com/spf13/viper"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter(
	authHttp presentation.AuthHttp,
) *Router {
	mux := http.NewServeMux()
	authHttp.RegisterHandlers(mux)
	return &Router{Mux: mux}
}

func (r *Router) ServeHTTP() error {
	sv := &http.Server{
		Addr:    net.JoinHostPort("", viper.GetString("PORT")),
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
