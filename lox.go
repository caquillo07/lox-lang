package main

import (
    "bufio"
    "fmt"
    "os"
)

type Lox struct {
    hadErr  bool
    lastErr error
}

func (l Lox) runFile(file string) error {
    fmt.Println("reading file ", file)
    bytes, err := os.ReadFile(file)
    if err != nil {
        return err
    }
    return l.run(string(bytes))
}

func (l Lox) runPrompt() error {
    bufScanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("> ")
    for bufScanner.Scan() {
        txt := bufScanner.Text()
        if txt == "exit" {
            return nil
        }
        if err := l.run(txt); err != nil {
            return err
        }
        l.hadErr = false
    }
    return nil
}

func (l Lox) run(source string) error {
    scanner := NewScanner(source)
    tokens := scanner.scanTokens()

    for _, token := range tokens {
        fmt.Println(token)
    }
    return nil
}

func (l Lox) err(line int, message string) {
    l.report(line, "", message)
}

func (l Lox) report(line int, where string, message string) {
    if _, err := fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s", line, where, message); err != nil {
        panic(err)
    }
    l.hadErr = true
}
