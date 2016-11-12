package main

import (
	// "strings"
	"bytes"
	"fmt"
)

const (
	space = ' '
	slash = '\\'
)

var (
	quotechars = [3]rune{'`','"','\''}
	Parser = quoteparse	// optional: change to spaceparse
)

func strune (s string) (r rune) {
	return []rune(s)[0]
}

func quoteparse (s string) (args []string, err error) {
	/*
		Complexity: O(len(s))

		SYNTAX RULES:
		1. Arguments are space-separated. To type a literal space, escape with a backslash.
				Example:	`a b\ c` => [`a`, `b c`]
		2. After a space character ` `, a quote character begins a quote block.
			Quote characters are single quotes `'`, double quotes `"`, and backticks "`".
			Quote blocks end with the same quote character they began with.
				Example:	`a "b c" 'd"e'` => [`a`, `b c`, `d"e`]
		3. When not in a quote block, and not immediately following a space character, a quote character is literal.
				Example:	`a b" c'd"e` => [`a`, `b"`, `c'd"e`]
		4. Backslashes are escape characters, i.e. whatever character follows the backslash is interpreted literally.
			(1) Double backslash gets interpreted to a single backslash literal.
				Example:	`a\\b` => [`a\b`]
			(2) Backslash followed by a quote character, makes that quote character have no special significance.
				It cannot open or close quote blocks.
				Example:	`\"a 'b\' c'` => [`"a`, `b" c`]
			(3) Backslash followed by anything else, results in the backslash being ignored.
				If backslash is the last character of your string, it is ignored.
				Example:	`\a\  b\` => [`a `,  `b`]
	*/
	var buffer bytes.Buffer
	slashing := false
	quoting := false
	var curquote rune
	for i,e := range s {	// e is a rune
		_ = i
		if slashing {	// escape this character, whatever it is.
			buffer.WriteRune(e)
			slashing = false
			continue
		}
		// now we're not escaped
		if e == slash {	// escape the next character
			slashing = true
			continue
		}
		// now we don't have to worry about slashes at all
		if quoting {	// we're in a quote
			if e == curquote {	// end the quote
				quoting = false
				args = append(args, buffer.String())
				buffer.Reset()
				continue
			} else {	// normal character
				buffer.WriteRune(e)
				continue
			}
		}
		// now we're not in a quote
		if buffer.Len() == 0 {	// not in an arg; possibly starting a quote
			for _,q := range quotechars {
				if e == q {	// start a quote block
					quoting = true
					curquote = q
					break
				}
			}
			if !quoting {	// we didn't hit any quote characters
				if e != space {	// ignore spaces outside of args
					buffer.WriteRune(e)
					continue
				}
			}
		} else {	// we're in an arg
			if e == space {	// arg is over
				args = append(args, buffer.String())
				buffer.Reset()
				continue
			} else {	// normal character
				buffer.WriteRune(e)
				continue
			}
		}
	}
	// we're done with our string.
	if quoting {
		args = append(args, buffer.String())
		err = fmt.Errorf("Ended while still in a quote block.")
	}
	if buffer.Len() != 0 {	// we ended while still in an arg - this is expected.
		args = append(args, buffer.String())
	}
	return
}

func spaceparse (s string) (args []string, err error) {
	/*
		Complexity: O(len(s))

		SYNTAX RULES:
		1. Arguments are space separated.
		2. If you want a literal space character, you must escape it with a backslash `\ `.
		3. If you want a literal backslash, you must escape it with a backslash `\\`.
		4. Multiple consecutive (unescaped) spaces count as a single space. Leading and trailing spaces have no effect.

	*/
	var buffer bytes.Buffer
	slashing := false
	for i,e := range s {
		_ = i
		if slashing {	// escape this character, whatever it is.
			buffer.WriteRune(e)
			slashing = false
			continue
		}
		// we're not slashing
		if e == slash {	// escape the next character
			slashing = true
			continue
		}
		// nothing is slash-related now
		if e == space {
			if buffer.Len() != 0 {	// arg is over
				args = append(args, buffer.String())
				buffer.Reset()
			}
			continue
		}
		// now we're just a normal character
		buffer.WriteRune(e)
	}
	if buffer.Len() != 0 {
		args = append(args, buffer.String())
		buffer.Reset()
	}
	return
}