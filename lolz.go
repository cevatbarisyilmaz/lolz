package main

import (
	"os"
	"lolz/reader"
	"lolz/lexer"
	"lolz/parser"
	"lolz/executor"
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