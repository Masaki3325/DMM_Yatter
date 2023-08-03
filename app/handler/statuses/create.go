package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	content string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Print(r.Body)
		return
	}

	a := auth.AccountOf(r)
	status := new(object.Status)
	status.AccountID = a.ID
	status.Content = req.content
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Println("😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠😠")
	fmt.Println(status.Content)
	fmt.Println(req.content)

	err := h.as.CreateNewStatus(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
