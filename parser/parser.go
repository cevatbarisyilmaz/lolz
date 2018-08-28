package parser

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const MaxVariableLength = 8
const MaxTypeLength = 3
const MaxOperatorLength = 3
const MaxFunctionLength = 2
const MaxRuneTypeLength = 2

type Node interface {
	Execute(*map[int]string) string
}

const (
	Let = iota
	Operator
	Function
	Variable
	Value
	Loop
)

var Nodes = [...]int{Let, Operator, Function, Variable, Value, Loop}

const (
	Summation = iota
	Subtraction
	Multiplication
	Division
	Power
	IsEqual
	IsGreater
	IsSmaller
)

var Operators = [...]int{Summation, Subtraction, Multiplication, Division, Power, IsEqual, IsGreater, IsSmaller}

const (
	Print = iota
	ScanString
	ScanInteger
)

var Functions = [...]int{Print, ScanString, ScanInteger}

const (
	UpperCaseLetter = iota
	LowerCaseLetter
	Number
	Sign
)

var UpperCaseLetters = [...]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var LowerCaseLetters = [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var Numbers = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var Signs = [...]rune{' ', '\n', '.', ',', ':', '!', '?'}

type LetNode struct {
	left  *Node
	right *Node
}

func (node LetNode) Execute(scope *map[int]string) string {
	r := (*node.right).Execute(scope)
	if variableNode, ok := (*node.left).(VariableNode); ok {
		(*scope)[variableNode.reference] = r
	}
	return r
}

type OperatorNode struct {
	operation int
	left      *Node
	right     *Node
}

func (node OperatorNode) Execute(scope *map[int]string) string {
	l := (*node.left).Execute(scope)
	r := (*node.right).Execute(scope)
	o := Operators[node.operation%len(Operators)]
	li, err1 := strconv.Atoi(l)
	ri, err2 := strconv.Atoi(r)
	integer := false
	if err1 == nil && err2 == nil {
		integer = true
	}
	if o == IsEqual {
		if l == r {
			return "1"
		} else {
			return "0"
		}
	}
	if !integer {
		if o == Summation {
			return l + r
		}
		return ""
	}
	var a int
	switch o {
	case Summation:
		a = li + ri
	case Subtraction:
		a = li - ri
	case Multiplication:
		a = li * ri
	case Division:
		if ri != 0 {
			a = li / ri
		} else {
			a = 0
		}
	case Power:
		a = int(math.Pow(float64(li), float64(ri)))
	case IsGreater:
		if li > ri {
			a = 1
		} else {
			a = 0
		}
	case IsSmaller:
		if li < ri {
			a = 1
		} else {
			a = 0
		}
	}
	return strconv.Itoa(a)
}

type FunctionNode struct {
	function  int
	parameter *Node
}

func (node FunctionNode) Execute(scope *map[int]string) string {
	f := Functions[node.function%len(Functions)]
	var r string
	switch f {
	case Print:
		p := (*node.parameter).Execute(scope)
		fmt.Print(p)
	case ScanString:
		in := bufio.NewReader(os.Stdin)
		r, _ = in.ReadString('\n')
		r = r[:len(r)-1]
	case ScanInteger:
		var i int
		fmt.Scanf("%d", &i)
		r = strconv.Itoa(i)
	}
	return r
}

type VariableNode struct {
	reference int
}

func (node VariableNode) Execute(scope *map[int]string) string {
	return (*scope)[node.reference]
}

type ValueNode struct {
	value string
}

func (node ValueNode) Execute(scope *map[int]string) string {
	return node.value
}

type LoopNode struct {
	condition *Node
	body      []*Node
}

func (node LoopNode) Execute(scope *map[int]string) string {
	for (*node.condition).Execute(scope) == "1" {
		for _, b := range node.body {
			(*b).Execute(scope)
		}
	}
	return ""
}

func Parse(lexicalTokens []int) []*Node {
	nodes := make([]*Node, 0)
	var node *Node
	for len(lexicalTokens) != 0 {
		node, lexicalTokens = parse(lexicalTokens)
		nodes = append(nodes, node)
	}
	return nodes
}

func parseCode(lexicalTokens []int, maxLength int) ([]int, int) {
	num := 0
	i := 0
	for i < len(lexicalTokens) && i < maxLength {
		if lexicalTokens[i] == 2 {
			i++
			break
		}
		num *= 2
		num += lexicalTokens[i]
		i++
	}
	lexicalTokens = lexicalTokens[i:]
	return lexicalTokens, num
}

func parse(lexicalTokens []int) (*Node, []int) {
	var t Node
	t = ValueNode{value: ""}
	if len(lexicalTokens) == 0 {
		return &t, lexicalTokens
	}
	lexicalTokens, num := parseCode(lexicalTokens, MaxTypeLength)
	switch num {
	case Let:
		return parseLet(lexicalTokens)
	case Operator:
		return parseOperator(lexicalTokens)
	case Function:
		return parseFunction(lexicalTokens)
	case Variable:
		return parseVariable(lexicalTokens)
	case Value:
		return parseValue(lexicalTokens)
	case Loop:
		return parseLoop(lexicalTokens)
	}
	return &t, lexicalTokens
}

func parseLet(lexicalTokens []int) (*Node, []int) {
	var node Node
	left, lexicalTokens := parseVariable(lexicalTokens)
	right, lexicalTokens := parse(lexicalTokens)
	node = LetNode{left: left, right: right}
	return &node, lexicalTokens
}

func parseOperator(lexicalTokens []int) (*Node, []int) {
	var node Node
	lexicalTokens, num := parseCode(lexicalTokens, MaxOperatorLength)
	left, lexicalTokens := parse(lexicalTokens)
	right, lexicalTokens := parse(lexicalTokens)
	node = OperatorNode{operation: num, left: left, right: right}
	return &node, lexicalTokens
}

func parseFunction(lexicalTokens []int) (*Node, []int) {
	var node Node
	lexicalTokens, num := parseCode(lexicalTokens, MaxFunctionLength)
	var parameter *Node
	if num == Print {
		parameter, lexicalTokens = parse(lexicalTokens)
	}
	node = FunctionNode{function: num, parameter: parameter}
	return &node, lexicalTokens
}

func parseVariable(lexicalTokens []int) (*Node, []int) {
	var node Node
	lexicalTokens, num := parseCode(lexicalTokens, MaxVariableLength)
	node = VariableNode{reference: num}
	return &node, lexicalTokens
}

func parseValue(lexicalTokens []int) (*Node, []int) {
	var node Node
	value := ""
	var r rune
	for len(lexicalTokens) > 0 && lexicalTokens[0] != 2 {
		r, lexicalTokens = parseRune(lexicalTokens)
		value += string(r)
	}
	node = ValueNode{value: value}
	if len(lexicalTokens) == 0 {
		return &node, lexicalTokens
	}
	return &node, lexicalTokens[1:]
}

func parseRune(lexicalTokens []int) (rune, []int) {
	var target []rune
	lexicalTokens, num := parseCode(lexicalTokens, MaxRuneTypeLength)
	switch num {
	case UpperCaseLetter:
		target = UpperCaseLetters[:]
	case LowerCaseLetter:
		target = LowerCaseLetters[:]
	case Number:
		target = Numbers[:]
	case Sign:
		target = Signs[:]
	}
	length := 0
	for math.Pow(2, float64(length)) < float64(len(target)) {
		length++
	}
	lexicalTokens, num = parseCode(lexicalTokens, length)
	return target[num%len(target)], lexicalTokens
}

func parseLoop(lexicalTokens []int) (*Node, []int) {
	var node Node
	c, lexicalTokens := parse(lexicalTokens)
	nodes := make([]*Node, 0)
	var n *Node
	for len(lexicalTokens) > 0 && lexicalTokens[0] != 2 {
		n, lexicalTokens = parse(lexicalTokens)
		nodes = append(nodes, n)
	}
	node = LoopNode{condition: c, body: nodes}
	if len(lexicalTokens) == 0 {
		return &node, lexicalTokens
	}
	return &node, lexicalTokens[1:]
}
