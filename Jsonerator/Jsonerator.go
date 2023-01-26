/*=============================================================================================================================================================
= Copyright (c) 2023 John Rinderer. All rights reserved                                                                                                      =
=============================================================================================================================================================*/

// Jsonerator.GetKeyVals("{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}")
package Jsonerator

import (
	"fmt"
	"strings"
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
	t.Key = strings.ReplaceAll(t.Key, "\"", "")
	//we know after we parse a key we must be in a value
	l.PrevState = "key"
	l.State = "value"
	l.Posit = posit
	//fmt.Println("The Key is: " + t.Key)
}

func (l *Lexer) parseSubVals(data []rune, posit int, t *Token) {
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
	t.Value = strings.ReplaceAll(t.Value, "\"", "")
	l.State = "key"
	l.PrevState = "value"
	l.Posit = posit
	//fmt.Println("The key is: " + t.Key)
	//fmt.Println("The Value is: " + t.Value)
}

func (l *Lexer) parseSubKeys(data []rune, posit int, t *Token, parentKey string) {
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
	t.Key = parentKey + "_" + holder

	//replace key stuff
	if strings.Count(t.Key, "\"") == 3 {
		t.Key = strings.Replace(t.Key, "\"_", "_", 3)
	} else {
		t.Key = strings.Replace(t.Key, "\"_\"", "_", 3)
	}
	t.Key = strings.ReplaceAll(t.Key, "\"", "")
	//

	//fmt.Println(t.Key)
	//we know after we parse a key we must be in a value
	l.PrevState = "key"
	l.State = "value"
	l.Posit = posit
	//fmt.Println("The Key is: " + t.Key)
}

// used to parse array of JSON
func (l *Lexer) parseKeysArray(data []rune, posit int, t *Token, ParentKey string, toks *Tokens) {
	var holder string

	size := len(data)
	count_of_pars_open := 1
	count_of_pars_clse := 0
	t.Key = ParentKey
	//consider writing to parse data until we get count of [ and closing ]
	for i := posit + 1; i < size; i++ {
		str := string(data[i])
		//we need to ensure we have closing ] for each open [
		if str == "[" {
			count_of_pars_open++
		}
		if str == "]" {
			count_of_pars_clse++
		}
		if count_of_pars_open == count_of_pars_clse {
			break
		}
		if str == "{" { //we know this is a key
			l.parseSubKeys(data, i, t, ParentKey)
			i = l.Posit
			//parsing simple JSON
		} else if str == ":" && l.peekChar(data, i) == "\"" && l.State == "value" {
			l.parseSubVals(data, i, t)
			i = l.Posit
			//fmt.Println(t)
			toks.Tokens = append(toks.Tokens, *t)
		} else if str == ":" && unicode.IsLetter((l.peekRune(data, i))) && l.State == "value" {
			l.parseSubVals(data, i, t)
			//fmt.Println(t)
			toks.Tokens = append(toks.Tokens, *t)
			i = l.Posit
		} else if str == ":" && unicode.IsNumber((l.peekRune(data, i))) && l.State == "value" {
			l.parseSubVals(data, i, t)
			//fmt.Println(t)
			toks.Tokens = append(toks.Tokens, *t)
			i = l.Posit
		} else if str == ":" && l.peekChar(data, i) == "[" && l.peekChar(data, i+1) == "{" && l.State == "key" {
			l.parseSubVals(data, i, t)
			//fmt.Println(t)
			toks.Tokens = append(toks.Tokens, *t)
			i = l.Posit
		} else if str == "[" && l.peekChar(data, i) == "\"" && l.PrevState == "key" {
			l.parseArrayVals(data, i, t)
			//fmt.Println(t)
			toks.Tokens = append(toks.Tokens, *t)
			i = l.Posit
		} else if str == "," && l.peekChar(data, i) == "\"" && l.PrevState == "value" {
			l.parseSubKeys(data, i, t, ParentKey)
			i = l.Posit
		} else if str == "\"" && string(data[i-1]) == "{" {
			l.parseSubKeys(data, i, t, ParentKey)
			i = l.Posit
		} else if str == "\"" {
			l.parseSubVals(data, i, t)
			i = l.Posit
			//fmt.Println(t)
			toks.Tokens = append(toks.Tokens, *t)
		}

		posit++
	}
	//fmt.Println(l.Posit)
	//fmt.Println(toks.Tokens)
	//fmt.Println(toks.Tokens)
	fmt.Println(holder)
	//t.Key += holder
	//we know after we parse a key we must be in a value
	l.PrevState = "key"
	l.State = "value"
	//l.Posit = posit
	//fmt.Println("The Key is: " + t.Key)
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
	t.Value = strings.ReplaceAll(t.Value, "\"", "")
	//this will change based on how the value end!
	l.State = "key"
	l.PrevState = "value"
	l.Posit = posit
	//fmt.Println("The Value is: " + t.Value)
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
	t.Value = strings.ReplaceAll(t.Value, "\"", "")
	//this will change based on how the value end!
	l.State = "key"
	l.PrevState = "value"
	l.Posit = posit
	//	fmt.Println("The Value is: " + t.Value)
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
			toks.Tokens = append(toks.Tokens, tok)
			//fmt.Println(tok.Key)
			//fmt.Println(tok.Value)
		} else if str == ":" && unicode.IsLetter((lex.peekRune(chars, i))) && lex.State == "value" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
			toks.Tokens = append(toks.Tokens, tok)
			//fmt.Println(tok.Key)
			//fmt.Println(tok.Value)
		} else if str == ":" && unicode.IsNumber((lex.peekRune(chars, i))) && lex.State == "value" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
			toks.Tokens = append(toks.Tokens, tok)
			//fmt.Println(tok.Key)
			//fmt.Println(tok.Value)
		} else if str == ":" && lex.peekChar(chars, i) == "[" && lex.peekChar(chars, i+1) == "{" && lex.State == "key" {
			lex.parseVals(chars, i, &tok)
			i = lex.Posit
			toks.Tokens = append(toks.Tokens, tok)
			//fmt.Println(tok.Key)
			//fmt.Println(tok.Value)
		} else if str == "[" && lex.peekChar(chars, i) == "{" && lex.PrevState == "key" {
			//this indicates we're in a new JSON and we need to preserve this top level key.
			//fmt.Println(tok.Key)
			//possibly pass toks array so we can append all necessary data
			lex.parseKeysArray(chars, i+1, &tok, tok.Key, &toks)
			//fmt.Println(tok.Value)
			i = lex.Posit
			//fmt.Println(tok.Key)
			//fmt.Println(tok.Value)
		} else if str == "[" && lex.peekChar(chars, i) == "\"" && lex.PrevState == "key" {
			lex.parseArrayVals(chars, i, &tok)
			i = lex.Posit

			//fmt.Println(tok.Key)
			//fmt.Println(tok.Value)
		} else if str == "," && lex.peekChar(chars, i) == "\"" && lex.PrevState == "value" {
			lex.parseKeys(chars, i, &tok)
			i = lex.Posit

		}

	}

	counter := 0
	duplicate_frequency := make(map[string]int)
	i := 0
	sizer := len(toks.Tokens)

	for a := 0; a < sizer; a++ {
		_, exist := duplicate_frequency[toks.Tokens[a].Key]

		if exist {
			val := fmt.Sprintf("%d", counter)
			toks.Tokens[a].Key += val
			duplicate_frequency[toks.Tokens[a].Key] += 1 // increase counter by 1 if already in the map
			counter++
		} else {
			duplicate_frequency[toks.Tokens[a].Key] = 1 // else start counting from 1
		}
		i++
	}

	//almost got it. Need to look at when tokens are populated
	return toks
}
