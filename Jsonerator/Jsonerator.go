/*=============================================================================================================================================================
= Copyright (c) 2023 John Rinderer. All rights reserved                                                                                                      =
=============================================================================================================================================================*/

package Jsonerator

import (
	"fmt"
	"unicode"
)

func (l *Lexer) peekChar(data []rune, posit int) string {
	return string(data[posit+1])
}

func (l *Lexer) peekRune(data []rune, posit int) rune {
	return data[posit+1]
}

func (l *Lexer) parseKeys(data []rune, posit int, t *Token) {
	var holder string
	size := len(data)

	for i := posit + 1; i < size; i++ {
		if string(data[i]) != ":" {
			holder += string(data[i])
		} else if string(data[i]) == ":" {
			break
		}
		posit++
	}
	t.Key = holder
	//we know after we parse a key we must be in a value
	l.PrevState = "key"
	l.State = "value"
	l.Posit = posit
	fmt.Println("The Key is: " + t.Key)
}

func (l *Lexer) parseArrayVals(data []rune, posit int, t *Token) {
	var holder string
	size := len(data)

	for i := posit + 1; i < size; i++ {
		//if the previous value is a key, then we know we're not in an array
		//need to check that the comma is outside an array
		if string(data[i]) != "]" {
			holder += string(data[i])
		} else {
			break
		}
		posit++
	}
	t.Value = holder
	//this will change based on how the value end!
	l.State = "key"
	l.PrevState = "value"
	l.Posit = posit
	fmt.Println("The Value is: " + t.Value)
}

func (l *Lexer) parseVals(data []rune, posit int, t *Token) {
	var holder string
	size := len(data)

	for i := posit + 1; i < size; i++ {
		//if the previous value is a key, then we know we're not in an array
		//need to check that the comma is outside an array
		if string(data[i]) != "," && string(data[i]) != "}" {
			holder += string(data[i])
		} else {
			break
		}
		posit++
	}
	t.Value = holder
	//this will change based on how the value end!
	l.State = "key"
	l.PrevState = "value"
	l.Posit = posit
	fmt.Println("The Value is: " + t.Value)
}

func GetKeyVals(data string) Tokens {
	var toks Tokens
	var tok Token
	var lex Lexer

	chars := []rune(data)

	size := len(chars)

	//reserved_chars := [7]string {"{","}","[","]","\"", ",",":"}
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"
	for i := 0; i < size; i++ {
		//first char will always be { and will always start with a key
		str := string(chars[i])
		if str == "{" { //we know this is a key
			lex.parseKeys(chars, i, &tok)
			i = lex.Posit
			//parsing simple JSON
		} else if str == ":" && lex.peekChar(chars, i) == "\"" && lex.State == "value" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
		} else if str == ":" && unicode.IsLetter((lex.peekRune(chars, i))) && lex.State == "value" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
		} else if str == ":" && unicode.IsNumber((lex.peekRune(chars, i))) && lex.State == "value" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
		} else if str == ":" && lex.peekChar(chars, i) == "[" && lex.peekChar(chars, i+1) == "{" && lex.State == "key" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
		} else if str == "[" && lex.peekChar(chars, i) == "{" && lex.PrevState == "key" {
			lex.parseKeys(chars, i+1, &tok)
			i = lex.Posit
		} else if str == "[" && lex.peekChar(chars, i) == "\"" && lex.PrevState == "key" {
			lex.parseArrayVals(chars, i, &tok)
			i = lex.Posit
		} else if str == "," && lex.peekChar(chars, i) == "\"" && lex.PrevState == "value" {
			lex.parseKeys(chars, i, &tok)
			i = lex.Posit
		}
	}

	return toks
}
