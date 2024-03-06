package sharedPresentation

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

	authPresentation "duonglt.net/internal/auth/presentation"
	"github.com/spf13/viper"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter(
	auth authPresentation.Http,
) *Router {
	mux := http.NewServeMux()
	// Register http handlers
	auth.RegisterHandlers(mux)

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
