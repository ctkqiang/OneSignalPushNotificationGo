package config

import (
	"flag"
	"io"
	"os"
	"valuefarm_pushnotification_services/internal/utilities"
)

func GetDevelopmentFlavours() (string, error) {
	var envFile string

	// Use a local flag set to avoid conflict with 'go test' flags
	fs := flag.NewFlagSet("config", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // Silently ignore unknown flags (like -test.*)
	mode := fs.String("m", "debug", "Mode to run the application [debug|release]")

	// Parse os.Args[1:] to find -m, but don't fail on unknown flags
	_ = fs.Parse(os.Args[1:])

	switch *mode {
	case "release":
		envFile = ".env"
	case "debug":
		envFile = ".env.dev"
	default:
		utilities.Error("Invalid mode: %s. Use 'debug' or 'prod'", *mode)
	}

	return envFile, nil
}