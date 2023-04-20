package main

import "fmt"

type Token struct {
    tokeType TokenType
    lexeme   string
    literal  any // ?
    line     int
}

func NewToken(t TokenType, l string, literal any, line int) Token {
    return Token{
        tokeType: t,
        lexeme:   l,
        literal:  literal,
        line:     line,
    }
}

func (t Token) String() string {
    return fmt.Sprintf("%v - %s %+v", t.tokeType, t.lexeme, t.literal)
}
