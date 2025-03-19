package user

import (
	"net/http"
	"portfolio/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	service types.UserStore
}

func NewHandler(s types.UserStore) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	// 
}

func (h *Handler) getRegistrationForm(w http.ResponseWriter, r *http.Request) {
	// now the url needs to include the user who invited them to the bot
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.registerUser).Methods(http.MethodPost)
	router.HandleFunc("/register", h.getRegistrationForm).Methods(http.MethodPost)
}
