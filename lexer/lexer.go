package lexer

import "bufio"

const (
	o = iota
	l
	z
)

func Lex(reader *bufio.Reader) []int {
	tokens := make([]int, 0)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return tokens
		}
		switch r {
		case 'o':
			tokens = append(tokens, o)
		case 'l':
			tokens = append(tokens, l)
		case 'z':
			tokens = append(tokens, z)
		}
	}
}
