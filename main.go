package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/bibliothek/api"
)

func main() {
	router := httprouter.New()

	// Technical routes.
	router.GET("/ping", api.Ping())

	// Command routes (usually verb + noun, imperative).
	router.POST("/bestand/erfasse-buch", api.ErfasseBuch())
	router.POST("/ausleihe/leihe-buch-aus", api.LeiheBuchAus())

	// Query routes (usually noun).
	router.GET("/bestand/buecher", api.WelcheBuecherGibtEs())

	port := 3000
	address := fmt.Sprintf(":%d", port)

	log.Println("starting server...", port)

	err := http.ListenAndServe(address, router)
	log.Fatal(err)
}
