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

Elements marked with (text) are for text output only. \
Elements marked with (html) are for HTML output only.

## Ideas to implement

Don't take this list for granted. 
Some ideas might not be implemented, 
just what I (or you) like/want/need.

* Documentation
  * [ ] better online help `-help`
  * [ ] manual `-manual`
  * [ ] more examples (COVID)
* [ ] maximum width of the label
* HTML output
  * [ ] header level (html)
  * [ ] dark mode output (html)
* [ ] Dynamic terminal width detection
* [ ] print a sum `-sum`
  * [ ] Text-elements `-sumlabel` below the label and `-sumtext` below the chart
* [ ] print a count `-count`
  * [ ] Text-elements `-countlabel` below the label and `-counttext` below the chart
* [ ] print a average `-average`
  * [ ] Text-elements `-avglabel` below the label and `-avgtext` below the chart
* [ ] define bar chart length, might be nice to define 100% etc.
  * [ ] define a maximum value bar length, if exceeded display `>`.
  * [ ] define a minimum value bar length, if below display `<`.
  * [ ] visual hint if zero axis not in visible range.
* unix tool scripts for use with bars
  * [ ] du
  * [ ] df
  * [ ] top processes (CPU, Mem, IO)
  * [ ] web server logfile analyzer
  * [ ] covid19 scanner
  * [ ] system load (needs definition of chart length)
  * [ ] ping milliseconds latency
* [ ] process two numbers per row for stacked charts like  *- Maybe this is overkill?*
  * [ ] ... for mixed charts 
  * [ ] ... for ranges
* defining limit and visuals
  * [ ] define low and high limit `-limitlow` and `-limithigh`, comparing *value <= limitlow* and *value >= limithigh*. This results in 3 ranges: low, mid, high.
  * [ ] define warning `!` and alert `‼` if value  within one range.
  * [ ] define color of bars/background? if a value is within a range

### Ideas which are unlikely to be implemented 

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
  