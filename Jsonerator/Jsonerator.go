/*=============================================================================================================================================================
= Copyright (c) 2023 John Rinderer. All rights reserved                                                                                                      =
=============================================================================================================================================================*/

package Jsonerator

func (l *Lexer) peekChar(data []rune, posit int) string {
	//fmt.Println(string(data[posit+1]))
	return string(data[posit+1])
}

func (l *Lexer) peekRune(data []rune, posit int) rune {
	return data[posit+1]
}

func (l *Lexer) ParseArray(data []rune, posit int, tks *Tokens) {

}

func (l *Lexer) BuildValues(data []rune, posit int, tk *Token) {
	var str_holder string
	size := len(data)

	//iterate through current position in our data
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"

	for i := posit; i < size; i++ {
		//store the character in string
		str_holder = string((data[i]))
		//this indicates the start of a key
		if str_holder == "{" && l.peekChar(data, (i+1)) == "\"" {

		}

	}
}

func (l *Lexer) BuildKey(data []rune, posit int, tk *Token) {
	var str_holder string
	var key_holder string
	counter := posit
	size := len(data)

	//iterate through current position in our data
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"

	for i := posit; i < size; i++ {
		//store the character in string
		str_holder = string((data[i]))
		//this indicates the start of a key
		if str_holder != ":" {
			key_holder += str_holder
		} else if str_holder == ":" {
			break
		}
		counter++
	}
	tk.Key = key_holder
	l.Posit = counter
}

func (l *Lexer) BuildToken(data []rune, posit int, tk *Token) {
	var str_holder string
	size := len(data)

	//iterate through current position in our data
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"

	for i := posit; i < size; i++ {
		//store the character in string
		str_holder = string((data[i]))
		//this indicates the start of a key
		if str_holder == "\"" && string(data[i-1]) == "{" {
			l.BuildKey(data, i, tk)
			i = l.Posit
			//this indicates the start of a simple KEY
		} else if string(data[i-1]) == ":" && str_holder == "\"" {
			l.BuildValues(data, i, tk)
			i = l.Posit
		}
	}

}

func (l *Lexer) ParseJson(data []rune, posit int, tks *Tokens) {
	//this will be the current token we're working with
	var tok_holder Token
	var str_holder string
	size := len(data)

	//iterate through current position in our data
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"

	for i := posit; i < size; i++ {
		//store the character in string
		str_holder = string((data[i]))
		//this indicates the start of a key
		if str_holder == "{" && l.peekChar(data, (i)) == "\"" {
			l.BuildToken(data, (i + 1), &tok_holder)
		} else if str_holder == ":" && l.peekChar(data, (i+1)) == "[" {
			l.ParseArray(data, i, tks)
		}

	}

}

func GetKeyVals(data string) Tokens {
	var toks Tokens
	/*
		var tok Token

	*/
	var lex Lexer

	chars := []rune(data)

	size := len(chars)

	//reserved_chars := [7]string {"{","}","[","]","\"", ",",":"}
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"
	for i := 0; i < size; i++ {
		lex.ParseJson(chars, i, &toks)
	}

	return toks
}

/*
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

*/
