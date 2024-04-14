//go:build js && wasm

package main

import (
	"fmt"
	"math"
	"strconv"
	"syscall/js"
)

// The main entrypoint for our Wasm module.
// This is called by the `run` method of the `Go` JavaScript class provided by `wasm_exec.js`.
func main() {
	fmt.Printf("Go Wasm module instantiated!\n")

	// We can register Global functions on the browsers window object, allowing JavaScript code to call it.
	// Thereby interacting with our Wasm module - the interface provided is very simple, yet powerful.
	//
	// The actual logic for declaring and assigning these functions is handled by the underlying `wasm_exec.js`
	// file provided by the Go standard library.
	js.Global().Set("GoSqrt", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return math.NaN()
		}

		return math.Sqrt(parseFloatJS(args[0]))
	}))

	// The Wasm module should wait perpetually in our case in order to be able to
	// make calls to the exported functions.
	//
	// If a web application wants to dynamically instantiate and teardown Wasm modules then one could,
	// for example, create a JS exported func that would close the channel or any other wait mechanism.
	<-make(chan bool)
}

// parseFloatJS returns a float64 from a js.Value, based on either a `number` or `string` js type. NaN otherwise.
func parseFloatJS(v js.Value) float64 {
	switch v.Type() {
	case js.TypeNumber:
		return v.Float()
	case js.TypeString:
		if f, err := strconv.ParseFloat(v.String(), 64); err == nil {
			return f
		}
	}
	return math.NaN()
}
