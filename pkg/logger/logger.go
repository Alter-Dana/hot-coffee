package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

var MyLogger *slog.Logger

// This function creates a new instance of structured logger
func GetLoggerObject(FilePath string) *slog.Logger {

	option := os.O_CREATE | os.O_TRUNC | os.O_RDWR | os.O_APPEND

	file, err := os.OpenFile(FilePath, option, 0666)
	if err != nil {
		log.Fatalln("Error opening log file: ", err)
	}

	// Creation of a new instance of logger. "File" variable is a destination file
	logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	// NewJSONHandler makes sure the logger gives output in JSON format
	// &slog.HandlerOptions is a pointer to a struct that allows customization of logger
	// AddSource: adds the source file and line number where the log message was called
	// Level: sets the minimum logging level to "Debug", ensuring all messages at this level and higher are logged

	return logger

}

func ErrorWrapper(layer, functionName, context string, err error) error {
	return fmt.Errorf("%s %w\n", fmt.Sprintf("[Layer:%s,Function: %s,Context: %s]--->", layer, functionName, context), err)
}
