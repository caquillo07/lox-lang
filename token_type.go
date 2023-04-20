package main

type TokenType rune

const (
    // Single-character tokens.

    LeftParenTokenType TokenType = iota << 1
    RightParenTokenType
    LeftBraceTokenType
    RightBraceTokenType
    CommaTokenType
    DotTokenType
    MinusTokenType
    PlusTokenType
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

// String implements fmt.Stringer interface
//
// thank the lord for copilot, my new favorite tool.
func (t TokenType) String() string {
    switch t {
    case LeftParenTokenType:
        return "LeftParenTokenType"
    case RightParenTokenType:
        return "RightParenTokenType"
    case LeftBraceTokenType:
        return "LeftBraceTokenType"
    case RightBraceTokenType:
        return "RightBraceTokenType"
    case CommaTokenType:
        return "CommaTokenType"
    case DotTokenType:
        return "DotTokenType"
    case MinusTokenType:
        return "MinusTokenType"
    case PlusTokenType:
        return "PlusTokenType"
    case SemicolonTokenType:
        return "SemicolonTokenType"
    case SlashTokenType:
        return "SlashTokenType"
    case StarTokenType:
        return "StarTokenType"
    case BangTokenType:
        return "BangTokenType"
    case BangEqualTokenType:
        return "BangEqualTokenType"
    case EqualTokenType:
        return "EqualTokenType"
    case EqualEqualTokenType:
        return "EqualEqualTokenType"
    case GreaterTokenType:
        return "GreaterTokenType"
    case GreaterEqualTokenType:
        return "GreaterEqualTokenType"
    case LessTokenType:
        return "LessTokenType"
    case LessEqualTokenType:
        return "LessEqualTokenType"
    case IdentifierTokenType:
        return "IdentifierTokenType"
    case StringTokenType:
        return "StringTokenType"
    case NumberTokenType:
        return "NumberTokenType"
    case AndTokenType:
        return "AndTokenType"
    case ClassTokenType:
        return "ClassTokenType"
    case ElseTokenType:
        return "ElseTokenType"
    case FalseTokenType:
        return "FalseTokenType"
    case FunTokenType:
        return "FunTokenType"
    case ForTokenType:
        return "ForTokenType"
    case IfTokenType:
        return "IfTokenType"
    case NilTokenType:
        return "NilTokenType"
    case OrTokenType:
        return "OrTokenType"
    case PrintTokenType:
        return "PrintTokenType"
    case ReturnTokenType:
        return "ReturnTokenType"
    case SuperTokenType:
        return "SuperTokenType"
    case ThisTokenType:
        return "ThisTokenType"
    case TrueTokenType:
        return "TrueTokenType"
    case VarTokenType:
        return "VarTokenType"
    case WhileTokenType:
        return "WhileTokenType"
    case EOF:
        return "EOF"
    default:
        return "UnknownTokenType"
    }
}
