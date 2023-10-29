package wathne

import (
	"fmt"
	"runtime"
)

const (
	NC     string = "\033[0;0m"
	Black  string = "\033[0;30m"
	Red    string = "\033[0;31m"
	Green  string = "\033[0;32m"
	Yellow string = "\033[0;33m"
	Blue   string = "\033[0;34m"
	Purple string = "\033[0;35m"
	Cyan   string = "\033[0;36m"
	White  string = "\033[0;37m"
)

// Initializes a safe buffered printer.
func NewPrinter(buffer int) (chan<- string) {
	var channel chan string = make(chan string, buffer)
	var receiver <-chan string = channel
	var sender chan<- string = channel
	go func(receiver <-chan string) {
		var str string
		for str = range receiver {
			fmt.Print(str)
		}
	}(receiver)
	return sender
}

func getCallerName(skip int) (string, bool) {
	var pc uintptr
	var ok bool
	pc, _, _, ok = runtime.Caller(skip)
	var fn *runtime.Func
	fn = runtime.FuncForPC(pc)
	if ok && fn != nil {
		return fn.Name(), true
	}
	return "", false
}

// fizzbizz
func getFizzBizzName(skip int) (string) {
	var callerName string
	var ok bool
	callerName, ok = getCallerName(skip)
	if !ok {
		return "function"
	}
	switch callerName {
	case "dat320/lab6/fizzbizz.(*SyncBlock).fizz":
		return "fizz"
	case "dat320/lab6/fizzbizz.(*SyncBlock).bizz":
		return "bizz"
	case "dat320/lab6/fizzbizz.(*SyncBlock).number":
		return "number"
	case "dat320/lab6/fizzbizz.(*SyncBlock).fizzBizz":
		return "fizzBizz"
	}
	return "function"
}

// fizzbizz
func getFizzBizzColor(name string) (string) {
	switch name {
	case "fizz":
		return Red
	case "bizz":
		return Green
	case "number":
		return Yellow
	case "fizzBizz":
		return Blue
	}
	return NC
}

// fizzbizz
func getFizzBizzNameAlignment(name string) (string) {
	switch name {
	case "fizz":
		return "    "
	case "bizz":
		return "    "
	case "number":
		return "  "
	case "fizzBizz":
		return ""
	}
	return ""
}

// fizzbizz
func LogFizzBizzRoutine(current int) (string) {
	var name string = getFizzBizzName(3)
	var color string = getFizzBizzColor(name)
	var spaces string = getFizzBizzNameAlignment(name)
	return fmt.Sprint("\n\n", color, current, " ", name, "()  ", spaces, NC)
}

// fizzbizz
func LogFizzBizzTrue(result string) (string) {
	var name string = getFizzBizzName(3)
	var color string = getFizzBizzColor(name)
	return fmt.Sprint(color, "true   ", result, NC)
}

// fizzbizz
func LogFizzBizzFalse(result string) (string) {
	var name string = getFizzBizzName(3)
	var color string = getFizzBizzColor(name)
	return fmt.Sprint(color, "false  ", result, NC)
}

// fizzbizz
func LogFizzBizzBreak() (string) {
	var name string = getFizzBizzName(3)
	var color string = getFizzBizzColor(name)
	return fmt.Sprint(color, "breaking... ", NC)
}

// fizzbizz
func LogFizzBizzDone() (string) {
	var name string = getFizzBizzName(3)
	var color string = getFizzBizzColor(name)
	return fmt.Sprint(color, "done", NC)
}
