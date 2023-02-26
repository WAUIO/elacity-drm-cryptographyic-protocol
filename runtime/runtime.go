package runtime

import (
	"syscall/js"

	"github.com/elacity/crypto-protocol/log"
)

// FuncOf is a wrapper for js.FuncOf that recovers from panics.
func FuncOf(fn func(this js.Value, args []js.Value) any) js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) any {
			defer func() {
				if r := recover(); r != nil {
					log.Errorf("Recovered. Error: %s", r)
				}
			}()
			return fn(this, args)
		},
	)
}

type fn func(this js.Value, args []js.Value) (any, error)

var (
	jsErr     js.Value = js.Global().Get("Error")
	jsPromise js.Value = js.Global().Get("Promise")
)

// AsyncFunc is a promise wrapper that recovers from panics.
// see https://clavinjune.dev/en/blogs/golang-wasm-async-function/
func AsyncFunc(innerFunc fn) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(_ js.Value, promFn []js.Value) any {
			resolve, reject := promFn[0], promFn[1]

			go func() {
				defer func() {
					if r := recover(); r != nil {
						reject.Invoke(jsErr.New(r.(string)))
					}
				}()

				res, err := innerFunc(this, args)
				if err != nil {
					reject.Invoke(jsErr.New(err.Error()))
				} else {
					resolve.Invoke(res)
				}
			}()

			return nil
		})

		return jsPromise.New(handler)
	})
}
