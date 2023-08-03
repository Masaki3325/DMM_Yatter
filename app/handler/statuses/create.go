package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Content string `json:"content"`
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a := auth.AccountOf(r)
	status := new(object.Status)
	status.AccountID = a.ID
	status.Content = req.Content

	err := h.sr.CreateNewStatus(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertedID, err := h.sr.LastInserted(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	addStatus, err := h.sr.FindByID(r.Context(), insertedID)
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
		ID:        int(addStatus.ID),
		Account:   *account,
		AccountID: int(addStatus.AccountID),
		Content:   addStatus.Content,
		CreatedAt: addStatus.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statusResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
