package handlers

import (
	"fmt"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "KorzadiVPN API - Estado: ONLINE")
}
