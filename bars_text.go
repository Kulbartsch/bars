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

type symbolType struct {
	headerFiller rune
	errors       rune
	ruler        rune
	zero         rune
	bar          rune
	exceedMark   string
}

var mySymbols symbolType

// MaxInt returns the bigger ont of to integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt returns the bigger ont of to integers
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateFormat() {
	// calculate width of separators
	separatorsWidth := 2
	// calculate number col width
	myValues.valueTxtLen = MaxInt(myValues.valueTxtLen, MaxInt(len(myValues.sumValText), MaxInt(len(myValues.cntValText), len(myValues.avgValText))))
	// check min width
	minWidth := separatorsWidth + MinInt(7, myValues.labelMaxLen) /* 7 = min label */ + 7 /* min bars */ + myValues.valueTxtLen
	if *myParam.outputWidth < minWidth {
		log.Fatal("Error: need min ", minWidth, " chars width, but there is only ", *myParam.outputWidth, " chars.")
	}
	// calculate label and bars length
	restWidth := *myParam.outputWidth - separatorsWidth - myValues.valueTxtLen
	myValues.labelLen = myValues.labelMaxLen
	myValues.chartLen = restWidth - myValues.labelLen
	if myValues.chartLen < 7 {
		myValues.chartLen = 7
		myValues.labelLen = restWidth - myValues.chartLen
		if myValues.labelLen < 2 {
			log.Fatal("Error: error calculating label/bars length. Label length: ", myValues.labelLen)
		}
	}
	calculateFooterLabels()
	// calculate bars-element size
	spread := myValues.valueMax
	if myValues.valueMin < 0 {
		spread = myValues.valueMax - myValues.valueMin
		myValues.oneVal = spread / float64(myValues.chartLen-1)
		myValues.chartNLen = int(-myValues.valueMin / myValues.oneVal)
		myValues.chartPLen = myValues.chartLen - myValues.chartNLen - 1
	} else { // Only positive values
		myValues.oneVal = spread / float64(myValues.chartLen)
		myValues.chartNLen = 0
		myValues.chartPLen = myValues.chartLen
	}
	if *myParam.verbose {
		log.Println("max label length    : " + strconv.Itoa(myValues.labelLen))
		log.Println("max value length    : " + strconv.Itoa(myValues.valueTxtLen))
		log.Println("chart length        : " + strconv.Itoa(myValues.chartLen))
		log.Println("... negative part   : " + strconv.Itoa(myValues.chartNLen))
		log.Println("... positive part   : " + strconv.Itoa(myValues.chartPLen))
		log.Println("one bar char length : " + strconv.FormatFloat(myValues.oneVal, 'G', -1, 32))
	}
}

func calculateFooterLabels() {
	if *myParam.sum {
		if myValues.labelLen < 3 {
			myValues.sumLabelText = "Σ"
		}
	}
	if *myParam.count {
		if myValues.labelLen < 5 && myValues.labelLen >= 3 {
			myValues.cntValText = "Cnt"
		} else if myValues.labelLen < 3 {
			myValues.cntValText = "#"
		}
	}
	if *myParam.average {
		if myValues.labelLen < 7 && myValues.labelLen >= 3 {
			myValues.cntValText = "Avg"
		} else if myValues.labelLen < 3 {
			myValues.cntValText = "⦵"
		}
	}
}

func FillText(text string, length int, filler rune, alignRight bool) string {
	var fill rune
	if utf8.RuneCountInString(string(filler)) != 1 {
		fill = ' '
	} else {
		fill = filler
	}
	ltx := utf8.RuneCountInString(text)
	if ltx >= length {
		return text // return the (to long) text
	}
	n := length - ltx
	if alignRight {
		return strings.Repeat(string(fill), n) + text
	}
	return text + strings.Repeat(string(fill), n)
}

func TextToLen(text string, length int, filler rune, alignRight bool, exceedMark string, exceedLeft bool, errorSymbol rune) string {
	var fill rune
	// pre checks
	if length < 1 { // do nothing
		return text
	}
	ltx := utf8.RuneCountInString(text)
	if length == 1 && ltx == 1 {
		return text
	}
	if utf8.RuneCountInString(string(filler)) != 1 {
		fill = ' '
	} else {
		fill = filler
	}
	lem := utf8.RuneCountInString(exceedMark)
	if ltx > length && lem > (length+1) {
		log.Println("Inter<nal Error: exceedMark greater than length")
		return TextToLen("", length, errorSymbol, false, "", false, errorSymbol)
	}
	// shorten text if necessary
	var result string
	if ltx > length {
		n := length - lem
		for _, x := range text {
			result = result + string(x)
			n -= 1
			if n == 0 {
				break
			}
		}
		if exceedLeft {
			result = exceedMark + result
		} else {
			result += exceedMark
		}
	} else {
		result = text
	}
	lre := utf8.RuneCountInString(result)
	if lre > length {
		log.Println("Internal Error: result exceeds required length")
		return result // return it anyway
	}
	// fill up
	return FillText(result, length, fill, alignRight)
}

func AnsiText(text string, format string) string {
	var f string
	if len(text) == 0 || len(format) == 0 {
		return text
	}
	switch format {
	case "bold":
		f = "1"
	case "title":
		f = "1" // bold
	case "footer":
		f = "3" // Itlaic
	case "underline":
		f = "4"
	case "header":
		f = "4;37" // underline, lightgrey
	case "lgreen", "value":
		f = "92"
	case "lblue", "label":
		f = "94"
	case "lmagenta", "positive":
		f = "95"
	case "lcyan", "negative":
		f = "96"
	case "zero":
		f = "39" // Default foreground color
	case "value_ul":
		f = "92;4"
	case "label_ul":
		f = "94;4"
	case "positive_ul":
		f = "95;4"
	case "negative_ul":
		f = "96;4"
	case "zero_ul":
		f = "39;4"
	default:
		f = format
	}
	return "\x1B[" + f + "m" + text + "\x1B[0m"
}

func colorize(text string, format string) string {
	if myValues.mode == "color" {
		return AnsiText(text, format)
	}
	return text
}

func displayTextRuler() {
	if !*myParam.noHR {
		filler := mySymbols.ruler
		fmt.Print(TextToLen("", myValues.labelLen, filler, false, "", false, mySymbols.errors) + " ")
		fmt.Print(TextToLen("", myValues.valueTxtLen, filler, false, "", false, mySymbols.errors) + " ")
		fmt.Println(TextToLen("", myValues.chartLen, filler, false, "", false, mySymbols.errors))
	}
}

func displayTextBarsHeader(exceedMark string) {
	if len(*myParam.title) > 0 {
		if myValues.mode == "color" {
			fmt.Println(AnsiText(TextToLen(*myParam.title, *myParam.outputWidth, ' ', false, exceedMark, false, mySymbols.errors), "title"))
		} else {
			fmt.Println(TextToLen(*myParam.title, *myParam.outputWidth, ' ', false, exceedMark, false, mySymbols.errors))
		}
	}
	if !myValues.headers {
		return
	}
	if myValues.mode == "color" {
		fmt.Print(AnsiText(TextToLen(*myParam.labelHeader, myValues.labelLen, ' ', false, exceedMark, false, mySymbols.errors), "underline") + " ")
		fmt.Print(AnsiText(TextToLen(*myParam.valueHeader, myValues.valueTxtLen, ' ', true, exceedMark, false, mySymbols.errors), "underline") + " ")
		fmt.Println(AnsiText(TextToLen(*myParam.chartHeader, myValues.chartLen, ' ', false, exceedMark, false, mySymbols.errors), "header"))
	} else {
		fmt.Print(TextToLen(*myParam.labelHeader, myValues.labelLen, mySymbols.headerFiller, false, exceedMark, false, mySymbols.errors) + " ")
		fmt.Print(TextToLen(*myParam.valueHeader, myValues.valueTxtLen, mySymbols.headerFiller, true, exceedMark, false, mySymbols.errors) + " ")
		fmt.Println(TextToLen(*myParam.chartHeader, myValues.chartLen, mySymbols.headerFiller, false, exceedMark, false, mySymbols.errors))
		displayTextRuler()
	}
}

func displayTextBarsOneFooter(lbl, val, txt string) {
	label := TextToLen(lbl, myValues.labelLen, ' ', false, mySymbols.exceedMark, false, mySymbols.errors)
	value := TextToLen(val, myValues.valueTxtLen, ' ', true, mySymbols.exceedMark, false, mySymbols.errors)
	text := TextToLen(txt, myValues.chartLen, ' ', false, mySymbols.exceedMark, false, mySymbols.errors)
	fmt.Print(colorize(label, "footer") + " ")
	fmt.Print(colorize(value, "footer") + " ")
	fmt.Println(colorize(text, "footer"))
}

func displayTextBarsFooter() {
	if myValues.mode == "plain" {
		displayTextRuler()
	}
	if *myParam.sum {
		displayTextBarsOneFooter(myValues.sumLabelText, myValues.sumValText, "")
	}
	if *myParam.count {
		displayTextBarsOneFooter(myValues.cntLabelText, myValues.cntValText, "")
	}
	if *myParam.average {
		displayTextBarsOneFooter(myValues.avgLabelText, myValues.avgValText, "")
	}
}

func displayTextBars() {
	//var label string
	displayTextBarsHeader(mySymbols.exceedMark)
	ul := ""
	for i, pair := range chartData {
		if i == len(chartData)-1 && myValues.mode == "color" && (*myParam.sum || *myParam.count || *myParam.average) { // if color and there are footer lines, last line should be underlined
			ul = "_ul"
		}
		label := TextToLen(pair.label, myValues.labelLen, ' ', false, mySymbols.exceedMark, false, mySymbols.errors)
		value := TextToLen(pair.valueText, myValues.valueTxtLen, ' ', true, mySymbols.exceedMark, false, mySymbols.errors)
		fmt.Print(colorize(label, "label"+ul) + " ")
		fmt.Print(colorize(value, "value"+ul) + " ")
		if myValues.valueMin < 0 {
			if pair.value < 0 {
				fmt.Print(strings.Repeat(" ", myValues.chartNLen+int(pair.value/myValues.oneVal)) +
					colorize(strings.Repeat(string(mySymbols.bar), int(-pair.value/myValues.oneVal)), "negative"+ul))
			} else {
				fmt.Print(strings.Repeat(" ", myValues.chartNLen))
			}
			fmt.Print(colorize(string(mySymbols.zero), "zero"+ul))
		}
		if pair.value > 0 {
			fmt.Println(colorize(strings.Repeat(string(mySymbols.bar), int(pair.value/myValues.oneVal)), "positive"+ul))
		} else {
			fmt.Println("")
		}
	}
	displayTextBarsFooter()
}

// EOF
