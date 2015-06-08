package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"

	"DockerBeanstalkSample/handler"
)

func main() {
	n := negroni.New(
	  negroni.NewRecovery(),
	  negroni.NewLogger(),
	  negroni.NewStatic(http.Dir("static")),
	)

	router := mux.NewRouter()

	//Init controllers
	router.HandleFunc("/", handler.Home)

	n.UseHandler(router)
	n.Run(":8080")
}
