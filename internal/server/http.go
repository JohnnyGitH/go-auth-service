package server

import (
	"net/http"
	"workspace/internal/handler"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router *mux.Router
	http   *http.Server
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: mux.NewRouter(),
	}
}

func (hs *HTTPServer) RegisterHanderls() {
	hs.router.Path("").Methods(http.MethodGet).HandlerFunc(handler.HomePage)

}

func (hs *HTTPServer) StartServer() {
	hs.RegisterHanderls()
	http.ListenAndServe(":9000", nil)
}
