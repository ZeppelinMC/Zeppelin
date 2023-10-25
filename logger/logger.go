package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/util"
	"github.com/fatih/color"
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

func getDateString() string {
	return gray(time.Now().Format("15:04:05"))
}

var gray = color.New(color.FgHiBlack, color.Bold).SprintFunc()
var blue = color.New(color.FgBlue, color.Bold).SprintFunc()
var cyan = color.New(color.FgCyan, color.Bold).SprintFunc()
var red = color.New(color.FgRed, color.Bold).SprintFunc()
var yellow = color.New(color.FgYellow, color.Bold).SprintFunc()

var colors = map[string]color.Attribute{
	"black":        color.FgBlack,
	"dark_blue":    color.FgBlue,
	"dark_green":   color.FgGreen,
	"dark_aqua":    color.FgCyan,
	"dark_red":     color.FgRed,
	"dark_purple":  color.FgMagenta,
	"gold":         color.FgYellow,
	"gray":         color.FgWhite,
	"dark_gray":    color.FgHiBlack,
	"blue":         color.FgHiBlue,
	"green":        color.FgHiGreen,
	"aqua":         color.FgHiCyan,
	"red":          color.FgHiRed,
	"light_purple": color.FgHiMagenta,
	"yellow":       color.FgHiYellow,
	"white":        color.FgHiWhite,
}

func parseChat(msg chat.Message) string {
	var str string
	texts := []chat.Message{msg}
	texts = append(texts, msg.Extra...)

	for _, text := range texts {
		attrs := []color.Attribute{colors[text.Color]}
		if text.Bold {
			attrs = append(attrs, color.Bold)
		}
		if text.Italic {
			attrs = append(attrs, color.Italic)
		}
		if text.Underlined {
			attrs = append(attrs, color.Underline)
		}
		str += color.New(attrs...).SprintFunc()(text.Text)
	}

	return str
}

func ParseChat(content string) string {
	content = strings.ReplaceAll(content, "§", "&")
	return parseChat(chat.NewMessage(content))
}

func (logger *Logger) Channel() chan Message {
	return logger.c
}

func (logger *Logger) Print(message string) {
	message = strings.ReplaceAll(message, "§", "&")
	msg := chat.NewMessage(message)
	logger.send(Message{
		Type:    "chat",
		Message: msg.String(),
	})
	fmt.Println(parseChat(msg))
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
	fmt.Printf("%s %s: %s\n", time, blue("INFO "), str)
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
	fmt.Printf("%s %s: %s\n", time, cyan("DEBUG"), str)
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
	fmt.Fprintf(os.Stderr, "%s %s: %s\n", time, red("ERROR"), str)
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
	fmt.Printf("%s %s: %s\n", time, yellow("WARN "), str)
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
