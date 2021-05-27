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
	_ "embed"
	"fmt"
	"log"
	"strconv"
)

//go:embed "bars.css"
var barsCss string

const gridCols = 200

func displayCSS() {
	print(barsCss)
}

func calculateHtml() {
	// calculate bars-element size
	spread := myValues.valueMax
	if myValues.valueMin < 0 {
		spread = myValues.valueMax - myValues.valueMin
		myValues.oneVal = spread / float64(gridCols)
		myValues.chartNLen = int(-myValues.valueMin / myValues.oneVal)
		myValues.chartPLen = gridCols - myValues.chartNLen
	} else { // Only positive values
		myValues.oneVal = spread / float64(gridCols)
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

func displayHtmlSnippet() {
	exceedMark := mySymbols.exceedMark
	offsetX := 5
	offsetY := 2
	if len(*myParam.title) > 0 {
		fmt.Println("<h2>" + *myParam.title + "</h2>")
	}
	n := myValues.linesValid
	if myValues.headers {
		n += 1
	}
	fmt.Println("<div class=\"bars_chart\" style=\"grid-template-rows: repeat(" + strconv.Itoa(n) + ", 1fr);\">")
	if myValues.headers {
		fmt.Println("  <div class=\"bars_header\" style=\"grid-column: 1/2;\">" + TextToLen(*myParam.labelHeader, myValues.labelLen, mySymbols.headerFiller, false, exceedMark, false, mySymbols.errors) + "</div>")
		fmt.Println("  <div class=\"bars_header\"  style=\"grid-column: 3/4;text-align:right\">" + *myParam.valueHeader + "</div>")
		fmt.Println("  <div class=\"bars_header\" style=\"grid-column: 5/206; text-align:left\">" + *myParam.chartHeader + "</div>")
	}
	var m int
	for _, pair := range chartData {
		label := TextToLen(pair.label, myValues.labelLen, ' ', false, mySymbols.exceedMark, false, mySymbols.errors)
		value := pair.valueText
		fmt.Println("  <div class=\"bars_label\">" + label + "</div>")
		fmt.Println("    <div class=\"bars_value\">" + value + "</div>")
		if pair.value < 0 {
			m = offsetX + myValues.chartNLen + int(pair.value/myValues.oneVal)
			fmt.Println("    <div class=\"bars_neg\" style=\"grid-column:" + strconv.Itoa(m) + "/" + strconv.Itoa(offsetX+myValues.chartNLen) + ";\"></div>")
		}
		// nothing for pair.value = 0
		if pair.value > 0 {
			m = offsetX + myValues.chartNLen + 1 + int(pair.value/myValues.oneVal)
			fmt.Println("    <div class=\"bars_pos\" style=\"grid-column: " + strconv.Itoa(offsetX+myValues.chartNLen+1) + "/" + strconv.Itoa(m) + ";\"></div>")
		}
	}
	// zero line
	fmt.Println("  <div class=\"bars_zero\" style=\"grid-column:" + strconv.Itoa(offsetX+myValues.chartNLen) + "/" + strconv.Itoa(offsetX+myValues.chartNLen+1) + "; grid-row:" + strconv.Itoa(offsetY) + "/" + strconv.Itoa(offsetY+myValues.linesValid) + ";\"></div>")
	// footer
	// TODO: implement footer
	// end
	fmt.Println("</div>")
}

func displayHtmlPage() {
	fmt.Println(`<!DOCTYPE html>
<html lang="de">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Test CSS bars</title>
  <style>
    body {margin: 5px;
          color: #444;
          background-color: #eee;
          font-family: Helvetica,Arial,Verdana,sans-serif; }
  </style>
<style>`)
	fmt.Println(barsCss)
	fmt.Println(`</style>
</head>
<body>`)
	displayHtmlSnippet()
	fmt.Println("</body></html>")
}

// EOF
