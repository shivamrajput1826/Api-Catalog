package logger

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Logger struct {
	logInstance zerolog.Logger
}

func CreateLogger(logContext string) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logInstance := zerolog.New(os.Stdout).With().Timestamp().Str("context", logContext).Logger()

	return &Logger{
		logInstance: logInstance,
	}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	newLogger := &Logger{
		logInstance: l.logInstance,
	}
	reqCtx, ok := ctx.Value("requestCtx").(map[string]interface{})
	if !ok {
		return newLogger
	}
	logContext := map[string]string{
		"requestId": reqCtx["requestId"].(string),
		"userId":    reqCtx["userId"].(string),
		"userEmail": reqCtx["userEmail"].(string),
	}
	for key, val := range logContext {
		if val != "" {
			newLogger.logInstance = newLogger.logInstance.With().Str(key, val).Logger()
		}
	}
	return newLogger
}

func (l *Logger) WithFiberContext(c *fiber.Ctx) *Logger {
	newLogger := &Logger{
		logInstance: l.logInstance,
	}
	contextKeys := []string{"requestId", "userId", "userEmail"}
	for _, key := range contextKeys {
		if val := c.Locals(key); val != nil {
			newLogger.logInstance = newLogger.logInstance.With().Str(key, val.(string)).Logger()
		}
	}
	return newLogger
}

func (l *Logger) Info(message string, fields ...interface{}) {
	l.logInstance.Info().Fields(fields).Msg(message)
}

func (l *Logger) Error(message string, fields ...interface{}) {
	l.logInstance.Error().Fields(fields).Msg(message)
}

func (l *Logger) Debug(message string, fields ...interface{}) {
	l.logInstance.Debug().Fields(fields).Msg(message)
}
