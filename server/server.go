package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"

	"github.com/renair/weather"
	"github.com/renair/weather/openweather"
	"github.com/renair/weather/persistence"
	"github.com/renair/weather/resolver"
)

const defaultPort = "8080"
const defaultRedisPort = "6379"

func main() {
	serverPort, ok := os.LookupEnv("PORT")
	if !ok {
		serverPort = defaultPort
	}
	apiKey, ok := os.LookupEnv("OPENWEATHER_API_KEY")
	if !ok {
		log.Fatal("There is no key for Openweather API.\nSet it to OPENWEATHER_API_KEY enironmental variable.")
	}
	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		log.Fatal("There is no key for Redis server.\nSet it to REDIS_HOST enironmental variable.")
	}
	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		log.Printf("There is no key for Redis port, using default - %s.\nIf required - set it to REDIS_PORT enironmental variable.", defaultRedisPort)
		redisPort = defaultRedisPort
	}

	weathrApi := openweather.Initialize(apiKey)
	weathrApi.SetMeasureUnits(openweather.METRIC)
	storage := persistence.NewStorage(redisHost, redisPort)

	reslvr := resolver.Resolver{
		ApiClient:   weathrApi,
		Persistance: storage,
	}

	http.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(weather.NewExecutableSchema(weather.Config{Resolvers: &reslvr})))
	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
