package log

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

var timeColor = color.New(color.FgBlack).SprintFunc()
var infoColor = color.New(color.FgHiBlue, color.Bold).SprintFunc()
var errorColor = color.New(color.FgRed, color.Bold).SprintFunc()
var warningColor = color.New(color.FgYellow, color.Bold).SprintFunc()

func timeString() string {
	time := time.Now()

	return fmt.Sprintf("%02d:%02d:%02d", time.Hour(), time.Minute(), time.Second())
}

func Time() string {
	return timeColor(timeString())
}

var strs = sync.Pool{
	New: func() any { return new(strings.Builder) },
}

func SprintText(msg text.TextComponent) string {
	str := strs.Get().(*strings.Builder)
	str.Reset()
	defer strs.Put(str)

	var components = append([]text.TextComponent{msg}, msg.Extra...)
	for _, component := range components {
		c := colors[component.Color]
		if c == nil {
			c = colors["white"]
		}
		if component.Bold {
			c = c.Add(color.Bold)
		}
		if component.Strikethrough {
			c = c.Add(color.CrossedOut)
		}
		if component.Underlined {
			c = c.Add(color.Underline)
		}
		if component.Italic {
			c = c.Add(color.Italic)
		}
		str.WriteString(c.Sprint(component.Text))
	}

	return strings.ReplaceAll(str.String(), "\n", "\n\r")
}

var colors = map[string]*color.Color{
	"black":        color.New(color.FgBlack),
	"dark_blue":    color.New(color.FgBlue),
	"dark_green":   color.New(color.FgGreen),
	"dark_aqua":    color.New(color.FgCyan),
	"dark_red":     color.New(color.FgRed, color.Bold),
	"dark_purple":  color.New(color.FgMagenta),
	"gold":         color.New(color.FgYellow, color.Bold),
	"gray":         color.New(color.FgWhite),
	"dark_gray":    color.New(color.FgHiBlack),
	"blue":         color.New(color.FgHiBlue, color.Bold),
	"green":        color.New(color.FgGreen),
	"aqua":         color.New(color.FgHiCyan),
	"red":          color.New(color.FgHiRed),
	"light_purple": color.New(color.FgHiMagenta),
	"yellow":       color.New(color.FgHiYellow),
	"white":        color.New(color.FgHiWhite),
}

/*
Println prints the content prefixed and suffixed with a carriage return with an endline in the end.
Unlike fmt.Println, this doesn't add spaces between the elements

This should be used if raw terminal is enabled, but it works without it aswell
*/
func Println(v ...any) (i int, err error) {
	if i0, err := fmt.Print("\r"); err != nil {
		return i0, err
	} else {
		i += i0
	}

	if i0, err := fmt.Print(v...); err != nil {
		return i0, err
	} else {
		i += i0
	}
	i0, err := fmt.Println("\r")

	return i + i0, err
}

/*
Print prints the content prefixed and suffixed with a carriage return.

This should be used if raw terminal is enabled, but it works without it aswell
*/
func Print(v ...any) (i int, err error) {
	if i0, err := fmt.Print("\r"); err != nil {
		return i0, err
	} else {
		i += i0
	}

	i0, err := fmt.Print(v...)
	return i + i0, err
}

/*
Printf prints the content formatted, and prefixed and suffixed with a carriage return.

This should be used if raw terminal is enabled, but it works without it aswell
*/
func Printf(format string, v ...any) (i int, err error) {
	return fmt.Printf("\r"+format, v...)
}

/*
Println prints the content formatted, and prefixed and suffixed with a carriage return with an endline in the end.
Unlike fmt.Println, this doesn't add spaces between the elements

This should be used if raw terminal is enabled, but it works without it aswell
*/
func Printlnf(format string, v ...any) (i int, err error) {
	if i0, err := fmt.Print("\r"); err != nil {
		return i0, err
	} else {
		i += i0
	}

	if i0, err := fmt.Printf(format, v...); err != nil {
		return i0, err
	} else {
		i += i0
	}
	i0, err := fmt.Println("\r")

	return i + i0, err
}

// prints the contents prefixed by a carriage return + time, blue info text and suffixed with a newline and "> "
func Infoln(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), infoColor("INFO"))
	fmt.Println(v...)
	fmt.Print("\r> ")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a new line
func InfolnClean(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), infoColor("INFO"))
	fmt.Print(v...)
	fmt.Println("\r")
}

// prints the contents formatted prefixed by a carriage return + blue info text and suffixed with a new line
func InfolnfClean(format string, v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), infoColor("INFO"))
	fmt.Printf(format, v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text
func Info(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), infoColor("INFO"))
	fmt.Print(v...)
}

// prints the contents prefixed by a carriage return + blue info text
func Infof(format string, v ...any) {
	fmt.Printf("\r%s %s: %s", timeColor(timeString()), infoColor("INFO"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Infolnf(format string, v ...any) {
	fmt.Printf("\r%s %s: %s\n\r> ", timeColor(timeString()), infoColor("INFO"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Errorln(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), errorColor("ERROR"))
	fmt.Println(v...)
	fmt.Print("\r> ")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a new line
func ErrorlnClean(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), errorColor("ERROR"))
	fmt.Print(v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text
func Error(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), errorColor("ERROR"))
	fmt.Print(v...)
}

// prints the contents prefixed by a carriage return + blue info text
func Errorf(format string, v ...any) {
	fmt.Printf("\r%s %s: %s", timeColor(timeString()), errorColor("ERROR"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Errorlnf(format string, v ...any) {
	fmt.Printf("\r%s %s: %s\n\r> ", timeColor(timeString()), errorColor("ERROR"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Warnln(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), warningColor("WARN"))
	fmt.Println(v...)
	fmt.Print("\r> ")
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a new line
func WarnlnClean(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), warningColor("WARN"))
	fmt.Print(v...)
	fmt.Println("\r")
}

// prints the contents prefixed by a carriage return + blue info text
func Warn(v ...any) {
	fmt.Printf("\r%s %s: ", timeColor(timeString()), warningColor("WARN"))
	fmt.Print(v...)
}

// prints the contents prefixed by a carriage return + blue info text
func Warnf(format string, v ...any) {
	fmt.Printf("\r%s %s: %s", timeColor(timeString()), warningColor("WARN"), fmt.Sprintf(format, v...))
}

// prints the contents prefixed by a carriage return + blue info text and suffixed with a newline and "> "
func Warnlnf(format string, v ...any) {
	fmt.Printf("\r%s %s: %s\n\r> ", timeColor(timeString()), warningColor("WARN"), fmt.Sprintf(format, v...))
}
