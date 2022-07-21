package main

import (
	"fmt"
	"syscall/js"

	"github.com/norunners/vert" // This is used to pass structs from Go to JS
)

type GoObj struct {
	Name string `js:"name"`
}

func hello(this js.Value, args []js.Value) interface{} {
	return "Hello " + args[0].String() + "!"
}

func main() {
	fmt.Println("Helloo!")
	done := make(chan bool)

	js.Global().Set("wasm", vert.ValueOf(GoObj{Name: "Thila"}))
	js.Global().Set("hello", js.FuncOf(hello))
	// js.Global().Set("wasm", js.FuncOf(hello))
	<-done // used to keep the program running.
}
