package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
	as repository.Status
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account, as repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, as}
	r.With(auth.Middleware(ar)).Post("/", h.Create)

	return r
}
