package logger

const (
	callerKey   = "caller"
	defaultSkip = 3

	JSONFormat Format = "json"
	TextFormat Format = "text"

	PanicLevel Level = "panic" // PanicLevel level, highest level of severity. Logs and then calls panic with the message passed to Debug, Info, ...
	FatalLevel Level = "fatal" // FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the logging level is set to Panic.
	ErrorLevel Level = "error" // ErrorLevel level. Logs. Used for errors that should definitely be noted.
	WarnLevel  Level = "warn"  // WarnLevel level. Non-critical entries that deserve eyes.
	InfoLevel  Level = "info"  // InfoLevel level. General operational entries about what's going on inside the application.
	DebugLevel Level = "debug" // DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	TraceLevel Level = "trace" // TraceLevel level. Designates finer-grained informational events than the Debug.
)

type Config struct {
	// Format json or text. Default text
	Format Format
	// Level threshold log shown
	Level Level
	// IsBeautify only if using JSONFormat
	IsBeautify bool
	// IDKey context key that used for storing context's ID
	IDKey IDKey
}

type Format string

type Level string

type IDKey string

type Meta map[string]interface{}
