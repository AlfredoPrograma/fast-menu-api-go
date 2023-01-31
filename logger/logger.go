package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const LOGS_FILENAME = "app.log"

var logger zerolog.Logger

func openFileWriter() *os.File {
	fileWriter, err := os.OpenFile(LOGS_FILENAME, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fileWriter
}

func Init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	fileWriter := openFileWriter()

	outputs := zerolog.MultiLevelWriter(fileWriter, consoleWriter)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger = zerolog.New(outputs).
		With().
		Timestamp().
		Stack().
		Logger()
}

func Use() *zerolog.Logger {
	return &logger
}
