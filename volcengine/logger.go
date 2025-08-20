package volcengine

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"io"
	"log"
	"os"
)

// A LogLevelType defines the level logging should be performed at. Used to instruct
// the SDK which statements should be logged.
type LogLevelType uint

// LogLevel returns the pointer to a LogLevel. Should be used to workaround
// not being able to take the address of a non-composite literal.
func LogLevel(l LogLevelType) *LogLevelType {
	return &l
}

// Value returns the LogLevel value or the default value LogOff if the LogLevel
// is nil. Safe to use on nil value LogLevelTypes.
func (l *LogLevelType) Value() LogLevelType {
	if l != nil {
		return *l
	}
	return LogOff
}

// Matches returns true if the v LogLevel is enabled by this LogLevel. Should be
// used with logging sub levels. Is safe to use on nil value LogLevelTypes. If
// LogLevel is nil, will default to LogOff comparison.
func (l *LogLevelType) Matches(v LogLevelType) bool {
	c := l.Value()
	return c&v == v
}

// AtLeast returns true if this LogLevel is at least high enough to satisfies v.
// Is safe to use on nil value LogLevelTypes. If LogLevel is nil, will default
// to LogOff comparison.
func (l *LogLevelType) AtLeast(v LogLevelType) bool {
	c := l.Value()
	return c >= v
}

const (
	// LogOff states that no logging should be performed by the SDK. This is the
	// default state of the SDK, and should be use to disable all logging.
	LogOff LogLevelType = iota * 0x10000

	// LogDebug state that debug output should be logged by the SDK. This should
	// be used to inspect request made and responses received.
	LogDebug
)

// Debug Logging Sub Levels
const (
	LogDebugWithRequest     LogLevelType = LogDebug | (1 << iota)
	LogDebugWithRequestBody LogLevelType = (1 << iota) | LogDebugWithRequest
	LogDebugWithRequestID
	LogDebugWithEndpoint
	LogDebugWithConfig

	// LogDebugWithSigning states that the SDK should log request signing and
	// presigning events. This should be used to log the signing details of
	// requests for debugging. Will also enable LogDebug.
	LogDebugWithSigning

	// LogDebugWithRequestRetries states the SDK should log when service requests will
	// be retried. This should be used to log when you want to log when service
	// requests are being retried. Will also enable LogDebug.
	LogDebugWithRequestRetries

	LogDebugWithResponse
	LogDebugWithResponseBody LogLevelType = (1 << iota) | LogDebugWithResponse

	// LogDebugWithHTTPBody states the SDK should log HTTP request and response
	// HTTP bodys in addition to the headers and path. This should be used to
	// see the volcenginebody content of requests and responses made while using the SDK
	// Will also enable LogDebug.
	LogDebugWithHTTPBody LogLevelType = (1 << iota) | LogDebugWithRequestBody | LogDebugWithResponseBody

	// LogDebugWithRequestErrors states the SDK should log when service requests fail
	// to build, send, validate, or unmarshal.
	LogDebugWithRequestErrors LogLevelType = LogDebug | (1 << iota)

	// LogDebugWithEventStreamBody states the SDK should log EventStream
	// request and response bodys. This should be used to log the EventStream
	// wire unmarshaled message content of requests and responses made while
	// using the SDK Will also enable LogDebug.
	LogDebugWithEventStreamBody

	// LogInfoWithInputAndOutput states the SDK should log STRUCT input and output
	// Will also enable LogInfo.
	LogInfoWithInputAndOutput

	// LogDebugWithInputAndOutput states the SDK should log STRUCT input and output
	// Will also enable LogDebug.
	LogDebugWithInputAndOutput

	LogDebugAll LogLevelType = LogDebugWithRequest | LogDebugWithRequestBody | LogDebugWithRequestID |
		LogDebugWithEndpoint | LogDebugWithConfig | LogDebugWithSigning | LogDebugWithRequestRetries |
		LogDebugWithResponse | LogDebugWithResponseBody | LogDebugWithHTTPBody | LogDebugWithRequestErrors |
		LogDebugWithEventStreamBody | LogInfoWithInputAndOutput | LogDebugWithInputAndOutput
)

// A Logger is a minimalistic interface for the SDK to log messages to. Should
// be used to provide custom logging writers for the SDK to use.
type Logger interface {
	Log(...interface{})
	Debug(...interface{})
	DebugByLevel(LogLevelType, ...interface{})
	SetDebug(enable *bool)
	SetDebugLogLevel(*LogLevelType)
	Warn(...interface{})
	Error(...interface{})
	SetIoWriter(io.Writer)
}

// A LoggerFunc is a convenience type to convert a function taking a variadic
// list of arguments and wrap it so the Logger interface can be used.
type LoggerFunc func(...interface{})

// Log calls the wrapped function with the arguments provided
func (f LoggerFunc) Log(args ...interface{}) {
	f(args...)
}

// NewDefaultLogger returns a Logger which will write log messages to stdout, and
// use same formatting runes as the stdlib log.Logger
func NewDefaultLogger() Logger {
	v := LogDebugAll
	debug := false
	return &defaultLogger{
		logger:        log.New(os.Stdout, "", log.LstdFlags),
		debugLogger:   log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		warnLogger:    log.New(os.Stdout, "WARN: ", log.LstdFlags),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		debug:         &debug,
		debugLogLevel: &v,
	}
}

// A defaultLogger provides a minimalistic logger satisfying the Logger interface.
type defaultLogger struct {
	logger        *log.Logger
	debugLogger   *log.Logger
	warnLogger    *log.Logger
	errorLogger   *log.Logger
	debug         *bool
	debugLogLevel *LogLevelType
}

func (l *defaultLogger) SetIoWriter(writer io.Writer) {
	if writer != nil {
		l.logger.SetOutput(writer)
		l.debugLogger.SetOutput(writer)
		l.warnLogger.SetOutput(writer)
		l.errorLogger.SetOutput(writer)
	}
}

func (l *defaultLogger) SetDebugLogLevel(levelType *LogLevelType) {
	l.debugLogLevel = levelType
}

func (l *defaultLogger) SetDebug(debug *bool) {
	l.debug = debug
}

// Log logs the parameters to the stdlib logger. See log.Println.
func (l *defaultLogger) Log(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *defaultLogger) Debug(args ...interface{}) {
	if l.debug == nil || !*l.debug {
		l.Warn("Debug disabled.")
		return
	}
	l.debugLogger.Println(args...)
}

func (l *defaultLogger) DebugByLevel(inputLogLevel LogLevelType, args ...interface{}) {
	if l.debugLogLevel.Matches(inputLogLevel) {
		l.Debug(args...)
	}
}

func (l *defaultLogger) Warn(args ...interface{}) {
	l.warnLogger.Println(args...)
}

func (l *defaultLogger) Error(args ...interface{}) {
	l.errorLogger.Println(args)
}
