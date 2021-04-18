// bars: generate a bar chart in the terminal or as HTML snippet
// Copyright © 2021 Alexander Kulbartsch
// License: AGPL-3.0-or-later (GNU Affero General Public License 3 or later)
// Version: v0.1.0-beta

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
	"strconv"
	"strings"
	"unicode/utf8"
)

type parameters struct {
	//	htmlOutput  *bool   // ... generate html output
	//	htmlFormat  *rune   // ... "t" using text chars (default) or "h" html css bars
	//  style (HTML, colorful, plain)
	comma       *bool   // use comma as decimal separator
	decimals    *int    // number of positions after decimal point for display of percentage. Default 0
	//  labelWidth	*int	// maximum width of the label width
	//	asciiShort  *bool   // ... if text-label length exceeds limit, "d" replace last three characters with 3 dots "..." or "e" the last char with unicode ellipsis (U+2026)  "…" (default)
	//	format      *string // ...
	//	help        *bool   // ... print help
	//	fileName    *string // ... input filename, default stdin
	comment     *string // lines starting with <t> are ignored as they are comment lines. Default "#". Omitting <t> will define that there are no comments
	//	outputFile  *string // ... Output filename. Output to stdout is default.
	//	separator   *string // ...
	outputWidth *int    // limit text output to n chars. Default 80
	verbose     *bool   // verbose mode, default false
	valueAtEnd  *bool   // Values at the end, default at the beginning (false)
	AddNumChars *string // additional number characters
	// sort
	// print a sum
	// labelCaption
	// valueCaption
}

type valuesType struct {
	valueMin    float64
	valueMax    float64
	labelMinLen int
	labelMaxLen int
	valueTxtLen int
	lines       int
	linesValid  int
	labelLen    int
	chartLen    int
	oneVal      float64
}

type chartDataType struct {
	value     float64
	valueText string
	label     string
}

var Description = "bars: generate a bar chart in the terminal or as HTML snippet"
var Copyright = "© 2021 Alexander Kulbartsch"
var License = "AGPL-3.0-or-later (GNU Affero General Public License 3 or later)"
var Version = "V0.1.0-beta"

var myParam parameters
var chartData []chartDataType
var myValues = valuesType{0.0, 0.0, 0, 0, 0, 0, 0,
						  0, 0, 0.0}

func initialize() {
	//myParam.htmlOutput    = flag.Bool("html", false, "generate HTML snippet")
	myParam.decimals = flag.Int("decimals", 0, "number of decimals")
	myParam.comma = flag.Bool("comma", false, "use comma as decimal separator")
	myParam.comment = flag.String("comment", "#", "comment line start")
	myParam.outputWidth = flag.Int("outputWidth", 80, "width of the text output") // limit text output to n chars. Default 80% of width
	myParam.AddNumChars = flag.String("addNumChars", "", "additional characters representing a number")
	myParam.verbose = flag.Bool("verbose", false, "print verbose parsing information")
	myParam.verbose = flag.Bool("v", false, "print verbose parsing information")
	myParam.valueAtEnd = flag.Bool("atEnd", false, "values are at the end of a line")
	flag.Parse()
}

func validateParameters() {

	// -s<t>    separator definition or alternative string

	// -f       Output format for text. Default "tsvsb"="text, space, value, space, char-bar-chart",
	//          possible format options are:
	//	        * b    character based bar
	//	        * B    inverted bar without values
	//			* c    Comma
	//			* e    semicolon
	//			* i    pipe
	//       	* l    label
	//			* o    colon
	//	        * p    percent = value in percent
	//	        * P    inverted bar with containing percentage
	//			* s    space
	//			* t    tab
	//	        * v    value = the given value
	//	        * V    inverted bar with containing value
	//			* 1-9  any self defined

}

// WhiteSpaceTrim trims space, tabs and new lines
func WhiteSpaceTrim(in string) string {
	return strings.Trim(in, " \t\n")
}

// RemoveInvalidChars from text checking against set of valid chars
func RemoveInvalidChars(text, valid string) string {
	textLen := len(text)
	validLen := len(valid)
	if textLen == 0 {
		return text
	}
	if validLen == 0 {
		return ""
	}
	var result string
	for _, l := range text {
		letter := string(l)
		if strings.ContainsAny(letter, valid) {
			result = result + letter
		}
	}
	return result
}

//
func purifyNumber(numberText string, comma bool) string {
	var numberChars string
	if comma {
		numberChars = "0123456789+-,E"
	} else {
		numberChars = "0123456789+-.E"
	}
	re := RemoveInvalidChars(numberText, numberChars)
	if comma {
		re = strings.ReplaceAll(re, ",", ".")
	}
	return re
}

/*// NumberCharsLength checks string for the number of characters given by numChars and
// returns the slice values for start and end of the number and the length
func NumberCharsLength(text string, numChars string, fromRight bool) (start int, end int, leng int) {
	textLen := len(text)
	if textLen == 0 {
		return 0, 0, 0
	}
	runes := []rune(text)
	leng = 0
	if fromRight {
		for i := len(runes) - 1; i >= 0; i -= 1 {
			letter := string(runes[i])
			if !strings.ContainsAny(letter, numChars) {
				return i+1, textLen, leng
			}
			leng += len(letter)
		}
	} else {
		for i := 0; i < textLen; i += 1 {
			letter := string(runes[i])
			if !strings.ContainsAny(letter, numChars) {
				return 0, i, leng
			}
			leng += len(letter)
		}
	}
	// the whole string is a number
	return 0, textLen, textLen
}*/

// SplitLabelNumber separates the label from the value
func SplitLabelNumber(text string, numChars string, fromRight bool, comma bool) (label string, valueText string, value float64, err error) {
	if len(text) == 0 || len(numChars) == 0 {
		return text, "", 0, nil
	}
	runes := []rune(text)
	l := len(runes)
	var r rune
	var nt string // number text
	var lbl string // label
	isnum := true
	for i := 0; i < l; i += 1 {
		if fromRight {
			r = runes[l-i-1]
		} else {
			r = runes[i]
		}
		sr := string(r)
		if isnum {
			if strings.ContainsAny(sr, numChars) {
				if fromRight {
					nt = sr + nt
				} else {
					nt = nt + sr
				}
			} else { // no number char
				isnum = false
			}
		}
		if ! isnum {
			if fromRight {
				lbl = sr + lbl
			} else {
				lbl = lbl + sr
			}
		}
	}
	nv, err := strconv.ParseFloat(purifyNumber(nt, comma), 64)
	return WhiteSpaceTrim(lbl), nt, nv, err

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
	numberChars := "0123456789+-.,E" + *myParam.AddNumChars
	// numStart, numEnd, numCount := NumberCharsLength(text, numberChars, *myParam.valueAtEnd)
	label, valTxt, value, err := SplitLabelNumber(text, numberChars, *myParam.valueAtEnd, *myParam.comma)
	if len(valTxt) == 0 {
		if *myParam.verbose {
			log.Println("Line " + strconv.Itoa(myValues.lines) + " has no valid number.")
		}
		return
	}
	if err != nil {
		if *myParam.verbose {
			log.Println("Line " + strconv.Itoa(myValues.lines) + " Parse error:")
			log.Println(err)
		}
		return
	}
	valueText := fmt.Sprintf("%."+strconv.Itoa(*myParam.decimals)+"f", value)
	valTxtLen := len(valueText)
	// ...
	labelLength := utf8.RuneCountInString(label)
	if *myParam.verbose {
		log.Println("label: \"" + label + "\" value: " + strconv.FormatFloat(value, 'G', -1, 64))
	}
	chartData = append(chartData, chartDataType{value, valueText, label})
	myValues.linesValid += 1
	if myValues.linesValid == 1 {
		myValues.labelMaxLen = labelLength
		myValues.labelMinLen = labelLength
		myValues.valueMax = value
		myValues.valueMin = value
		myValues.valueTxtLen = valTxtLen
	} else {
		if myValues.labelMaxLen < labelLength {
			myValues.labelMaxLen = labelLength
		}
		if myValues.labelMinLen > labelLength {
			myValues.labelMinLen = labelLength
		}
		if myValues.valueMax < value {
			myValues.valueMax = value
		}
		if myValues.valueMin > value {
			myValues.valueMin = value
		}
		if myValues.valueTxtLen < valTxtLen {
			myValues.valueTxtLen = valTxtLen
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
		if *myParam.verbose {
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
	myValues.labelLen = myValues.labelMaxLen
	myValues.chartLen = *myParam.outputWidth - 2 - myValues.labelLen - myValues.valueTxtLen
	myValues.oneVal = myValues.valueMax / float64(myValues.chartLen)
	if *myParam.verbose {
		log.Println("max label length    : " + strconv.Itoa(myValues.labelLen))
		log.Println("max value length    : " + strconv.Itoa(myValues.valueTxtLen))
		log.Println("one bar char length : " + strconv.FormatFloat(myValues.oneVal, 'G', -1, 32))
	}
}

//
func displayBars() {
	for _, pair := range chartData {
		ll := utf8.RuneCountInString(pair.label)
		vl := utf8.RuneCountInString(pair.valueText)
		println(pair.label + strings.Repeat(" ", myValues.labelLen-ll+1) +
			strings.Repeat(" ", myValues.valueTxtLen-vl) + pair.valueText + " " +
			strings.Repeat("#", int(pair.value/myValues.oneVal)))
	}
}

// main function
func main() {
	initialize()
	validateParameters()
	parseInput()
	calculateFormat()
	displayBars()
}

// EOF
