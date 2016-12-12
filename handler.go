package main

import "net/http"
import "encoding/json"

type Handler struct {
	// mailer mailer.Mailer
	// store  store.Store
}

func (h *Handler) start(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(map[string]string{"who-you-are": "WINNER!!!"})
}
