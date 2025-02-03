package server

import (
	"net/http"
	"portfolio/cmd/web"
	"portfolio/cmd/web/components"
	"portfolio/cmd/web/containers"
	"portfolio/internal/middleware"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.AuthorizeUser)

	r.Handle("/", templ.Handler(web.HomePage()))

	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		components.DailyTasks().Render(r.Context(), w)
	})

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		containers.Home().Render(r.Context(), w)
	})

	r.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		components.Faq().Render(r.Context(), w)
	})

	r.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		containers.Profile().Render(r.Context(), w)
	})

	r.HandleFunc("/topusers", func(w http.ResponseWriter, r *http.Request) {
		containers.TopUsers().Render(r.Context(), w)
	})

	fileServer := http.FileServer(http.FS(web.Files))
	r.PathPrefix("/assets/").Handler(fileServer)

	return r
}

// CORS middleware
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
