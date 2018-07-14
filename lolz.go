package main

import (
	"os"
	reader "github.com/cevatbarisyilmaz/lolz/reader"
	lexer "github.com/cevatbarisyilmaz/lolz/lexer"
	parser "github.com/cevatbarisyilmaz/lolz/parser"
	executor "github.com/cevatbarisyilmaz/lolz/executor"
)

func main() {
	if !(len(os.Args) > 1) {
		os.Exit(0)
	}
	buffer, err := reader.Read(os.Args[1])
	if err != nil{
		os.Exit(0)
	}
	lexicalTokens := lexer.Lex(buffer)
	parseNodes := parser.Parse(lexicalTokens)
	executor.Execute(parseNodes)
}
