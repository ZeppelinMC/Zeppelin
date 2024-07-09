package log

import (
	"fmt"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue, color.Bold).SprintFunc()
var red = color.New(color.FgRed, color.Bold).SprintFunc()

func Infoln(v ...any) {
	fmt.Printf("%s: ", blue("INFO"))
	fmt.Println(v...)
}

func Info(v ...any) {
	fmt.Printf("%s: ", blue("INFO"))
	fmt.Print(v...)
}

func Infof(format string, v ...any) {
	fmt.Printf("%s: %s", blue("INFO"), fmt.Sprintf(format, v...))
}

func Errorln(v ...any) {
	fmt.Printf("%s: ", red("ERROR"))
	fmt.Println(v...)
}

func Error(v ...any) {
	fmt.Printf("%s: ", red("ERROR"))
	fmt.Print(v...)
}

func Errorf(format string, v ...any) {
	fmt.Printf("%s: %s", red("ERROR"), fmt.Sprintf(format, v...))
}
