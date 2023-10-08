package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dynamitemc/dynamite/util"
	"github.com/fatih/color"
)

type Logger struct {
	text string
	file *os.File
}

func getDateString() string {
	return time.Now().Format("15:04:05")
}

func (logger *Logger) Info(format string, a ...interface{}) {
	blue := color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.write(fmt.Sprintf("[%s INFO]: %s\n", time, str))
	fmt.Printf("[%s %s]: %s\n", time, blue("INFO"), str)
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	if !util.HasArg("-debug") {
		return
	}
	cyan := color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	str := fmt.Sprintf(format, a...)
	time := getDateString()
	logger.write(fmt.Sprintf("[%s DEBUG]: %s\n", time, str))
	fmt.Printf("[%s %s]: %s\n", time, cyan("DEBUG"), str)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	red := color.New(color.FgRed).Add(color.Bold).SprintFunc()

	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.write(fmt.Sprintf("[%s ERROR]: %s\n", time, str))
	fmt.Fprintf(os.Stderr, "[%s %s]: %s\n", time, red("ERROR"), str)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	yellow := color.New(color.FgYellow).Add(color.Bold).SprintFunc()

	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.write(fmt.Sprintf("[%s WARN]: %s\n", time, str))
	fmt.Printf("[%s %s]: %s\n", time, yellow("WARN"), str)
}

func (logger *Logger) write(str string) {
	logger.text += str
	t, _ := time.Parse("02-01-2006", strings.TrimSuffix(logger.file.Name(), ".log"))
	now := time.Now()
	if t.Day() != now.Day() {
		logger.reset()
	}
	logger.file.WriteString(logger.text)
}

func (logger *Logger) Close() {
	logger.file.Close()
}

func New() *Logger {
	os.Mkdir("log", 0755)
	file, _ := os.Create(fmt.Sprintf("log/%s.log", formatDay()))
	return &Logger{file: file}
}

func (logger *Logger) reset() {
	logger.file.Close()
	file, _ := os.Create(fmt.Sprintf("log/%s.log", formatDay()))
	logger.file = file
}

func formatDay() string {
	return time.Now().Format("02-01-2006")
}
