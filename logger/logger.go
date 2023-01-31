package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
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
	fileWriter := openFileWriter()

	outputs := zerolog.MultiLevelWriter(fileWriter, os.Stdout)

	logger = zerolog.New(outputs).With().Timestamp().Logger()
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Error(msg string) {
	logger.Error().Msg(msg)
	os.Exit(1)
}
