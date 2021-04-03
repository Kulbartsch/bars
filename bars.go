// bars: generate a bar chart in the terminal or as HTML snippet
// Copyright © 2021 Alexander Kulbartsch
// License: AGPL-3.0-or-later (GNU Affero General Public License 3 or later)

/*
    This file is part of bars.

    bars is free software: you can redistribute it and/or modify it
    under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of
    the License, or any later version.

    bars is distributed in the hope that it will be useful, but
    WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public
	License along with bars.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type parameters struct {
	htmlOutput  *bool
	htmlFormat  rune // "t" using text chars (default) or "h" html css bars
	decimalChar rune // -c<c>    Which character? Default "."
	decimals    int  // -d<n>    number of positions after decimal point for display of percentage. Default 0
	textShorten rune // -e<t>    if text exceeds limit, "d" replace last three characters with 3 dots "..." or
	//          "e" the last char with unicode ellipsis (U+2026)  "…" (default)
	format string // -f       Output format for text. Default "tsb"="text, separator, char-bar-chart",
	//          possible format options are:
	//       	* t    text
	//	        * v    value = the given value
	//	        * p    percent = value in percent
	//	        * b    character based bar
	//	        * s    separator as defined
	//	        * V    inverted bar with containing value
	//	        * P    inverted bar with containing percentage
	//	        * B    inverted bar without values
	help         bool   // -h       print help
	fileName     string // -i<t>    input filename, default stdin
	outputLength int    // -l<n>    output length for text output  (Default terminal-width or 80 chars)
	comment      *string // -m<t>    lines starting with <t> are ignored as they are comment lines. Default "#". Omitting <t> will define that there are no comments
	// -n       (reserved)
	outputFile string // -o<t>    Output filename. Output to stdout is default.
	separator  string // -s<t>    separator definition or alternative string
	//			*  s      space (default)
	//			*  t      tab
	//			*  c      Comma
	//			*  p      pipe
	//			*  e      semicolon
	//			*  o      colon
	//			*  other  any self defined
	textLimit  int  // -t<n>    limit text output to n chars. Default 80% of width
	verbose    bool // -v       verbose mode, default false
	valueAtEnd bool // -z       Values at the end, default at the beginning (false)
}

var myParam parameters

func init() {
	myParam.htmlOutput =  flag.Bool("html", false, "generate HTML snippet")
	//decimalChar rune // -c<c>    Which character? Default "."
	//decimals    int  // -d<n>    number of positions after decimal point for display of percentage. Default 0
	//textShorten rune // -e<t>    if text exceeds limit, "d" replace last three characters with 3 dots "..." or
	////          "e" the last char with unicode ellipsis (U+2026)  "…" (default)
	//format string // -f       Output format for text. Default "tsb"="text, separator, char-bar-chart",
	////          possible format options are:
	////       	* t    text
	////	        * v    value = the given value
	////	        * p    percent = value in percent
	////	        * b    character based bar
	////	        * s    separator as defined
	////	        * V    inverted bar with containing value
	////	        * P    inverted bar with containing percentage
	////	        * B    inverted bar without values
	//help         bool   // -h       print help
	//fileName     string // -i<t>    input filename, default stdin
	//outputLength int    // -l<n>    output length for text output  (Default terminal-width or 80 chars)
	//comment      string // -m<t>    lines starting with <t> are ignored as they are comment lines. Default "#". Omitting <t> will define that there are no comments
	myParam.comment = flag.String("comment","#","comment line start")
	//// -n       (reserved)
	//outputFile string // -o<t>    Output filename. Output to stdout is default.
	//separator  string // -s<t>    separator definition or alternative string
	////			*  s      space (default)
	////			*  t      tab
	////			*  c      Comma
	////			*  p      pipe
	////			*  e      semicolon
	////			*  o      colon
	////			*  other  any self defined
	//textLimit  int  // -t<n>    limit text output to n chars. Default 80% of width
	//verbose    bool // -v       verbose mode, default false
	//valueAtEnd bool // -z       Values at the end, default at the beginning (false)
}

// WhiteSpaceTrim trims space, tabs and new lines
func WhiteSpaceTrim(in string) string {
	return strings.Trim(in, " \t\n")
}

// text StartsWith start
func StartsWith(text string, start string) bool {
	l := len(start)
	if len(text) < l {
		return false
	}
	if text[0:l] == start {
		return true
	}
	return false
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	return r
}

func main() {
	//filename := flag.Arg(0)
	r := openStdinOrFile()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := WhiteSpaceTrim(scanner.Text())
		if StartsWith(text, *myParam.comment) {
			continue
		}
		fmt.Println(text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// EOF