package internal

import (
	"net/http"

	vHttp "duonglt.net/pkg/http"
)

type HttpHandler struct {
	rootHandler rootHandler
}
type rootHandler struct {
}

func NewHttpHandler() HttpHandler {
	return HttpHandler{
		rootHandler: newRootHandler(),
	}
}

func (h HttpHandler) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /", h.rootHandler)
}

func newRootHandler() rootHandler {
	return rootHandler{}
}

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vHttp.Ok(w, map[string]string{})
}
