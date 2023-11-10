package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/logger/color"
	"github.com/dynamitemc/dynamite/util"
)

type Message struct {
	Type    string `json:"type"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

type Logger struct {
	mu       sync.Mutex
	text     *strings.Builder
	chane    bool
	c        chan Message
	messages []Message
	file     *os.File
}

func Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Print(a...)
	fmt.Print("\n\r")
	return
}

func getDateString() string {
	return time.Now().Format("15:04:05")
}

var GB = color.Color{color.FgBlack, color.Bold}.Colorize
var BB = color.Color{color.FgBlue, color.Bold}.Colorize
var CB = color.Color{color.FgCyan, color.Bold}.Colorize
var RB = color.Color{color.FgRed, color.Bold}.Colorize
var YB = color.Color{color.FgYellow, color.Bold}.Colorize
var GG = color.Color{color.FgGreen, color.Bold}.Colorize

var HR = color.Color{color.FgRed}.Colorize
var C = color.Color{color.FgCyan}.Colorize

func (logger *Logger) Channel() chan Message {
	return logger.c
}

func (logger *Logger) Print(msg chat.Message) {
	time := getDateString()
	logger.send(Message{
		Type:    "chat",
		Message: msg.String(),
	})
	fmt.Printf("\r%s %s: %s\n\r> ", GB(time), GG("CHAT "), color.FromChat(msg))
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
	str = strings.ReplaceAll(str, "\n", "\n\r")
	fmt.Printf("\r%s %s: %s\n\r> ", GB(time), BB("INFO "), str)
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
	str = strings.ReplaceAll(str, "\n", "\n\r")
	fmt.Printf("\r%s %s: %s\n\r> ", GB(time), CB("DEBUG"), str)
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
	str = strings.ReplaceAll(str, "\n", "\n\r")
	fmt.Fprintf(os.Stderr, "\r%s %s: %s\n\r> ", GB(time), RB("ERROR"), str)
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
	str = strings.ReplaceAll(str, "\n", "\n\r")
	fmt.Printf("\r%s %s: %s\n\r> ", GB(time), YB("WARN "), str)
}

func (logger *Logger) EnableChannel() {
	logger.chane = true
	for _, m := range logger.messages {
		logger.c <- m
	}
}

func (logger *Logger) send(message Message) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	if logger.chane {
		logger.c <- message
	} else {
		logger.messages = append(logger.messages, message)
	}
}

func (logger *Logger) write(str string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
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
