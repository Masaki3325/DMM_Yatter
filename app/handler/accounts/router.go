package accounts

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar}
	r.Post("/", h.Create)
	r.With(auth.Middleware(ar)).Post("/update_credentials", h.Update)
	r.With(auth.Middleware(ar)).Post("/{username}/follow", h.Follow)
	r.Get("/{username}", h.Find)

	return r
}
