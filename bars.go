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
	"strings"
	"unicode/utf8"
)

type parameters struct {
	comma       *bool   // use comma as decimal separator
	decimals    *int    // number of positions after decimal point for display of percentage. Default 0
	ascii       *bool   // ... if text-label length exceeds limit, "d" replace last three characters with 3 dots "..." or "e" the last char with unicode ellipsis (U+2026)  "…" (default)
	comment     *string // lines starting with <t> are ignored as they are comment lines. Default "#". Omitting <t> will define that there are no comments
	outputWidth *int    // limit text output to n chars. Default 80
	verbose     *bool   // verbose mode, default false
	valueAtEnd  *bool   // Values at the end, default at the beginning (false)
	addNumChars *string // additional number characters
	about       *bool   // write info about this program
	zero        *string // the symbol representing the 0 in text chart
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
	chartNLen   int
	chartPLen   int
	oneVal      float64
}

type chartDataType struct {
	value     float64
	valueText string
	label     string
}

var Description = "bars: generate a bar chart in the terminal or as HTML snippet"
var Copyright = "Copyright © 2021 Alexander Kulbartsch"
var License = "License: AGPL-3.0-or-later (GNU Affero General Public License 3 or later)"
var Version = "Version: v0.4.0"
var Source  = "Source: https://github.com/Kulbartsch/bars"

var myParam parameters
var chartData []chartDataType
var myValues = valuesType{0.0, 0.0, 0, 0, 0, 0, 0,
	0, 0, 0, 0,0.0}

func initialize() {
	//myParam.htmlOutput    = flag.Bool("html", false, "generate HTML snippet")
	myParam.decimals = flag.Int("decimals", 0, "number of decimals")
	myParam.comma = flag.Bool("comma", false, "use comma as decimal separator")
	myParam.comment = flag.String("comment", "#", "comment line start")
	myParam.outputWidth = flag.Int("outputWidth", 80, "width of the text output") // limit text output to n chars. Default 80% of width
	myParam.addNumChars = flag.String("addNumChars", "", "additional characters representing a number")
	myParam.verbose = flag.Bool("verbose", false, "print verbose parsing information")
	myParam.verbose = flag.Bool("v", false, "print verbose parsing information")
	myParam.valueAtEnd = flag.Bool("atEnd", false, "values are at the end of a line")
	myParam.ascii = flag.Bool("ascii", false, "use ascii dots instead of UTF8 ellipses")
	myParam.about = flag.Bool("about", false, "display information about this program, with this option other parameters are ignored")
	myParam.zero = flag.String("zero", "0", "symbol to represent the 0 line in text chart")
	flag.Parse()
}

func validateParameters() {
  if utf8.RuneCountInString(*myParam.zero) != 1 {
  	log.Fatal("Error: parameter 'zero' must be exactly one symbol")
  }
}

func about() {
	println(Description)
	println(Version)
	println(Copyright)
	println(License)
	println(Source)
	os.Exit(0)
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

// main function
func main() {
	initialize()
	if *myParam.about {
		about()
	}
	validateParameters()
	parseInput()
	calculateFormat()
	displayBars()
}

// EOF
