package c6

import "github.com/stretchr/testify/assert"
import "testing"
import "fmt"
import "c6/ast"

func AssertLexerTokenSequenceFromState(t *testing.T, scss string, fn stateFn, tokenList []ast.TokenType) {
	fmt.Printf("Testing SCSS: %s\n", scss)
	var lexer = NewLexerWithString(scss)
	assert.NotNil(t, lexer)
	lexer.runFrom(fn)
	lexer.run()
	AssertTokenSequence(t, lexer, tokenList)
	lexer.close()
}

func AssertLexerTokenSequence(t *testing.T, scss string, tokenList []ast.TokenType) {
	fmt.Printf("Testing SCSS: %s\n", scss)
	var lexer = NewLexerWithString(scss)
	assert.NotNil(t, lexer)
	lexer.run()
	AssertTokenSequence(t, lexer, tokenList)
	lexer.close()
}

func OutputGreen(msg string, args ...interface{}) {
	fmt.Printf("\033[32m")
	fmt.Printf(msg, args...)
	fmt.Printf("\033[0m\n")
}

func OutputRed(msg string, args ...interface{}) {
	fmt.Printf("\033[31m")
	fmt.Printf(msg, args...)
	fmt.Printf("\033[0m\n")
}

func AssertTokenSequence(t *testing.T, l *Lexer, tokenList []ast.TokenType) []ast.Token {

	var tokens = []ast.Token{}
	var failure = false
	for idx, expectingToken := range tokenList {

		var token = <-l.Output

		if token == nil {
			failure = true
			OutputRed("not ok ---- got nil expecting %s", expectingToken.String())
			break
		}

		tokens = append(tokens, *token)

		if expectingToken == token.Type {
			// OutputGreen("ok %s '%s'", token.Type.String(), token.Str)
		} else {
			failure = true
			OutputRed("not ok ---- %d token => got %s '%s' expecting %s", idx, token.Type.String(), token.Str, expectingToken.String())
		}
		assert.Equal(t, expectingToken, token.Type)
	}

	if l.remaining() {
		var token *ast.Token = nil
		for token = <-l.Output; token != nil; token = <-l.Output {
			OutputRed("not ok ---- Remaining expecting %s '%s'", token.Type.String(), token.Str)
		}
	}
	if failure {
		t.Fatal()
	}

	return tokens
}

func AssertTokenType(t *testing.T, tokenType ast.TokenType, token *ast.Token) {
	assert.NotNil(t, token)
	if tokenType != token.Type {
		OutputRed("not ok - expecting %s. Got %s '%s'", tokenType.String(), token.Type.String(), token.Str)
	} else {
		OutputGreen("ok - expecting %s. Got %s '%s'", tokenType.String(), token.Type.String(), token.Str)
	}
	assert.Equal(t, tokenType, token.Type)
}
