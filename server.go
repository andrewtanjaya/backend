package main

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"tpaWeb/graph"
	"tpaWeb/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	opt, err := pg.ParseURL("postgres://oadxreyqmffgnh:1bd75bcd83cf9e171d63dc62dc19be625150e7541b0dc2f2b4ff8fe56706fa69@ec2-54-247-118-139.eu-west-1.compute.amazonaws.com:5432/d7sbv5fgdtfjdr?sslmode=require")
	if err != nil {
		panic(err)
	}

	pgDB := pg.Connect(opt)

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"postgres://oadxreyqmffgnh:1bd75bcd83cf9e171d63dc62dc19be625150e7541b0dc2f2b4ff8fe56706fa69@ec2-54-247-118-139.eu-west-1.compute.amazonaws.com:5432/d7sbv5fgdtfjdr"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: pgDB}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}