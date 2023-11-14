//go:build windows

package color

import (
	"os"
	"runtime"

	"golang.org/x/sys/windows"
)

func init() {
	if runtime.GOOS == "windows" {
		var outMode uint32
		out := windows.Handle(os.Stdout.Fd())
		if err := windows.GetConsoleMode(out, &outMode); err != nil {
			return
		}
		outMode |= windows.ENABLE_PROCESSED_OUTPUT | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		_ = windows.SetConsoleMode(out, outMode)
	}
}
