package log

import (
	"fmt"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue, color.Bold).SprintFunc()
var red = color.New(color.FgRed, color.Bold).SprintFunc()
var yellow = color.New(color.FgYellow, color.Bold).SprintFunc()

// Infoln prints the content with a carriage return in the start and "> " in the end
func Infoln(v ...any) {
	fmt.Printf("\r%s: ", blue("INFO"))
	fmt.Println(v...)
	fmt.Print("> ")
}

// Infoln prints the content without a "> " in the end
func InfolnClean(v ...any) {
	fmt.Printf("\r%s: ", blue("INFO"))
	fmt.Println(v...)
}

func Info(v ...any) {
	fmt.Printf("\r%s: ", blue("INFO"))
	fmt.Print(v...)
}

func Infof(format string, v ...any) {
	fmt.Printf("\r%s: %s", blue("INFO"), fmt.Sprintf(format, v...))
}

// Errorln prints the content with a carriage return in the start and "> " in the end
func Errorln(v ...any) {
	fmt.Printf("\r%s: ", red("ERROR"))
	fmt.Println(v...)
	fmt.Print("> ")
}

func Error(v ...any) {
	fmt.Printf("\r%s: ", red("ERROR"))
	fmt.Print(v...)
}

func Errorf(format string, v ...any) {
	fmt.Printf("\r%s: %s", red("ERROR"), fmt.Sprintf(format, v...))
}

// Warnln prints the content with a carriage return in the start and "> " in the end
func Warnln(v ...any) {
	fmt.Printf("\r%s: ", yellow("WARN"))
	fmt.Println(v...)
	fmt.Print("> ")
}

func Warn(v ...any) {
	fmt.Printf("\r%s: ", yellow("WARN"))
	fmt.Print(v...)
}

func Warnf(format string, v ...any) {
	fmt.Printf("\r%s: %s", yellow("WARN"), fmt.Sprintf(format, v...))
}
