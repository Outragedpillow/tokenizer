package main

import (
    "fmt"
    "strings"
    "strconv"
)

type Token struct {
    Value string
    Type string
}

var TokenList []Token

func main() {
    str := "50.0 * 25 ="
    makeToken(str)
    fmt.Println(TokenList)
}


func makeToken(chars string) {
    words := strings.Split(chars, " ")

    for i := 0; i < len(words); i++ {
        switch words[i] {
            case "=":
                token := Token{
                    Value: words[i],
                    Type: "equalSign",
                }
                TokenList = append(TokenList, token)

            case "+", "-", "/", "*":
                token := Token{
                    Value: words[i],
                    Type: "operand",
                }
                TokenList = append(TokenList, token)

            case words[i]:
                if num, err := strconv.Atoi(words[i]); err == nil || (len(words[i]) > 1 && words[i][0] == '-' && strconv.Itoa(num) == words[i][1:]) {    
                 token := Token{
                    Value: words[i],
                    Type: "integer",
                }
                TokenList = append(TokenList, token)
                }

            default:
                fmt.Println("Invalid input")
        }
    }
}

func readinput() {
    
}
