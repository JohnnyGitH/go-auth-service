package server

import (
	"net/http"
	"workspace/internal/handler"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router *mux.Router // Gorilla/Mux for routing
	http   *http.Server
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: mux.NewRouter(),
	}
}

func (hs *HTTPServer) RegisterHandlers() {
	hs.router.Path("").Methods(http.MethodGet).HandlerFunc(handler.HomePage) // Routing to home Page

}

func (hs *HTTPServer) StartServer() {
	hs.RegisterHandlers()
	http.ListenAndServe(":9000", nil)
}
