package configs

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Port *int
	Dir  *string
}

func NewConfiguration() (*Config, error) {

	var config *Config = new(Config)

	config.Port = flag.Int("port", 8888, "Port number to run the server on")
	config.Dir = flag.String("dir", "data", "Path to the storage directory")
	help := flag.Bool("help", false, "Show usage information")

	flag.Usage = func() {
		fmt.Println(`./hot-coffee --help
  Coffee Shop Management System
  
  Usage:
    hot-coffee [--port <N>] [--dir <S>] 
    hot-coffee --help
  
  Options:
    --help       Show this screen.
    --port N     Port number.
    --dir S      Path to the data directory.`)
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	FlagUsageCount := make(map[string]int)
	for _, flags := range os.Args[1:] {
		if strings.HasPrefix(flags, "-") {
			flagName := strings.TrimLeft(flags, "-")
			FlagUsageCount[flagName]++
			if FlagUsageCount[flagName] > 1 {
				return nil, fmt.Errorf("duplicate flag detected: flag %s used multiple times", flagName)
			}
		}
	}

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: unexpected arguments: %v\n", flag.Args())
		os.Exit(1)
	}

	if err := config.validateConfig(); err != nil {
		return nil, err
	}

	return config, nil
}

func (config *Config) validateConfig() error {
	// Port validation
	if *config.Port < 1024 || *config.Port > 49151 {
		return fmt.Errorf("invalid port number: %d. Must be between 1 and 65535", *config.Port)
	}

	// Use default directory if not specified
	if *config.Dir == "" {
		*config.Dir = "data"
	}

	cmdPath, _ := filepath.Abs("cmd")
	internalPath, _ := filepath.Abs("internal")
	modelsPath, _ := filepath.Abs("models")
	if strings.HasPrefix(*config.Dir, cmdPath) || strings.HasPrefix(*config.Dir, internalPath) || strings.HasPrefix(*config.Dir, modelsPath) {
		return errors.New("path to storage directory is inside a program file")
	}

	absPath, err := filepath.Abs(*config.Dir)
	if err != nil {
		return fmt.Errorf("could not parse path for data directory")
	}

	if strings.Contains(*config.Dir, "..") {
		return fmt.Errorf("breaking code traversal")
	}

	if filepath.IsAbs(*config.Dir) && strings.Contains(*config.Dir, "..") {
		return fmt.Errorf("relative paths or directory traversal is not allowed")
	}

	projectRootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not start server on %s", projectRootDir)
	}
	if !strings.HasPrefix(absPath, projectRootDir) {
		return fmt.Errorf("path is outside of the project directory")
	}

	if absPath == projectRootDir {
		return fmt.Errorf("cannot create file directly inside the project root directory")
	}

	return nil
}
