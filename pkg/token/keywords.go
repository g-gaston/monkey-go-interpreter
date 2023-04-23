package token

var keywords = map[string]Type{
	"let":    Let,
	"fn":     Function,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func IsKeyword(word string) (Type, bool) {
	t, ok := keywords[word]
	return t, ok
}
