package main

type Scanner struct {
    source  string
    tokens  []Token
    start   int
    current int
    line    int
}

func NewScanner(source string) Scanner {
    return Scanner{
        source: source,
        line:   1,
    }
}

func (s *Scanner) scanTokens() []Token {
    for !isAtEnd() {
        s.start = s.current
        s.scanToken()
    }
    s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
}

func isAtEnd() bool {

}

func (s *Scanner) scanToken() {

}
