package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	ID          int
	Following   bool
	Followed_by bool
}

// Handle request for `POST /v1/accounts/{username}/follow`
func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	username := chi.URLParam(r, "username")
	var res Response

	account := auth.AccountOf(r)

	followaccount, err := h.ar.FindByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.ar.FollowAccount(int(account.ID), int(followaccount.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.ID = int(account.ID)
	res.Following = true

	result, err := h.ar.CheckFollow(r.Context(), int(account.ID), int(followaccount.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(result)
	res.Followed_by = result

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
