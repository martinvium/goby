package main

import (
	"fmt"
)

type Node interface{}
type DefNode struct {
	name     string
	argNames []string
	body     ExprNode
}
type ExprNode struct{}

type Parser struct {
	tokens []Token
}

func (p *Parser) consume(tokenSymbol string) Token {
	token := p.tokens[0]
	if token.Symbol == tokenSymbol {
		p.tokens = p.tokens[1:len(p.tokens)]
		return token
	} else {
		panic(fmt.Sprintf("Unexpected token %s, expected %s", token, tokenSymbol))
	}
}

func (p *Parser) parse() Node {
	defNode := p.parseDef()
	fmt.Println(defNode)
	return defNode
}

func (p *Parser) parseDef() DefNode {
	p.consume("def")
	name := p.consume("identifier").Value
	argNames := p.parseArgNames()
	body := p.parseExpr()
	p.consume("end")
	return DefNode{name, argNames, body}
}

func (p *Parser) parseArgNames() []string {
	p.consume("oparen")
	id := p.consume("identifier").Value
	p.consume("cparen")
	return []string{id}
}

func (p *Parser) parseExpr() ExprNode {
	integer := p.consume("integer").Value
	fmt.Println(integer)
	return ExprNode{}
}

func main() {
	fmt.Println("vim-go")

	tokens := tokenize("def a(x) 1 end")
	for _, token := range tokens {
		fmt.Println(token)
	}

	parser := &Parser{tokens}
	tree := parser.parse()
	fmt.Println(tree)
}
