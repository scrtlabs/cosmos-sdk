package types

import (
	"github.com/rs/zerolog"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"os"
)

type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
	// Values less than TraceLevel are handled as numbers.
)

var _ tmlog.Logger = (*SdkLogger)(nil)

// SdkLogger provides a wrapper around a zerolog.Logger instance. It implements
type SdkLogger struct {
	zerolog.Logger
}

func (z SdkLogger) GetLevel() Level {
	return Level(z.Logger.GetLevel())
}

func (z SdkLogger) WithSdkLogger(keyVals ...interface{}) SdkLogger {
	return SdkLogger{z.Logger.With().Fields(getLogFields(keyVals...)).Logger()}
}

// Info implements Tendermint's Logger interface and logs with level INFO. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z SdkLogger) Info(msg string, keyVals ...interface{}) {
	z.Logger.Info().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Error implements Tendermint's Logger interface and logs with level ERR. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z SdkLogger) Error(msg string, keyVals ...interface{}) {
	z.Logger.Error().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Debug implements Tendermint's Logger interface and logs with level DEBUG. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z SdkLogger) Debug(msg string, keyVals ...interface{}) {
	z.Logger.Debug().Fields(getLogFields(keyVals...)).Msg(msg)
}

// With returns a new wrapped logger with additional context provided by a set
// of key/value tuples. The number of tuples must be even and the key of the
// tuple must be a string.
func (z SdkLogger) With(keyVals ...interface{}) tmlog.Logger {
	return SdkLogger{z.Logger.With().Fields(getLogFields(keyVals...)).Logger()}
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals)%2 != 0 {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}

func (z SdkLogger) Level(lvl Level) SdkLogger {
	return SdkLogger{
		z.Logger.Level(zerolog.Level(lvl)),
	}
}

func NewDisabledLogger() SdkLogger {
	return SdkLogger{zerolog.New(os.Stdout)}.WithSdkLogger("module", "sdk/app").Level(Disabled)
}

func DisabledLogLevelManager() LogLevelManager {
	return NewLogLevelManager(Disabled)
}

func NewLogLevelManager(defaultLevel Level) LogLevelManager {
	return LogLevelManager{
		logLevels:    map[string]Level{},
		defaultLevel: defaultLevel,
	}
}

type LogLevelManager struct {
	logLevels    map[string]Level
	defaultLevel Level
}

func (manager LogLevelManager) RegisterLogLevel(key string, level Level) {
	manager.logLevels[key] = level
}

func (manager LogLevelManager) SetDefaultLevel(level Level) {
	manager.defaultLevel = level
}

func (manager LogLevelManager) GetLogLevel(key string) Level {
	lvl, found := manager.logLevels[key]
	if !found {
		return manager.defaultLevel
	}
	return lvl
}
