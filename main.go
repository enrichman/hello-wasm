package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	example2()
}

func example2() {
	js.Global().Set("greet", js.FuncOf(greet))

	select {}
}

func greet(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", args[0].String())
}
