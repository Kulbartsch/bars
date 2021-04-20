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
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

//
func calculateFormat() {
	// calculate width of separators
	separatorsWidth := 2
	// check min width
	minWidth := separatorsWidth + 7 /* min label */ + 7 /* min bars */ + myValues.valueTxtLen
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
		if myValues.labelLen < 7 {
			log.Fatal("Error: error calculating label/bars length. Label length: ", myValues.labelLen)
		}
	}
	// calculate bars-element size
	spread := myValues.valueMax
	if myValues.valueMin < 0 {
		spread = myValues.valueMax - myValues.valueMin
	}
	myValues.oneVal = spread / float64(myValues.chartLen)
	if *myParam.verbose {
		log.Println("max label length    : " + strconv.Itoa(myValues.labelLen))
		log.Println("max value length    : " + strconv.Itoa(myValues.valueTxtLen))
		log.Println("chart length        : " + strconv.Itoa(myValues.chartLen))
		log.Println("one bar char length : " + strconv.FormatFloat(myValues.oneVal, 'G', -1, 32))
	}
}

//
func displayBars() {
	var n int
	var label string
	for _, pair := range chartData {
		ll := utf8.RuneCountInString(pair.label)
		vl := utf8.RuneCountInString(pair.valueText)
		label = ""
		if ll > myValues.labelLen {
			if *myParam.ascii {
				n = myValues.labelLen - 3
			} else {
				n = myValues.labelLen - 1
			}
			for _, x := range pair.label {
				label = label + string(x)
				n -= 1
				if n == 0 {
					break
				}
			}
			if *myParam.ascii {
				label = label + "..."
			} else {
				label = label + "…"
			}
		} else {
			label = pair.label
		}
		ll = utf8.RuneCountInString(label)
		println(label + strings.Repeat(" ", myValues.labelLen-ll+1) +
			strings.Repeat(" ", myValues.valueTxtLen-vl) + pair.valueText + " " +
			strings.Repeat("#", int(pair.value/myValues.oneVal)))
	}
}

// EOF
