package executor

import "github.com/cevatbarisyilmaz/lolz/parser"

func Execute(parseNodes []*parser.Node){
	scope := make(map[int]string)
	for _, node := range parseNodes{
		(*node).Execute(&scope)
	}
}
