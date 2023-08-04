package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
)

// Handle request for `GET /v1/timelines/public`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	maxID := r.URL.Query().Get("max_id")
	sinceID := r.URL.Query().Get("since_id")
	limit := r.URL.Query().Get("limit")
	var response []object.Status

	statuses, err := h.tr.GetStatuses(r.Context(), maxID, sinceID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, status := range statuses {
		userID := status.AccountID
		account, err := h.ar.FindByID(r.Context(), int(userID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		status.Account = account
		status.AccountID = 0
		response = append(response, status)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
