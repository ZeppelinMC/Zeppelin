package logger

import (
	"fmt"
	"time"

	"github.com/dynamitemc/dynamite/util"
	"github.com/fatih/color"
)

type Logger struct {
	FilePath    string
	ConsoleText []string
}

func (logger *Logger) append(str string) {
	logger.ConsoleText = append(logger.ConsoleText, str)
	//web.Log(str)
}

func getDateString() string {
	return time.Now().Format("15:04:05")
}

func (logger *Logger) Info(format string, a ...interface{}) {
	blue := color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.append(fmt.Sprintf("[%s INFO]: %s", time, str))
	fmt.Printf("[%s %s]: %s\n", time, blue("INFO"), str)
}

func (logger *Logger) Print(format string, a ...interface{}) {
	format += "\n"
	logger.append(format)
	fmt.Printf(format, a...)
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	if !util.HasArg("-debug") {
		return
	}
	cyan := color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	str := fmt.Sprintf(format, a...)
	time := getDateString()
	logger.append(fmt.Sprintf("[%s DEBUG]: %s", time, str))
	fmt.Printf("[%s %s]: %s\n", time, cyan("DEBUG"), str)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	red := color.New(color.FgRed).Add(color.Bold).SprintFunc()

	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.append(fmt.Sprintf("[%s ERROR]: %s", time, str))
	fmt.Printf("[%s %s]: %s\n", time, red("ERROR"), str)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	yellow := color.New(color.FgYellow).Add(color.Bold).SprintFunc()

	time := getDateString()
	str := fmt.Sprintf(format, a...)
	logger.append(fmt.Sprintf("[%s WARN]: %s", time, str))
	fmt.Printf("[%s %s]: %s\n", time, yellow("WARN"), str)
}
