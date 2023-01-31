package main

type TokenType rune

const (
    // Single-character tokens.

    LeftParenTokenType TokenType = iota << 1
    RightParenTokenType
    LeftBraceTokenType
    RightBraceTokenType
    CommaTokenType
    OtTokenType
    MinusTokenType
    PluTokenType
    SemicolonTokenType
    SlashTokenType
    StarTokenType

    // One or two character tokens.

    BangTokenType
    BangEqualTokenType
    EqualTokenType
    EqualEqualTokenType
    GreaterTokenType
    GreaterEqualTokenType
    LessTokenType
    LessEqualTokenType

    // Literals.

    IdentifierTokenType
    StringTokenType
    NumberTokenType

    // Keywords.

    AndTokenType
    ClassTokenType
    ElseTokenType
    FalseTokenType
    FunTokenType
    ForTokenType
    IfTokenType
    NilTokenType
    OrTokenType
    PrintTokenType
    ReturnTokenType
    SuperTokenType
    ThisTokenType
    TrueTokenType
    VarTokenType
    WhileTokenType

    EOF
)
