package configs

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

	// To make sure flags are not used twice
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

	if err := config.ValidateConfig(); err != nil {
		return nil, err
	}

	return config, nil
}

func (config *Config) ValidateConfig() error {
	// Port validation
	if *config.Port < 1024 || *config.Port > 49151 {
		return fmt.Errorf("invalid port number: %d. Must be between 1024 and 49151", *config.Port)
	}

	// Dir validation
	validName := regexp.MustCompile(`^[a-z0-9_-]+$`)

	if !validName.MatchString(*config.Dir) {
		return errors.New("invalid directory name: use only lowercase letters, digits, hyphens, and underscores")
	}

	// Use default directory if not specified
	if strings.TrimSpace(*config.Dir) == "" {
		*config.Dir = "data"
	}

	// Get absolute paths
	absDir, err := filepath.Abs(*config.Dir)
	if err != nil {
		return fmt.Errorf("invalid storage directory path: %w", err)
	}

	// Gets absolute path for the cmd directory
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not determine project root: %w", err)
	}

	// Restricted directories
	restrictedDirs := []string{"cmd", "internal", "models"}
	for _, dir := range restrictedDirs {
		if strings.Contains(absDir, filepath.Join(projectRoot, dir)) {
			return fmt.Errorf("storage directory cannot be inside %s/", dir)
		}
	}

	// Prevent directory traversal attempts
	if strings.Contains(*config.Dir, "..") {
		return errors.New("directory traversal using '..' is not allowed")
	}

	// Ensure the storage directory is inside the project but not the root itself
	if !strings.HasPrefix(absDir, projectRoot) {
		return errors.New("storage directory must be inside the project directory")
	}
	if absDir == projectRoot {
		return errors.New("storage directory cannot be the project root")
	}

	return nil
}
