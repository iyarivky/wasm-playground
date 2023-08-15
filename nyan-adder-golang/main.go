package main

import (
    "fmt"
    "syscall/js"
)

func FuncName(M int, durations []int) int {
    fmt.Println("[GO] Log file.go", M, durations)
    return 10
}

func main() {
    fmt.Println("Hello, WebAssembly!")

    // Mount the function on the JavaScript global object.
    js.Global().Set("FuncName", js.FuncOf(func(this js.Value, args []js.Value) any {
        if len(args) != 2 {
            fmt.Println("invalid number of args")
            return nil
        }
        if args[0].Type() != js.TypeNumber {
            fmt.Println("the first argument should be a number")
            return nil
        }

        arg := args[1]
        if arg.Type() != js.TypeObject {
            fmt.Println("the second argument should be an array")
            return nil
        }

        durations := make([]int, arg.Length())
        for i := 0; i < len(durations); i++ {
            item := arg.Index(i)
            if item.Type() != js.TypeNumber {
                fmt.Printf("the item at index %d should be a number\n", i)
                return nil
            }
            durations[i] = item.Int()
        }

        // Call the actual func.
        return FuncName(args[0].Int(), durations)
    }))

    // Prevent the program from exiting.
    // Note: the exported func should be released if you don't need it any more,
    // and let the program exit after then. To simplify this demo, this is
    // omitted. See https://pkg.go.dev/syscall/js#Func.Release for more information.
    select {}
}