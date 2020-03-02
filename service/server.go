package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

//NewServer Creates a new server and fits route Handlers
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	webClient := fulfillmentWebClient{
		rootURL: "http://localhost:3001/products",
	}

	initRoutes(mx, formatter, webClient)

	n.UseHandler(mx)
	return n
}

//NewServerFromClient configures and returns a server
func NewServerFromClient(webClient fulfillmentClient) *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()

	mx := mux.NewRouter()

	initRoutes(mx, formatter, webClient)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, webClient fulfillmentClient) {
	mx.HandleFunc("/", rootHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog",
		getAllCatalogItemsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog/{productId}",
		getCatalogItemDetailsHandler(formatter, webClient)).Methods("GET")
}
