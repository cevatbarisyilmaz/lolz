package main

import (
	"github.com/cevatbarisyilmaz/lolz/executor"
	"github.com/cevatbarisyilmaz/lolz/lexer"
	"github.com/cevatbarisyilmaz/lolz/parser"
	"github.com/cevatbarisyilmaz/lolz/reader"
	"os"
)

func main() {
	defer func() {
		recover()
	}()
	if !(len(os.Args) > 1) {
		os.Exit(0)
	}
	buffer, err := reader.Read(os.Args[1])
	if err != nil {
		os.Exit(0)
	}
	lexicalTokens := lexer.Lex(buffer)
	parseNodes := parser.Parse(lexicalTokens)
	executor.Execute(parseNodes)
}
