package optionalstring

import (
	"fmt"

	parsec "github.com/prataprc/goparsec"
)

const (
	opensqr             = "OPEN_SQR"
	closesqr            = "CLOSE_SQR"
	single_quote        = "SINGLE_QUOTE"
	escaped_char        = "ESCAPED_CHAR"
	normalchars         = "NORMAL_CHARS"
	char                = "CHAR"
	chars               = "CHARS"
	char_within_escape  = "CHAR_WITHIN_ESCAPE"
	chars_within_escape = "CHARS_WITHIN_ESCAPE"
	escaped             = "ESCAPED"
	item                = "ITEM"
	items               = "ITEMS"
	optional_part       = "OPTIONAL_PART"
	optional_string     = "OPTIONAL_STRING"
)

var (
	opensqr_     parsec.Parser = parsec.Atom(`[`, opensqr)
	closesqr_                  = parsec.Atom(`]`, closesqr)
	squote_                    = parsec.Atom(`'`, single_quote)
	escapedchar_               = parsec.Token(`\\.`, escaped_char)
	normalchars_               = parsec.Token(`[^\[\]\\']+`, normalchars)
)

func newOptionalStringParser(ast *parsec.AST) parsec.Parser {
	char := ast.OrdChoice(char, nil, escapedchar_, normalchars_)
	chars := ast.Many(chars, nil, char)
	charWithinEscape := ast.OrdChoice(char_within_escape, nil, escapedchar_, normalchars_, opensqr_, closesqr_)
	charsWithinEscape := ast.Many(chars_within_escape, nil, charWithinEscape)

	var optional parsec.Parser
	escaped := ast.And(escaped, nil, squote_, charsWithinEscape, squote_)
	item := ast.OrdChoice(item, nil, chars, escaped, &optional)
	items := ast.Kleene(items, nil, item)
	optional = ast.And(optional_part, nil, opensqr_, items, closesqr_)
	return ast.Kleene(optional_string, nil, ast.OrdChoice("items", nil, optional, item))
}

type SyntaxError struct {
	Input    string
	ParsedAs string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf(
		"syntax error: maybe no opening/closing sqrt? parsed result = %s, input = %s",
		e.ParsedAs,
		e.Input,
	)
}

func EnumerateOptionalStringRaw(optionalString string) (enumerated []RawString, err error) {
	var node parsec.Queryable
	func() {
		defer func() {
			if rcv := recover(); rcv != nil {
				err = fmt.Errorf("%+v", rcv)
			}
		}()

		ast := parsec.NewAST("optionalString", 100)
		p := newOptionalStringParser(ast)
		s := parsec.NewScanner([]byte(optionalString))
		node, _ = ast.Parsewith(p, s)
	}()

	if err != nil {
		return
	}

	if parsedAs := node.GetValue(); len(parsedAs) != len(optionalString) {
		return []RawString{}, &SyntaxError{
			Input:    optionalString,
			ParsedAs: parsedAs,
		}
	}

	root := decode(node)

	return root.Flatten(), nil
}

func EnumerateOptionalString(optionalString string) (enumerated []string, err error) {
	raw, err := EnumerateOptionalStringRaw(optionalString)
	if err != nil {
		return []string{}, err
	}

	out := make([]string, len(raw))
	for idx, v := range raw {
		out[idx] = v.String()
	}
	return out, nil
}

func decode(node parsec.Queryable) *treeNode {
	root := &treeNode{}
	recursiveDecode(node.GetChildren(), root)
	return root
}

func recursiveDecode(nodes []parsec.Queryable, ctx *treeNode) {
	var onceFound bool

	for i := 0; i < len(nodes); i++ {
		if onceFound {
			recursiveDecode(nodes[i:], ctx.Right())
			return
		}

		switch nodes[i].GetName() {
		case optional_string:
			// skipping first node.
			recursiveDecode(nodes[i].GetChildren(), ctx)
		case optional_part:
			var optNext *treeNode
			if !onceFound {
				onceFound = true
				optNext = ctx.Left()
			} else {
				panic(
					fmt.Sprintf(
						"incorrect implementation: %s, %s",
						nodes[i].GetName(),
						nodes[i].GetValue(),
					),
				)
			}
			optNext.SetAsOptional()
			recursiveDecode(nodes[i].GetChildren(), optNext)
		case chars:
			for _, v := range nodes[i].GetChildren() {
				switch v.GetName() {
				case normalchars:
					ctx.AddValue(v.GetValue(), Normal)
				case escaped_char:
					ctx.AddValue(v.GetValue(), SingleQuoteEscaped)
				default:
					panic(fmt.Sprintf("incorrect implementation: %s, %s", v.GetName(), v.GetValue()))
				}
			}
		case escaped:
			ctx.AddValue(nodes[i].GetValue(), SingleQuoteEscaped)
		case items:
			recursiveDecode(nodes[i].GetChildren(), ctx)
		}
	}
}
