package main

import (
	"fmt"
	"syscall/js"
)

func nyanify(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return "Invalid number of arguments"
	}

	input := p[0].String()
	result := input + "nyan"

	return js.ValueOf(result)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("nyanify", js.FuncOf(nyanify))

	<-c
}