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
	// read all the form data into the user and them validate them then create the user
	// update the referal count of the person who invited him
	// and also update the value of the person who made the referal
}

func (h *Handler) getRegistrationForm(w http.ResponseWriter, r *http.Request) {
	// now the url needs to include the user who invited them to the bot
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.registerUser).Methods(http.MethodPost)
	router.HandleFunc("/register", h.getRegistrationForm).Methods(http.MethodPost)
}
