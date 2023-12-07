package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
  name string
}

  var muxRouterInstance* mux.Router = mux.NewRouter()

func NewMuxRouter() Router{
  return &muxRouter{name:"Prixas"}
}

func (mx *muxRouter) GET(uri string, f func(w http.ResponseWriter, req *http.Request)) {
  muxRouterInstance.HandleFunc(uri, f).Methods("GET")
}
func (mx *muxRouter) POST(uri string, f func(w http.ResponseWriter, req *http.Request)) {
  muxRouterInstance.HandleFunc(uri, f).Methods("POST")
}

func (mx *muxRouter) SERVE(port string) {
  fmt.Printf("Mux Http server running on port %v", port)
  http.ListenAndServe(port,  muxRouterInstance)
}

