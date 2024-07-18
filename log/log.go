package log

import (
	"fmt"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue, color.Bold).SprintFunc()
var red = color.New(color.FgRed, color.Bold).SprintFunc()
var yellow = color.New(color.FgYellow, color.Bold).SprintFunc()

// this function should be used instead of fmt's because of the custom terminal
func Println(v ...any) {
	fmt.Print("\r")
	fmt.Print(v...)
	fmt.Println("\r")
}

// this function should be used instead of fmt's because of the custom terminal
func Print(v ...any) {
	fmt.Print("\r")
	fmt.Print(v...)
}

// this function should be used instead of fmt's because of the custom terminal
func Printf(format string, v ...any) {
	fmt.Print("\r")
	fmt.Printf(format, v...)
}

func Printlnf(format string, v ...any) {
	fmt.Print("\r")
	fmt.Printf(format, v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Infoln(v ...any) {
	fmt.Printf("\r%s: ", blue("INFO"))
	fmt.Println(v...)
	fmt.Print("\r> ")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a new line
func InfolnClean(v ...any) {
	fmt.Printf("\r%s: ", blue("INFO"))
	fmt.Print(v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text
func Info(v ...any) {
	fmt.Printf("\r%s: ", blue("INFO"))
	fmt.Print(v...)
}

// prints the contents prefixed by a carriage return + blue info text
func Infof(format string, v ...any) {
	fmt.Printf("\r%s: %s", blue("INFO"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Infolnf(format string, v ...any) {
	fmt.Printf("\r%s: %s\n\r> ", blue("INFO"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Errorln(v ...any) {
	fmt.Printf("\r%s: ", red("ERROR"))
	fmt.Println(v...)
	fmt.Print("\r> ")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a new line
func ErrorlnClean(v ...any) {
	fmt.Printf("\r%s: ", red("ERROR"))
	fmt.Print(v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text
func Error(v ...any) {
	fmt.Printf("\r%s: ", red("ERROR"))
	fmt.Print(v...)
}

// prints the contents prefixed by a carriage return + blue info text
func Errorf(format string, v ...any) {
	fmt.Printf("\r%s: %s", red("ERROR"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Errorlnf(format string, v ...any) {
	fmt.Printf("\r%s: %s\n\r> ", red("ERROR"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Warnln(v ...any) {
	fmt.Printf("\r%s: ", yellow("WARN"))
	fmt.Println(v...)
	fmt.Print("\r> ")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a new line
func WarnlnClean(v ...any) {
	fmt.Printf("\r%s: ", yellow("WARN"))
	fmt.Print(v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text
func Warn(v ...any) {
	fmt.Printf("\r%s: ", yellow("WARN"))
	fmt.Print(v...)
}

// prints the contents prefixed by a carriage return + blue info text
func Warnf(format string, v ...any) {
	fmt.Printf("\r%s: %s", yellow("WARN"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Warnlnf(format string, v ...any) {
	fmt.Printf("\r%s: %s\n\r> ", yellow("WARN"), fmt.Sprintf(format, v...))
}
