package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/beforesecond/fiber-gql/src/graph/generated"
	"github.com/beforesecond/fiber-gql/src/resolvers"
	fiber "github.com/gofiber/fiber/v2"
)

const defaultPort = "5051"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := fiber.New()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	gqlHandler := srv.Handler()
	playground := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		playground(c.Context())
		return nil
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(app.Listen(":" + port))
}
