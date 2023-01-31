package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) > 1 {
        fmt.Println("Usage: golox [script]")
        os.Exit(64)
    }

    var lox Lox
    if len(args) == 1 {
        if err := lox.runFile(args[0]); err != nil {
            os.Exit(65)
        }
    } else {
        if err := lox.runPrompt(); err != nil {
            os.Exit(65)
        }
    }
}
