package server

import (
	"encoding/json"
	"libstack/pkgs/auth"
	"net/http"

	"github.com/gorilla/mux"
)

type Payload[T any] struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
	Data  T      `json:"data,omitempty"`
}

type Server struct {
	Router *mux.Router
}

func New() Server {
	return Server{Router: createRoutes()}
}

func createRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(corsOriginMiddleware)

	// Login should not be behind authentication
	r.Path("/api/v1/login").Handler(http.HandlerFunc(login)).Methods(http.MethodPost, http.MethodOptions)

	// Authentication required
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(auth.Middleware)
	api.Path("/logout").Handler(http.HandlerFunc(logout)).Methods(http.MethodPost, http.MethodOptions)

	api.Path("/librarian.title.add").Handler(http.HandlerFunc(addTitle)).Methods(http.MethodPost, http.MethodOptions)
	api.Path("/librarian.title.archive").Handler(http.HandlerFunc(archiveTitle)).Methods(http.MethodPost, http.MethodOptions)
	api.Path("/patron.title.borrow").Handler(http.HandlerFunc(borrowTitle)).Methods(http.MethodPost, http.MethodOptions)
	api.Path("/patron.title.hold").Handler(http.HandlerFunc(holdTitle)).Methods(http.MethodPost, http.MethodOptions)
	api.Path("/patron.title.return").Handler(http.HandlerFunc(returnTitle)).Methods(http.MethodPost, http.MethodOptions)
	// TODO(mchenryc): add some kind of route for viewing all the titles
	return r
}

const JSON = "application/json"
const INVALID_CONTENT_TYPE = "invalid_content_type"
const INVALID_CREDS = "invalid_creds"

func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"ok": false, error: "internal_error"}`))
}

func login(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != JSON {
		payload := Payload[interface{}]{Ok: false, Error: INVALID_CONTENT_TYPE}
		bs, err := json.Marshal(payload)
		if err != nil {
			internalError(w)
			return
		}
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Set("Content-Type", JSON)
		w.Write(bs)
		return
	}
	// TODO(mchenryc): handle login
	// TODO(mchenryc): handle invalid login
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", JSON)
	// TODO(mchenryc): replace with actual token
	payload := Payload[string]{Ok: true, Data: "todo-real-token"}
	bs, err := json.Marshal(payload)
	if err != nil {
		internalError(w)
		return
	}
	w.Write(bs)
}

func logout(w http.ResponseWriter, r *http.Request) {
	// TODO(mchenryc): handle logout
}

func addTitle(w http.ResponseWriter, r *http.Request) {
	// TODO(mchenryc): handle addTitle
	w.WriteHeader(http.StatusOK)
}

func archiveTitle(w http.ResponseWriter, r *http.Request) {
	// TODO(mchenryc): handle archiveTitle
	w.WriteHeader(http.StatusOK)
}

func borrowTitle(w http.ResponseWriter, r *http.Request) {
	// TODO(mchenryc): handle borrowTitle
	w.WriteHeader(http.StatusOK)
}

func holdTitle(w http.ResponseWriter, r *http.Request) {
	// TODO(mchenryc): handle holdTitle
	w.WriteHeader(http.StatusOK)
}

func returnTitle(w http.ResponseWriter, r *http.Request) {
	// TODO(mchenryc): handle returnTitle
	w.WriteHeader(http.StatusOK)
}

func corsOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(mchenryc): make ALLOWED_ORIGIN and env variable
		allowedOrigin := "localhost:3000"
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		next.ServeHTTP(w, r)
	})
}
