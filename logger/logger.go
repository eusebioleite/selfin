package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger configura o zerolog para escrever no stdout e num arquivo echo.log simultaneamente.
// O modo "dev" utiliza um ConsoleWriter amigável no stdout.
func InitLogger() {
	// Configurar formato de tempo
	zerolog.TimeFieldFormat = time.RFC3339

	// Arquivo de log
	file, err := os.OpenFile("selfin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening log file.")
	}

	// Console writer pro stdout (modo dev, colorido e formatado)
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// MultiWriter para jogar pro console e pro arquivo
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Substituir o logger global do pacote log do zerolog
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}
