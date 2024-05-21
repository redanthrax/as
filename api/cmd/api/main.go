package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/redanthrax/as/api/internal/handlers"
	"github.com/redanthrax/as/api/internal/repository"
	"github.com/redanthrax/as/api/internal/services"
	"github.com/redanthrax/as/api/server"
	"github.com/rs/zerolog/log"
)

func main() {
  listenAddr := "8080"
  if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
    listenAddr = val
  }

  _ = godotenv.Load()
	dbConfig := repository.Config{
		StorageAccount: os.Getenv("AzureWebJobsStorage"),
	}

	db, err := repository.NewDB(dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

  repo := repository.NewRepository(db)
  serv := services.NewServices(repo)
  handler := handlers.NewHandler(serv)
  srv := new(server.Server)
  go func() {
    if err := srv.Run(listenAddr, handler.InitRoutes());
      err != nil && err != http.ErrServerClosed {
      log.Fatal().Err(err).Msg("error running server")
    }
  }()

  log.Info().Str("port", listenAddr).Msg("server listening")
  ctx, stop := signal.NotifyContext(
    context.Background(), 
    os.Interrupt, 
    syscall.SIGTERM, syscall.SIGINT)
  defer stop()
  <-ctx.Done()
  log.Info().Msg("server shutting down")
  if err := srv.Shutdown(context.Background()); err != nil {
    log.Error().Err(err).Msg("error shutting down the server")
  }
}
