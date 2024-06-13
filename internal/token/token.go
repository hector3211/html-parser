package token

type TokenType int

const (
	DOCTYPE TokenType = iota
	START
	END
	COMMENT
	CHARACTER
	ENDOFFILE
)

type Token struct {
	Token          TokenType
	Literal        string
	Name           string           // DOCTYPE
	PublicID       string           // DOCTYPE
	SystemID       string           // DOCTYPE
	ForceQuirkFlag bool             // DOCTYPE default false
	TagName        string           // Start & End tags
	SelfClosingTag bool             // Start tag default false for "unset"
	Attributes     []AttributeToken // Start & End
	Data           string           // Comment & Character
}

type AttributeToken struct {
	Name  string
	Value string
}

func NewDoctypeToken() Token {
	return Token{
		Token:          DOCTYPE,
		Literal:        "DOCTYPE",
		Name:           "",
		PublicID:       "",
		SystemID:       "",
		ForceQuirkFlag: false,
	}
}

func NewStartToken(tagName string) Token {
	return Token{
		Token:          START,
		Literal:        "start",
		TagName:        tagName,
		SelfClosingTag: false,
		Attributes:     []AttributeToken{},
	}
}

func NewEndToken(tagName string) Token {
	return Token{
		Token:      END,
		Literal:    "end",
		TagName:    tagName,
		Attributes: []AttributeToken{},
	}
}

func NewCommentToken(comment string) Token {
	return Token{
		Token:   COMMENT,
		Literal: "comment",
		Data:    comment,
	}
}
func NewCharacterToken(character string) Token {
	return Token{
		Token:   CHARACTER,
		Literal: "character",
		Data:    character,
	}
}

func NewEndOfFileToken() Token {
	return Token{
		Token:   ENDOFFILE,
		Literal: "end-of-file",
	}
}
