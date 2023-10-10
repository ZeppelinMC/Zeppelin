package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/dynamitemc/dynamite/util"
	"github.com/fatih/color"
)

type Message struct {
	Type    string `json:"type"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

type Logger struct {
	text     *strings.Builder
	chane    bool
	c        chan Message
	messages []Message
	file     *os.File
}

func getDateString() string {
	return time.Now().Format("15:04:05")
}

var blue = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var cyan = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()
var yellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()

func (logger *Logger) Channel() chan Message {
	return logger.c
}

func (logger *Logger) Info(format string, a ...interface{}) {
	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.write(fmt.Sprintf("[%s INFO]: %s\n", time, str))
	logger.send(Message{
		Type:    "info",
		Time:    time,
		Message: str,
	})
	fmt.Printf("[%s %s]: %s\n", time, blue("INFO"), str)
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	if !util.HasArg("-debug") {
		return
	}
	str := fmt.Sprintf(format, a...)
	time := getDateString()
	logger.write(fmt.Sprintf("[%s DEBUG]: %s\n", time, str))
	logger.send(Message{
		Type:    "debug",
		Time:    time,
		Message: str,
	})
	fmt.Printf("[%s %s]: %s\n", time, cyan("DEBUG"), str)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.write(fmt.Sprintf("[%s ERROR]: %s\n", time, str))
	logger.send(Message{
		Type:    "error",
		Time:    time,
		Message: str,
	})
	fmt.Fprintf(os.Stderr, "[%s %s]: %s\n", time, red("ERROR"), str)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.write(fmt.Sprintf("[%s WARN]: %s\n", time, str))
	logger.send(Message{
		Type:    "warn",
		Time:    time,
		Message: str,
	})
	fmt.Printf("[%s %s]: %s\n", time, yellow("WARN"), str)
}

func (logger *Logger) EnableChannel() {
	logger.chane = true
	for _, m := range logger.messages {
		logger.c <- m
	}
}

func (logger *Logger) send(message Message) {
	if logger.chane {
		logger.c <- message
	} else {
		logger.messages = append(logger.messages, message)
	}
}

func (logger *Logger) write(str string) {
	logger.text.WriteString(str)
	t, _ := time.Parse("02-01-2006", strings.TrimSuffix(logger.file.Name(), ".log"))
	now := time.Now()
	if t.Day() != now.Day() {
		logger.reset()
	}
	logger.file.WriteString(logger.text.String())
}

func (logger *Logger) Close() {
	logger.file.Close()
}

func New() *Logger {
	os.Mkdir("log", 0755)
	file, err := os.Open(fmt.Sprintf("log/%s.log", formatDay()))
	text := &strings.Builder{}
	if err != nil {
		file, _ = os.Create(fmt.Sprintf("log/%s.log", formatDay()))
	} else {
		t, _ := io.ReadAll(file)
		text.Write(t)
		if text.Len() != 0 {
			text.WriteString("\n\n")
		}
	}
	return &Logger{file: file, text: text, c: make(chan Message, 1)}
}

func (logger *Logger) reset() {
	logger.file.Close()
	file, _ := os.Create(fmt.Sprintf("log/%s.log", formatDay()))
	logger.file = file
}

func formatDay() string {
	return time.Now().Format("02-01-2006")
}
