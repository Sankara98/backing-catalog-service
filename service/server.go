package service

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	cftools "github.com/cloudnativego/cf-tools"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

//NewServerFromCFEnv decides the url to use for a webclient
func NewServerFromCFEnv(appEnv *cfenv.App) *negroni.Negroni {
	webClient := fulfillmentWebClient{
		rootURL: "http://localhost:3001/products",
	}

	val, err := cftools.GetVCAPServiceProperty("backing-fulfill", "url", appEnv)
	if err == nil {
		webClient.rootURL = val
	} else {
		fmt.Printf("Failed to get URL property from bound service: %v\n", err)
	}

	fmt.Printf("Using the following URL for fulfillment backing service: %s\n", webClient.rootURL)

	return NewServerFromClient(webClient)
}

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
