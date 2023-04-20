package main

var keywords = map[string]TokenType{
    "and":    AndTokenType,
    "class":  ClassTokenType,
    "else":   ElseTokenType,
    "false":  FalseTokenType,
    "for":    ForTokenType,
    "fun":    FunTokenType,
    "if":     IfTokenType,
    "nil":    NilTokenType,
    "or":     OrTokenType,
    "print":  PrintTokenType,
    "return": ReturnTokenType,
    "super":  SuperTokenType,
    "this":   ThisTokenType,
    "true":   TrueTokenType,
    "var":    VarTokenType,
    "while":  WhileTokenType,
}

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
    for !s.isAtEnd() {
        s.start = s.current
        s.scanToken()
    }
    s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
    return s.tokens
}

func (s *Scanner) isAtEnd() bool {
    return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
    c := s.advance()
    switch c {
    case '(':
        s.addToken(LeftParenTokenType)
    case ')':
        s.addToken(RightParenTokenType)
    case '{':
        s.addToken(LeftBraceTokenType)
    case '}':
        s.addToken(RightBraceTokenType)
    case ',':
        s.addToken(CommaTokenType)
    case '.':
        s.addToken(DotTokenType)
    case '-':
        s.addToken(MinusTokenType)
    case '+':
        s.addToken(PlusTokenType)
    case ';':
        s.addToken(SemicolonTokenType)
    case '*':
        s.addToken(StarTokenType)
    case '!':
        if s.match('=') {
            s.addToken(BangEqualTokenType)
        } else {
            s.addToken(BangTokenType)
        }
    case '=':
        if s.match('=') {
            s.addToken(EqualEqualTokenType)
        } else {
            s.addToken(EqualTokenType)
        }
    case '<':
        if s.match('=') {
            s.addToken(LessEqualTokenType)
        } else {
            s.addToken(LessTokenType)
        }
    case '>':
        if s.match('=') {
            s.addToken(GreaterEqualTokenType)
        } else {
            s.addToken(GreaterTokenType)
        }
    case '/':
        if s.match('/') {
            // A comment goes until the end of the line.
            for s.peek() != '\n' && !s.isAtEnd() {
                s.advance()
            }
        } else {
            s.addToken(SlashTokenType)
        }
    case ' ', '\r', '\t':
        // Ignore whitespace.
    case '\n':
        s.line++
    case '"':
        s.string()
    default:
        if s.isDigit(c) {
            s.number()
        } else if s.isAlpha(c) {
            s.identifier()
        } else {
            // todo temp
            Lox{}.err(s.line, "Unexpected character.")
        }
    }
}

func (s *Scanner) advance() rune {
    s.current++
    return rune(s.source[s.current-1])
}

func (s *Scanner) addToken(t TokenType) {
    s.addTokenLiteral(t, nil)
}

func (s *Scanner) addTokenLiteral(t TokenType, literal any) {
    text := s.source[s.start:s.current]
    s.tokens = append(s.tokens, NewToken(t, text, literal, s.line))
}

func (s *Scanner) match(expected rune) bool {
    if s.isAtEnd() {
        return false
    }
    if rune(s.source[s.current]) != expected {
        return false
    }
    s.current++
    return true
}

func (s *Scanner) peek() rune {
    if s.isAtEnd() {
        // return '\000'
        return '\x00'
    }
    return rune(s.source[s.current])
}

func (s *Scanner) string() {
    for s.peek() != '"' && !s.isAtEnd() {
        if s.peek() == '\n' {
            s.line++
        }
        s.advance()
    }
    // Unterminated string.
    if s.isAtEnd() {
        Lox{}.err(s.line, "Unterminated string.")
        return
    }
    // The closing ".
    s.advance()
    // Trim the surrounding quotes.
    value := s.source[s.start+1 : s.current-1]
    s.addTokenLiteral(StringTokenType, value)
}

func (s *Scanner) isDigit(c rune) bool {
    return c >= '0' && c <= '9'
}

func (s *Scanner) number() {
    for s.isDigit(s.peek()) {
        s.advance()
    }
    // Look for a fractional part.
    if s.peek() == '.' && s.isDigit(s.peekNext()) {
        // Consume the "."
        s.advance()
        for s.isDigit(s.peek()) {
            s.advance()
        }
    }

    // todo(hector) convert this value to float
    s.addTokenLiteral(NumberTokenType, s.source[s.start:s.current])
}

func (s *Scanner) peekNext() rune {
    if s.current+1 >= len(s.source) {
        return '\x00'
    }
    return rune(s.source[s.current+1])
}

func (s *Scanner) identifier() {
    for s.isAlphaNumeric(s.peek()) {
        s.advance()
    }
    text := s.source[s.start:s.current]
    t, ok := keywords[text]
    if ok {
        s.addToken(t)
    } else {
        s.addToken(IdentifierTokenType)
    }
}

func (s *Scanner) isAlpha(c rune) bool {
    return (c >= 'a' && c <= 'z') ||
        (c >= 'A' && c <= 'Z') ||
        c == '_'
}

func (s *Scanner) isAlphaNumeric(c rune) bool {
    return s.isAlpha(c) || s.isDigit(c)
}
