package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "text/scanner"
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
        fmt.Printf("> ")
    }
    return nil
}

func (l Lox) run(source string) error {
    s := newSourceScanner(strings.NewReader(source))
    for {
        token := s.Scan()
        if token == scanner.EOF {
            fmt.Println("got EOF, we done")
            return nil // todo(hector) remove?
        }
        fmt.Printf(">> %v - %q - %q\n", token, scanner.TokenString(token), s.TokenText())
    }
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

func newSourceScanner(r io.Reader) scanner.Scanner {
    var s scanner.Scanner
    s.Init(r)
    s.Mode = scanner.ScanChars |
        scanner.ScanFloats |
        scanner.ScanIdents |
        scanner.ScanInts |
        scanner.ScanStrings |
        scanner.ScanRawStrings |
        scanner.ScanComments
    s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
    return s
}
