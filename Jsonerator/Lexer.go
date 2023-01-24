/*=============================================================================================================================================================
= Copyright (c) 2023 John Rinderer. All rights reserved                                                                                                      =
=============================================================================================================================================================*/

package Jsonerator

/*
State is important. For KEY the process is simple. Once we've finished parsing a key
the state moves to VALUE. However VALUE parsing is slightly different.
The complexity arises in []. Inside [ can have a few scenarios.

The states would be

KEY
VALUE_SIMPLE
VALUE_ARRAY
VALUE_ARRY_JSON

Scenario 1
{
simple...json,
"Key":
[
	{
		"key1":"value1",
		"key2":"value2",
	},
	{
		"key3":"value3",
		"key4":"value4",
	}
]

In this situation our value is an array of JSON. While parsing KEY is easy parsing the value is difficult here.

In the above example Our KEY parser would flip the state to VALUE. We'd need to peak ahead 1 character and see if the
next char after the : is [. If so, we also need to check if the next char is {.

This would create a token like this.

Token
	Token.Key="Key"
	Token.Value="Key_key1-value1,Key_key2-value2,Key_key3-value3,Key_key4-value4"
*/

type Lexer struct {
	Posit     int
	PrevState string
	State     string
}
