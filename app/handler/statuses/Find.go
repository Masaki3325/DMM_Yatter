package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"yatter-backend-go/app/domain/object"

	"github.com/go-chi/chi/v5"
)

type StatusResponse struct {
	ID        int            `json:"id"`
	Account   object.Account `json:"account"`
	AccountID int            `json:"account_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"create_at"`
}

// Handle request for `GET /v1/statuses/{id}`
func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	id := chi.URLParam(r, "id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("cna't change into int")
		return
	}

	status, err := h.as.FindByID(r.Context(), intid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := status.AccountID
	account, err := h.ar.FindByID(r.Context(), int(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statusResponse := StatusResponse{
		ID:        int(status.ID),
		Account:   *account,
		AccountID: int(status.AccountID),
		Content:   status.Content,
		CreatedAt: status.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statusResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
