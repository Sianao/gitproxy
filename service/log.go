package service

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type consoleColorModeValue int

const (
	autoColor consoleColorModeValue = iota
	disableColor
	forceColor
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

var consoleColorMode = autoColor

// LoggerConfig defines the config for Logger middleware.
type LoggerConfig struct {
	// Optional. Default value is gin.defaultLogFormatter
	Formatter LogFormatter

	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
	Output io.Writer

	// SkipPaths is an url path array which logs are not written.
	// Optional.
	SkipPaths []string

	// Skip is a Skipper that indicates which logs should not be written.
	// Optional.
}

// Skipper is a function to skip logs based on provided Context

// LogFormatter gives the signature of the formatter function passed to LoggerWithFormatter
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {

	// TimeStamp shows the time after the server returns a response.
	// StatusCode is HTTP response code.
	StatusCode    int
	ContentLength string
	// Latency is how much time the server cost to process a certain request.
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode

	switch {
	case code >= http.StatusContinue && code < http.StatusOK:
		return white
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (p *LogFormatterParams) MethodColor() string {
	method := p.Method

	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor indicates whether can colors be outputted to the log.
func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor)
}

func DefaultLogFormatter(param LogFormatterParams) {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	fmt.Printf("[Proxy] %s |%s %d %s| %13v | %15s |%s %-7s %s %#v %s\n",
		statusColor, time.Now().Format("01-02 15:04:05"), param.StatusCode, resetColor,
		param.ContentLength,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// DisableConsoleColor disables color output in the console.
func DisableConsoleColor() {
	consoleColorMode = disableColor
}

// ForceConsoleColor force color output in the console.
func ForceConsoleColor() {
	consoleColorMode = forceColor
}
