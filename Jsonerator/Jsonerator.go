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

func (t *Token) parseVal(data []rune, posit int) int {

	var holder string
	size := len(data)
	counter := posit

	for i := posit + 1; i < size; i++ {
		//fmt.Println(string(data[i]))
		if string(data[i]) == "[" || string(data[i]) == "{" {
			i++
		} else if string(data[i]) == "," {
			break
		} else if string(data[i]) != "," || string(data[i]) != "[" || string(data[i]) != "{" {
			//fmt.Println(string(data[i]))
			holder += string(data[i])
		}

		posit++
		counter = posit
	}
	t.Value = holder
	fmt.Println("The value is: " + t.Value)
	return counter
}

func (t *Token) parseKey(data []rune, posit int) int {

	var holder string
	size := len(data)
	counter := posit

	for i := posit + 1; i < size; i++ {
		if string(data[i]) == "[" || string(data[i]) == "{" {
			i++
		} else if string(data[i]) != ":" {
			holder += string(data[i])
		} else if string(data[i]) == ":" {
			break
		}
		posit++
		counter = posit
	}
	t.Key = holder
	fmt.Println("The Key is: " + t.Key)
	return counter
}

func nextChar(data []rune, curnt_posit int) int {
	//if the next char is a reserved char
	var next_car string
	var return_posit int

	return_posit = curnt_posit

	next_car = string(data[curnt_posit+1])

	if next_car == "{" {
		return_posit = curnt_posit + 2
	}
	return return_posit
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

	//reserved_chars := [7]string {"{","}","[","]","\"", ",",":"}
	//data = "{\"$implementationId\":\"deviceConfiguration--hardenedUncPathEnabled\",\"hardenedUncPaths\":[{\"serverPath\":\"\\\\\\\\*\\\\SYSVOL\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]},{\"serverPath\":\"\\\\\\\\*\\\\NETLOGON\",\"securityFlags\":[\"requireMutualAuthentication\",\"requireIntegrity\"]}]}"
	for i := 0; i < size; i++ {
		//fmt.Println(string(chars[i]))
		//build a token and compare
		holder += string(chars[i])
		fmt.Println(holder)
		switch holder {
		case "{":
			i = token.parseKey(chars, i)
			tokens.Tokens = append(tokens.Tokens, token)
			holder = ""
			//if this is the case we should expect a " next and a KEY
			//fmt.Println("d")
		case "}":

			//if this is the case we should expect either a , or the end of file
			//fmt.Println("d")
		case "[":
			//check to see if the next char is a { or the start of a string
			i = nextChar(chars, i)
			//holder = ""
			//if this is the case we should expect a " to start some values seperated by comma
			//or { which should start parsing like the {
			//fmt.Println("d")
		case "]":

			//this should be followed by a , and end an array.
			//fmt.Println("d")
		case "\"":
			
			//this should be the start and end of all keys and values
			//we should expect a KEY or VALUE
			//fmt.Println("d")
		case ",":
			//comma should in MOST cases be followed by a key
			i = token.parseKey(chars, i)
			tokens.Tokens = append(tokens.Tokens, token)
			holder = ""
			//this seperates our key valu pairs. It can be followed by a " or { or [
			//fmt.Println("d")
		case ":":
			i = token.parseVal(chars, i)
			tokens.Tokens = append(tokens.Tokens, token)
			holder = ""
			//this should tell us we've ended a KEY and are starting a value
			//we will in a value here
			//this should be followed by either a " or a [
			//fmt.Println("d")
		default:
			fmt.Println("DICKS")

		}
		counter++

	}
}
