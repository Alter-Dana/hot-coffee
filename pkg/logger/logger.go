package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

var MyLogger *slog.Logger

func GetLoggerObject(FilePath string) *slog.Logger {

	option := os.O_CREATE | os.O_TRUNC | os.O_RDWR | os.O_APPEND

	file, err := os.OpenFile(FilePath, option, 0666)
	if err != nil {
		log.Fatalln("Error opening  log file: ", err)
	}

	logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	return logger

}

func ErrorWrapper(layer, functionName, context string, err error) error {
	return fmt.Errorf("%s %w\n", fmt.Sprintf("[Layer:%s,Function: %s,Context: %s]--->", layer, functionName, context), err)
}
