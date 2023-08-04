package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/statuses/{id}`
func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	id := chi.URLParam(r, "id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("cna't change into int")
		return
	}

	status, err := h.sr.FindByID(r.Context(), intid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accountID := status.AccountID
	account, err := h.ar.FindByID(r.Context(), int(accountID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status.Account = account
	status.AccountID = 0

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
