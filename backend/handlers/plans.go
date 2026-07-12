package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/models"
)

func Plans(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Plans)
}
