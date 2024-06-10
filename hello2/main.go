package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("goHello", js.FuncOf(hello))

	select {}
}

func hello(this js.Value, args []js.Value) any {
	if len(args) == 0 || args[0].String() == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", args[0].String())
}
