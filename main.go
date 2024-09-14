package main

import (
	"C"
	"fmt"

	"github.com/fatih/color"

	"github.com/iawia002/lux/app"
)
import "unsafe"

func main() {}

// Convert a C array of strings to Go []string
func cArrayToGoStrings(argc C.int, argv **C.char) []string {
	length := int(argc)
	goArgs := make([]string, length)
	for i := 0; i < length; i++ {
		cStr := C.GoString(*argv)
		goArgs[i] = cStr
		argv = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + unsafe.Sizeof(*argv)))
	}
	return goArgs
}

//export RunApp
func RunApp(argc C.int, argv **C.char) C.int {
	goArgs := cArrayToGoStrings(argc, argv)
	if err := app.New().Run(goArgs); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
		return 1 // Failure
	}
	return 0 // Success
}
