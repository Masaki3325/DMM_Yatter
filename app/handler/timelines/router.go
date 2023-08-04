package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
	sr repository.Status
	tr repository.Timeline
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account, sr repository.Status, tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, sr, tr}
	r.Get("/public", h.Get)

	return r
}
