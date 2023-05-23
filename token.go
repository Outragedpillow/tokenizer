package main

import (
	"fmt"
	"unicode"
)

type Token struct {
	Value string
	Type  TokenType
}
type TokenType string

const (
	Equals      TokenType = "Equals"
	Number      TokenType = "Number"
	Plus        TokenType = "Plus"
	Minus       TokenType = "Minus"
	Divide      TokenType = "Divide"
	Times       TokenType = "Times"
	OpenParen   TokenType = "OpenParen"
	ClosedParen TokenType = "ClosedParen"
)

func main() {
	str := "50.0 * 25 ="
	tokens := Tokenize(str)
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func Tokenize(chars string) []Token {
	var TokenList []Token
	var number string = ""

	for _, char := range []byte(chars) {
		if unicode.IsDigit(rune(char)) {
			number += string(char)
			continue
		} else {
			if number != "" {
				token := Token{
					Value: number,
					Type:  Number,
				}
				TokenList = append(TokenList, token)
			}
			number = ""
		}
		switch char {
		case '=':
			token := Token{
				Value: string(char),
				Type:  Equals,
			}
			TokenList = append(TokenList, token)

		case '+':
			token := Token{
				Value: string(char),
				Type:  Plus,
			}
			TokenList = append(TokenList, token)

		case '-':
			token := Token{
				Value: string(char),
				Type:  Minus,
			}
			TokenList = append(TokenList, token)
		case '*':
			token := Token{
				Value: string(char),
				Type:  Times,
			}
			TokenList = append(TokenList, token)
		case '(':
			token := Token{
				Value: string(char),
				Type:  OpenParen,
			}
			TokenList = append(TokenList, token)
		case ')':
			token := Token{
				Value: string(char),
				Type:  ClosedParen,
			}
			TokenList = append(TokenList, token)

		default:
			fmt.Println("Invalid input")
		}
	}
	return TokenList
}
