package main

import (
	"fmt"
	"os"
	"store/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Just4Ease/axon/v2"
	"github.com/Just4Ease/axon/v2/options"
	"github.com/Just4Ease/axon/v2/systems/jetstream"
	"github.com/Just4Ease/graphrpc"
	"github.com/Just4Ease/graphrpc/server"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var eventStore axon.EventStore
	var err error

	if eventStore, err = jetstream.Init(options.Options{
		ServiceName:         "store",
		Address:             "0.0.0.0:4222",
		AuthenticationToken: "",
	}); err != nil {
		logrus.Fatal(err)
	}

	address := fmt.Sprintf("0.0.0.0:%s", port)
	if err := graphrpc.NewServer(
		eventStore,
		handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})),
		server.SetGraphHTTPServerAddress(address),
		server.SetGraphQLPath("/graphql"),
	).Serve(); err != nil {
		logrus.Fatalf("Could not start server on %s. Got error: %s", address, err.Error())
	}
}
