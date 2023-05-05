package main

import (
	"fmt"
	"net/http"

	"github.com/kabaf81/BuildAWebApplication/pkg/handlers"
)

const portNumber = ":9991"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/About", handlers.About)
	http.HandleFunc("/SiteMap", handlers.SiteMap)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
