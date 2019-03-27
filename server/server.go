package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/renair/weather"
	"github.com/renair/weather/openweather"
	"github.com/renair/weather/resolver"
)

const defaultPort = "8080"

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	apiKey, ok := os.LookupEnv("OPENWEATHER_API_KEY")
	if !ok {
		log.Fatal("There is no key for Openweather API.\nSet it to OPENWEATHER_API_KEY enironmental variable.")
	}

	weathrApi := openweather.Initialize(apiKey)

	reslvr := resolver.Resolver{
		ApiClient: weathrApi,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(weather.NewExecutableSchema(weather.Config{Resolvers: &reslvr})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
