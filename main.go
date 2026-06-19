package main

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eusebioleite/selfin/controllers"
	"github.com/eusebioleite/selfin/logger"
	"github.com/eusebioleite/selfin/routes"
	_ "github.com/eusebioleite/selfin/views"
	"github.com/joho/godotenv"
	_ "github.com/ncruces/go-sqlite3/driver"
	"github.com/rs/zerolog/log"
)

func main() {

	// 1. new log instance
	logger.InitLogger()

	// 2. load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error while loading .env file.")
	}
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		log.Fatal().Msg("Error while reading PORT environment variable.")
	}

	// 3. setup ctrl + c as shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// 4. setup database
	db, err := sql.Open("sqlite3", "file:selfin.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Error while trying to connect to database.")
	}
	defer func() {
		log.Info().Msg("Closing database...")
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Connection error.")
	}

	controllers.InitDB(db)

	// 5. setup gin routes
	r := routes.SetupRouter()

	// 6. setup webserver
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	for _, route := range r.Routes() {
		log.Info().
			Str("method", route.Method).
			Str("url", fmt.Sprintf("http://localhost%s%s", port, route.Path)).
			Msg("Route")
	}

	go func() {
		log.Info().Msg("Server started!.")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Error while starting webserver.")
		}
	}()

	<-ctx.Done()
	log.Warn().Msg("Shutting down...")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown.")
	}

	log.Info().Msg("Bye!.")
}
