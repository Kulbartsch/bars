## Features implemented

* [X] Display labels, values and character charts - obviously. ;)
* [X] Handle UTF8 characters.
* [X] Values can be on the left or right side. Parse right with `-atEnd`.
* [X] Accept units and separators as part of the value like `12.90$;` with `-addNumChars="$;"`.
* [X] Alternatively accept a comma as a decimal separator `-comma`.
* [X] Define the number of decimals to display `-decimals=<n>`.
* [X] Display license information with `-about`.
* [X] Handle negative numbers.
* [X] Choose the 0 axis separator with `-zero=<t>`. (text)
* [X] Label will be shortened if there are not at least seven characters for the bar. (text)
* [X] Header line with texts for label `-labelHeader=<t>`, value `-valueHeader=<t>` and chart `-chartHeader=<t>`.
* [X] Header line separator/ruler, can be deactivated with `--noruler`. (text)
* [X] Nice UTF8 symbol support for output, can be deactivated with `-ascii`. (text)
* [X] Colored terminal output, `-mode=color` is default. Does only work on ANSI/VT100 compatible terminals. For the plain monochrome text output use `-mode=plain`. (text)
* [X] HTML snippet output for integration in a website `-mode=html`. (html)
* [X] Output CSS template with for integration in a website `-mode=css`. Does not parse anything. (html)
* [X] Full HTML page with title using `-mode=page`. (html)
* [X] self defined chart symbol (text)
* [X] print a sum `-sum`
* [X] Text-elements `-sum-label` below the label and `-sum-text` below the chart
* [X] print a count `-count`
* [X] Text-elements `-count-label` below the label and `-count-text` below the chart
* [X] print a average `-average`
* [X] Text-elements `-avg-label` below the label and `-avg-text` below the chart
* [X] Online-Manual


Elements marked with (text) are for text output only. \
Elements marked with (html) are for HTML output only.

## Ideas to implement

Don't take this list for granted. 
Some ideas might not be implemented, 
just what I (or you) like/want/need.

### WIP 



### Open 

* [ ] Fixed width of the label text
* [ ] Fixed Width of the value text
* [ ] define bar chart length, might be nice to define a 100% value etc.
  * [ ] define a maximum value bar length, if exceeded display `>`.
  * [ ] define a minimum value bar length, if below display `<`.
  * [ ] visual hint if zero axis not in visible range.
* [ ] chart for the "average" summary line
* [ ] Documentation
  * [ ] better online help `-help`?
  * [ ] more/update examples (COVID)
* [ ] Dynamic terminal width detection
* [ ] unix tool scripts for use with bars
  * [ ] du
  * [ ] df
  * [ ] top processes (CPU, Mem, IO)
  * [ ] web server logfile analyzer
  * [ ] covid19 scanner
  * [ ] system load (needs definition of chart length)
  * [ ] ping milliseconds latency
  * [ ] system temperature
* [ ] defining limit and visuals
  * [ ] define low and high limit `-limit-low` and `-limit-high`, comparing *value <= limit-low* and *value >= limit-high*. This results in 3 ranges: low, mid, high.
  * [ ] define warning `!` and alert `‼` if value  within one range.
  * [ ] define color of bars/background? if a value is within a range
* [ ] display scale respectively high and low values at end
* [ ] process two numbers per row for stacked charts like  *- Maybe this is overkill?*
  * [ ] ... for mixed charts
  * [ ] ... for ranges

### Ideas which are unlikely to be implemented 

* [ ] HTML output
  * [ ] header level for title (html)
  * [ ] dark mode output (html) 
* [ ] Use individual column separator, i.e. a semicolon to use the output as a CSV file. (text)
* [ ] Output format definition *- I think the way it is now is nice. YAGNI!*
* [ ] frame around the output (text)
* [ ] Output filename, Output to stdout is default. *- Isn't stdout all you need? Writing `-o file` or `> file` isn't such a difference.*
* HTML snippet output to integrate in a website (html)
  * [ ] inbound css (html) *- it's a snippet, that does not make sense*
  * [ ] choose color (background, text, bars) (html) *- you can/should modify the css template*
* [ ] autodetect if numbers are on the right or left of the text *- I must be bored to do so (Maybe it's easy: Just try to parse from left and right, count the successful ones. The most successful side wins.)*
* Colored terminal output (text)
  * [ ] coloring bars from green to red or other way round *- To die in beauty. ;)*
  * [ ] choose coloring
* [ ] sort input *- there are other tools for this*
* [ ] logarithmic scaled output
* [ ] `-combined` view of values within the bar charts. Could become challenging when chart elements are short. Inverted bars in text mode.
* [ ] i18n

## Resources

* HTML bars
  * http://www.coding-dude.com/wp/html5/bar-chart-html/
  * https://css-tricks.com/making-a-bar-chart-with-css-grid/
  * https://speckyboy.com/code-snippets-css3-bar-graphs/
* Terminal colors
  * https://en.wikipedia.org/wiki/ANSI_escape_code#Unix-like_systems
  * https://stackoverflow.com/questions/4842424/list-of-ansi-color-escape-sequences
  * https://misc.flogisoft.com/bash/tip_colors_and_formatting
* COVID
  * https://www.worldometers.info/coronavirus/worldwide-graphs/#daily-cases
* Alternative
  * [Termgraph](https://github.com/mkaz/termgraph) (Python program)
  
