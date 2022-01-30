package logger

import (
	"os"
	"testing"
)

// TestNew tests the Logger::New functionality
func TestNew(*testing.T) {
	New(os.Stdout, os.Stderr)
}

func TestLogging(t *testing.T) {
	const message = "Hello, world!"

	logger := New(os.Stdout, os.Stderr)
	t.Run("Debug", func(*testing.T) { logger.Debug("%v", message) })
	t.Run("Info", func(*testing.T) { logger.Info("%v", message) })
	t.Run("Warn", func(*testing.T) { logger.Warn("%v", message) })
	t.Run("Error", func(*testing.T) { logger.Error("%v", message) })
	t.Run("Log", func(*testing.T) { logger.Log("%v", message) })
}
