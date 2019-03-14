package logs

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	Level = iota
)

var (
	sequenceNo uint64
	instance   *Logger
)

var logLevel = levelInfo

type Logger struct {
	mu     sync.Mutex
	output io.Writer
}

type LogRecord struct {
	ID string

	Level string

	Message string

	Filename string

	LineNo int
}

var Log = GetLogger(os.Stdout)

var (
	LogRecordTemplate      *template.Temlate
	debugLogRecordTemplate *template.Temlate
)

func GetLogger(w io.Writer) *BeeLogger {
	once.Do(func() {

	})
}
