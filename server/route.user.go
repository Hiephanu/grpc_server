package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegsterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.HandleGetUser).Methods("GET")
	router.HandleFunc("/users", h.HandlePostUSer).Methods("POSt")
}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.store.GetUser()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) HandlePostUSer(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the PostUser method to save the user to the database
	if err := h.store.PostUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
