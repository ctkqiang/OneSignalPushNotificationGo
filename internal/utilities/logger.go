package utilities

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type LogLevel int

const (
	APP_NAME = "VALUEFARM"
	VERSION  = "1.0.0"
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	VVERBOSE
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPink   = "\033[35m"
)

var (
	startTime       = time.Now()
	CurrentLevel    = INFO
	errorCallback   func(string)
	statusCallbacks = make(map[string]func(StatusUpdate))
	statusMutex     sync.RWMutex
)

type StatusUpdate struct {
	StockNo         string
	LastPrice       float64
	Volume          float64
	BestAskSize     float64
	BestAskCount    float64
	BestBidPrice    float64
	TimeReceived    string
	TimeReceivedISO string
	TotalRecords    string
	SequenceNumber  string
	CreatedAt       string
	Message         string
}

func RegisterErrorCallback(cb func(string)) {
	errorCallback = cb
}

func RegisterStatusCallback(id string, cb func(StatusUpdate)) {
	statusMutex.Lock()
	defer statusMutex.Unlock()
	statusCallbacks[id] = cb
}

func UnregisterStatusCallback(id string) {
	statusMutex.Lock()
	defer statusMutex.Unlock()
	delete(statusCallbacks, id)
}

func SetLogLevel(levelStr string) {
	level := strings.ToUpper(levelStr)
	switch level {
	case "DEBUG":
		CurrentLevel = DEBUG
	case "INFO":
		CurrentLevel = INFO
	case "WARN":
		CurrentLevel = WARN
	case "ERROR":
		CurrentLevel = ERROR
	case "VVERBOSE":
		CurrentLevel = VVERBOSE
	default:
		CurrentLevel = INFO
	}
}

// ToFloat64 converts any interface value (primarily string or numeric) to float64 safely.
func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		s := strings.TrimSpace(val)
		if s == "" || strings.EqualFold(s, "null") || strings.EqualFold(s, "none") {
			return 0
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0
		}
		return f
	default:
		return 0
	}
}

func init() {
	SetLogLevel(os.Getenv("LOG_LEVEL"))
}

func Log(level LogLevel, format string, a ...interface{}) {
	if level < CurrentLevel {
		return
	}

	levelStr := ""
	color := ""

	switch level {
	case DEBUG:
		levelStr = "DEBUG"
		color = colorYellow
	case INFO:
		levelStr = "INFO"
		color = colorBlue
	case WARN:
		levelStr = "WARN"
		color = colorPink
	case ERROR:
		levelStr = "ERROR"
		color = colorRed
	case VVERBOSE:
		levelStr = "VVERBOSE"
		color = "" // Default
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, a...)

	if level == ERROR && errorCallback != nil {
		errorCallback(msg)
	}

	if color != "" {
		fmt.Printf("["+APP_NAME+"] %s[%s] [%s] %s%s\n", color, timestamp, levelStr, msg, colorReset)
	} else {
		fmt.Printf("["+APP_NAME+"] [%s] [%s] %s\n", timestamp, levelStr, msg)
	}
}

func Info(format string, a ...interface{})     { Log(INFO, format, a...) }
func Debug(format string, a ...interface{})    { Log(DEBUG, format, a...) }
func Warn(format string, a ...interface{})     { Log(WARN, format, a...) }
func Error(format string, a ...interface{})    { Log(ERROR, format, a...) }
func VVerbose(format string, a ...interface{}) { Log(VVERBOSE, format, a...) }

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Mask(s string) string {
	runes := []rune(s)
	n := len(runes)

	if n <= 4 {
		return "****"
	}

	showCount := 10
	if n <= showCount {
		showCount = n / 3
	}

	return string(runes[:showCount]) + "[REDACTED]"
}

func CheckEnvFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		Error("CRITICAL: .env file NOT FOUND at: %s", filePath)
		return
	}

	Warn("Confirmed: .env file exists at: %s", filePath)
}

func CheckCUrrentMemory() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	toMB := func(bytes uint64) float64 {
		return float64(bytes) / 1024 / 1024
	}

	uptime := time.Since(startTime).Round(time.Second)
	numGoroutine := runtime.NumGoroutine()

	status := fmt.Sprintf(
		"--- [APP_STATE] Memory: Heap=%-7.2fMB | Total=%-7.2fMB | Sys=%-7.2fMB | G-Routines: %-4d | Uptime: %s (since %s)",
		toMB(m.Alloc),
		toMB(m.TotalAlloc),
		toMB(m.Sys),
		numGoroutine,
		uptime.String(),
		startTime.Format("2006-01-02 15:04:05"),
	)

	Log(INFO, "%s", status)
	return status
}
