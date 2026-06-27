package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {

	zerolog.TimeFieldFormat = time.RFC3339

	file, err := os.OpenFile("selfin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening log file.")
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}
