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
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)


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
func PurifyNumber(numberText string, comma bool) string {
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


// SplitLabelNumber separates the label from the value
func SplitLabelNumber(text string, numChars string, fromRight bool, comma bool) (label string, valueText string, value float64, err error) {
	if len(text) == 0 || len(numChars) == 0 {
		return text, "", 0, nil
	}
	runes := []rune(text)
	l := len(runes)
	var r rune
	var nt string  // number text
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
		if !isnum {
			if fromRight {
				lbl = sr + lbl
			} else {
				lbl = lbl + sr
			}
		}
	}
	nv, err := strconv.ParseFloat(PurifyNumber(nt, comma), 64)
	return WhiteSpaceTrim(lbl), nt, nv, err
}


func parseLine(text string) {
	numberChars := "0123456789+-.,E" + *myParam.addNumChars
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

// EOF