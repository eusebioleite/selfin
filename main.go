package main

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eusebioleite/selfin/database"
	logger "github.com/eusebioleite/selfin/log"
	"github.com/eusebioleite/selfin/routes"
	"github.com/eusebioleite/selfin/security"
	_ "github.com/eusebioleite/selfin/views"
	"github.com/joho/godotenv"
	_ "github.com/ncruces/go-sqlite3/driver"
	"github.com/rs/zerolog/log"
)

func main() {

	logger.InitLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error while loading .env file.")
	}

	if len(os.Args) > 1 {

		if os.Args[1] != "-password" {
			log.Fatal().Msg("Usage: ./selfin.exe -password <new_password>")
		}

		passwordPtr := flag.String("password", "", "Set new admin password")
		flag.Parse()
		if *passwordPtr != "" {
			db, err := sql.Open("sqlite3", "file:selfin.db")
			if err != nil {
				log.Fatal().Err(err).Msg("Error while trying to connect to database.")
			}

			database.InitDB(db)

			err = security.ResetPassword(*passwordPtr)
			if err != nil {
				db.Close()
				log.Fatal().Err(err).Msg("Error resetting password.")
			}

			log.Info().Msg("Password successfully reset.")
			db.Close()
		}
		os.Exit(0)
	}

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		log.Fatal().Msg("Error while reading PORT environment variable.")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

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

	database.InitDB(db)

	r := routes.SetupRouter()

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
