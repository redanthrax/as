package main

import (
	"net/http"
	"os"
  "context"
  "os/signal"
  "syscall"

	"github.com/redanthrax/as/api/internal/handlers"
	"github.com/redanthrax/as/api/internal/services"
	"github.com/redanthrax/as/api/server"
	"github.com/rs/zerolog/log"
)

func main() {
  listenAddr := "8080"
  if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
    listenAddr = val
  }

  serv := services.NewServices()
  handler := handlers.NewHandler(serv)
  srv := new(server.Server)
  go func() {
    if err := srv.Run(listenAddr, handler.InitRoutes()); err != nil && err != http.ErrServerClosed {
      log.Fatal().Err(err).Msg("error running server")
    }
  }()

  log.Info().Str("port", listenAddr).Msg("server listening")
  ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
  defer stop()
  <-ctx.Done()
  log.Info().Msg("server shutting down")
  if err := srv.Shutdown(context.Background()); err != nil {
    log.Error().Err(err).Msg("error shutting down the server")
  }
}
