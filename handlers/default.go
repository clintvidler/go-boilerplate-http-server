package handlers

import (
	"fmt"
	"net/http"

	"github.com/clintvidler/go-boilerplate-http-server/services"
)

type DefaultHandler struct {
	logger *services.Logger
}

func NewDefaultHandler(l *services.Logger) *DefaultHandler {
	return &DefaultHandler{logger: l}
}

func (h *DefaultHandler) Placeholder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Not implemented")
}
