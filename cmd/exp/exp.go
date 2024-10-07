package main

import (
	"log/slog"
	"os"
)

func main() {
	//h := slog.NewTextHandler(os.Stderr, nil)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	log := slog.New(h)
	log.Info("info mensage")
	log.Debug("debug message")
	log.Warn("ward message")
	log.Error("error message")
}
