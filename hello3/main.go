package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("goSum", js.FuncOf(sum))

	select {}
}

func sum(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return "error"
	}

	v1 := args[0].Float()
	v2 := args[1].Float()

	sum := v1 + v2

	js.Global().Get("window").Call("handleResult", sum)

	if sum > 1000 {
		js.Global().Get("window").Call("handleWarning", sum)
	} else {
		js.Global().Get("window").Call("handleWarning")
	}

	return sum
}
