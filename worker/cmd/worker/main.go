package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/redanthrax/as/worker/internal/handlers"
	"github.com/redanthrax/as/worker/internal/repository"
	"github.com/redanthrax/as/worker/internal/services"
	"github.com/redanthrax/as/worker/server"
	"github.com/rs/zerolog/log"
)

func queueHandler(w http.ResponseWriter, r *http.Request) {
  var message map[string]interface{}

    // Read the message body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Unmarshal the message body
    err = json.Unmarshal(body, &message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Process the message
    fmt.Printf("Received message: %v\n", message)

    // Respond with a 200 status code
    w.WriteHeader(http.StatusOK)
}

func main() {
  //http.HandleFunc("/QueueTrigger", queueHandler)
  //port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
  //log.Printf("Listening on port %s\n", port)
  //log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

  listenAddr := "8080"
  if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
    listenAddr = val
  }

	dbConfig := repository.Config{
		StorageAccount: os.Getenv("AzureWebJobsStorage"),
	}

	db, queue, err := repository.NewDB(dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

  repo := repository.NewRepository(db, queue)
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

