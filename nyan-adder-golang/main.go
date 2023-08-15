package main

import (
    "fmt"
    "syscall/js"
)


func nyanify(word string) string {
    result := word + "nyan"
    //fmt.Println(result)
    return result
}

func main() {
    fmt.Println("Hello, WebAssembly!")
    // Mount the function on the JavaScript global object.
    js.Global().Set("nyanify", js.FuncOf(func(this js.Value, args []js.Value) any {
        if len(args) != 1 {
            //fmt.Println("invalid number of args")
            return "Invalid no of arguments passed"
        }
        input := args[0].String()
        fmt.Printf("input %s\n", input)
        return nyanify(input)
    }))

    // Prevent the program from exiting.
    // Note: the exported func should be released if you don't need it any more,
    // and let the program exit after then. To simplify this demo, this is
    // omitted. See https://pkg.go.dev/syscall/js#Func.Release for more information.
    select {}
}