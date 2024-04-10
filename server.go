package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PICT-LibraryAutomation/granthpal/auth"
	"github.com/PICT-LibraryAutomation/granthpal/graph"
	"github.com/PICT-LibraryAutomation/granthpal/graph/resolvers"
	"github.com/PICT-LibraryAutomation/granthpal/remote"
	"github.com/PICT-LibraryAutomation/granthpal/search"
	"github.com/PICT-LibraryAutomation/granthpal/sessions"
	"github.com/PICT-LibraryAutomation/granthpal/tsdb"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't load .env file")
	}

	addr := os.Getenv("GRANTHPAL_ADDR")
	if addr == "" {
		log.Fatal().Msg("GRANTHPAL_ADDR not provided")
	}

	remoteDB := remote.ConnectToRemote()
	t := tsdb.ConnectToTSDB()
	sm := sessions.NewSessionManager()
	s := search.ConnectToMeilisearch()

	router := chi.NewRouter()
	router.Use(remote.RemoteMiddleware(remoteDB))
	router.Use(tsdb.TSDBMiddleware(t))
	router.Use(sessions.SessionManagerMiddleware(sm))
	router.Use(auth.AuthMiddleware())
	router.Use(search.SearchMiddleware(s))

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		Remote:         remoteDB,
		SessionManager: sm,
		Search:         s,
	}}))

	router.Handle("/", playground.ApolloSandboxHandler("GraphQL Sandbox", "/gql"))
	router.Handle("/gql", srv)

	log.Info().
		Str("Address", addr).
		Msg("Listening...")
	err = http.ListenAndServe(addr, router)
	log.Fatal().AnErr("Error", err).Msg("Server terminated")
}
