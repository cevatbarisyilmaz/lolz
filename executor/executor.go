package executor

import "lolz/parser"

func Execute(parseNodes []*parser.Node){
	scope := make(map[int]string)
	for _, node := range parseNodes{
		(*node).Execute(&scope)
	}
}