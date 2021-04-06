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
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type parameters struct {
	htmlOutput  *bool
	htmlFormat  *rune // "t" using text chars (default) or "h" html css bars
	decimalChar *rune // -c<c>    Which character? Default "."
	decimals    *int  // -d<n>    number of positions after decimal point for display of percentage. Default 0
	textShorten *rune // -e<t>    if text exceeds limit, "d" replace last three characters with 3 dots "..." or
	//          "e" the last char with unicode ellipsis (U+2026)  "…" (default)
	format *string // -f       Output format for text. Default "tsb"="text, separator, char-bar-chart",
	//          possible format options are:
	//       	* t    text
	//	        * v    value = the given value
	//	        * p    percent = value in percent
	//	        * b    character based bar
	//	        * s    separator as defined
	//	        * V    inverted bar with containing value
	//	        * P    inverted bar with containing percentage
	//	        * B    inverted bar without values
	help     *bool   // -h       print help
	fileName *string // -i<t>    input filename, default stdin
	comment  *string // -m<t>    lines starting with <t> are ignored as they are comment lines. Default "#". Omitting <t> will define that there are no comments
	// -n       (reserved)
	outputFile *string // -o<t>    Output filename. Output to stdout is default.
	separator  *string // -s<t>    separator definition or alternative string
	//			*  s      space (default)
	//			*  t      tab
	//			*  c      Comma
	//			*  p      pipe
	//			*  e      semicolon
	//			*  o      colon
	//			*  other  any self defined
	outputWidth *int  // -t<n>    limit text output to n chars. Default 80% of width
	verbose     *bool // -v       verbose mode, default false
	debug       *bool
	valueAtEnd  *bool   // -z       Values at the end, default at the beginning (false)
	numChars    *string //
}

type valuesType struct {
	valueMin   float32
	valueMax   float32
	lineMinLen int
	lineMaxLen int
	lines      int
	linesValid int
	textLen    int
	chartLen   int
	oneVal     float32
}

type chartDataType struct {
	value float32
	label string
}

var myParam parameters
var chartData []chartDataType
var myValues = valuesType{0.0, 0.0, 0, 0, 0, 0, 0, 0, 0.0}

func init() {
	myParam.htmlOutput = flag.Bool("html", false, "generate HTML snippet")
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
	myParam.comment = flag.String("comment", "#", "comment line start")
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
	myParam.outputWidth = flag.Int("outputWidth", 80, "width of the text output") // limit text output to n chars. Default 80% of width
	//verbose    bool // -v       verbose mode, default false
	//valueAtEnd bool // -z       Values at the end, default at the beginning (false)
	myParam.numChars = flag.String("numChars", "0123456789+-.,_E", "characters representing a number")

	myParam.debug = flag.Bool("debug", false, "print debug information")

	flag.Parse()
}

// WhiteSpaceTrim trims space, tabs and new lines
func WhiteSpaceTrim(in string) string {
	return strings.Trim(in, " \t\n")
}

//
func RemoveInvalidChars(text, valid string) string {
	textLen := len(text)
	validLen := len(valid)
	if textLen == 0 || validLen == 0 {
		return text
	}
	var result string
	tr := []rune(text)
	for i := 0; i < textLen; i += 1 {
		letter := string(tr[i])
		if strings.ContainsAny(letter, valid) {
			result = result + letter
		}
	}
	return result
}

//
func purifyNumber(numberText string) string {
	numberChars := "0123456789+-.E"
	// TODO: Respect decimal comma as well, drop the sign from the numChars list before calling RemoveInvalidChars
	re := RemoveInvalidChars(numberText, numberChars)
	return re
}

//
func NumberCharsLenght(text string, numChars string) int {
	textLen := len(text)
	if textLen == 0 {
		return 0
	}
	runes := []rune(text)
	for i := 0; i < textLen; i += 1 {
		letter := string(runes[i])
		if !strings.ContainsAny(letter, *myParam.numChars) {
			return i
		}
	}
	return len(text)
}

func openStdinOrFile() io.Reader {
	var err error
	var r *os.File
	if len(flag.Args()) > 0 {
		r, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r = os.Stdin
	}
	return r
}

func parseLine(text string) {
	textLength := len(text)
	numCount := NumberCharsLenght(text, *myParam.numChars)
	if numCount == 0 {
		log.Println("Line " + strconv.Itoa(myValues.lines) + " has no valid number.")
		return
	}
	v64, err := strconv.ParseFloat(purifyNumber(text[0:numCount]), 32)
	value := float32(v64)
	if err != nil {
		log.Println(err)
	}
	label := WhiteSpaceTrim(text[numCount:])
	if *myParam.debug {
		log.Println("label: \"" + label + "\" value: " + strconv.FormatFloat(v64, 'G', -1, 32))
	}
	chartData = append(chartData, chartDataType{value, label})
	myValues.linesValid += 1
	if myValues.linesValid == 1 {
		myValues.lineMaxLen = textLength
		myValues.lineMinLen = textLength
		myValues.valueMax = value
		myValues.valueMin = value
	} else {
		if myValues.lineMaxLen < textLength {
			myValues.lineMaxLen = textLength
		}
		if myValues.lineMinLen > textLength {
			myValues.lineMinLen = textLength
		}
		if myValues.valueMax < value {
			myValues.valueMax = value
		}
		if myValues.valueMin > value {
			myValues.valueMin = value
		}
	}
}

func parseInput() {
	//filename := flag.Arg(0)
	r := openStdinOrFile()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		myValues.lines += 1
		text := WhiteSpaceTrim(scanner.Text())
		if strings.HasPrefix(text, *myParam.comment) {
			continue
		}
		if *myParam.debug {
			log.Println("Input Line: " + text)
		}
		parseLine(text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

//
func calculateFormat() {
	myValues.textLen = myValues.lineMaxLen
	myValues.chartLen = 80 - 1 - myValues.textLen
	myValues.oneVal = myValues.valueMax / float32(myValues.chartLen)
}

//
func displayBars() {
	for _, pair := range chartData {
		l := len(pair.label)
		println(pair.label + strings.Repeat(" ", myValues.textLen-l+1) +
			strings.Repeat("#", int(pair.value/myValues.oneVal)))
	}
}

//
func main() {
	parseInput()
	calculateFormat()
	displayBars()
}

// EOF
