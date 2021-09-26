package main

import (
	"github.com/skhatri/api-router-go/router"
	"github.com/skhatri/pokedex/server/poke"
	"log"
	"net/http"
)

/*
http://localhost:6100/search/by-type?q=flying,water
http://localhost:6100/search/by-name?q=Charizard
http://localhost:6100/search/effective-against/by-type?q=dragon
http://localhost:6100/search/effective-against/by-name?q=Charizard
*/
func main() {
	httpRouter := router.NewHttpRouterBuilder().
		WithOptions(router.HttpRouterOptions{
			LogRequest: false,
		}).Configure(func(configurer router.ApiConfigurer) {
		configurer.Get("/search/by-name", poke.FindByName).
			Get("/search/by-type", poke.FindByType).
			Get("/search/effective-against/by-type", poke.FindEffectiveAgainstType).
			Get("/search/effective-against/by-name", poke.FindEffectiveAgainstName)
	}).Build()
	var address = "0.0.0.0:6100"
	log.Printf("Listening on %s\n", address)
	http.ListenAndServe(address, httpRouter)
}
