/*=============================================================================================================================================================
= Copyright (c) 2023 John Rinderer. All rights reserved                                                                                                      =
=============================================================================================================================================================*/

package Jsonerator

import "fmt"

type Tokens struct {
	Tokens []Token
}

type Token struct {
	Key   string
	Value string
}

func parseVal() {

}

func (t *Token) parseKey(data string, posit int) int {
	chars := []rune(data)
	var holder string
	size := len(chars)
	counter := posit

	for i := 0; i < size; i++ {
		if string(chars[i]) != ":" {
			holder += string(chars[i])
		}
		posit++
		counter = posit
	}
	t.Key = holder
	return counter
}

func GetKeyVals(data string) {
	//var key string
	//var val string
	//var inKey bool
	//var inVal bool
	var holder string
	var tokens Tokens
	var token Token
	var counter int
	counter = 0

	chars := []rune(data)

	size := len(chars)

	reserved_chars := make(map[string]string)
	//reserved_chars := [7]string {"{","}","[","]","\"", ",",":"}
	reserved_chars["{"] = "{"
	reserved_chars["}"] = "}"
	reserved_chars["["] = "["
	reserved_chars["]"] = "]"
	reserved_chars["\""] = "\""
	reserved_chars[","] = ","
	reserved_chars[":"] = ":"
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"
	for i := 0; i < size; i++ {
		fmt.Println(string(chars[i]))
		//build a token and compare
		holder += string(chars[i])
		switch holder {
		case "{":
			holder = ""
			//if this is the case we should expect a " next and a KEY
			fmt.Println("d")
		case "}":
			holder = ""
			//if this is the case we should expect either a , or the end of file
			fmt.Println("d")
		case "[":
			holder = ""
			//if this is the case we should expect a " to start some values seperated by comma
			//or { which should start parsing like the {
			fmt.Println("d")
		case "]":
			holder = ""
			//this should be followed by a , and end an array.
			fmt.Println("d")
		case "\"":
			holder = ""
			//this should be the start and end of all keys and values
			//we should expect a KEY or VALUE
			fmt.Println("d")
		case ",":
			holder = ""
			//this seperates our key valu pairs. It can be followed by a " or { or [
			fmt.Println("d")
		case ":":
			holder = ""
			//this should tell us we've ended a KEY and are starting a value
			//we will in a value here
			//this should be followed by either a " or a [
			fmt.Println("d")
		default:
			i = token.parseKey(holder, i)
			tokens.Tokens = append(tokens.Tokens, token)
			fmt.Println("DICKS")

		}
		counter++

	}
}
