## Ideas implemented

* [X] Handles UTF8 characters
* [X] Values can be on the left or right side
* [X] Accept units as part of the value like ``12.90$``
* [X] Alternatively accept a comma as a decimal separator
* [X] Define the number of decimals to display
* [X] Documentation
  * [X] license information (-about)## Ideas to implement
* [X] handle negative numbers
  * [X] choose the 0 axis separator (-zero=t) (text)
* [X] header line with texts for label, value and chart
* [X] header line separator (text)

## Ideas 

Don't take this list for granted. 
Some ideas might not be implemented, 
just what I (or you) like/want/need.

* [X] Documentation
  * [ ] better online help (-help)
  * [ ] more examples (COVID)
* [ ] maximum width of the label
* [ ] HTML snippet output to integrate in a website (html)
  * [ ] output CSS template (html)
  * [ ] full HTML page with title / subtitle (html)  
* [ ] Colored terminal output (text)
  * [ ] dynamic terminal width detection
* [ ] print a sum
* [ ] print a average
* [ ] print a count
* [ ] define bar chart length, might be nice to define 100% etc. (text)
* [ ] define a maximum value bar length, if exceeded display >
* [ ] define a minimum value bar length, if below display <
* [x] unix tool scripts for use with bars
  * [ ] du
  * [ ] df
  * [ ] top processes (CPU, Mem, IO)
  * [ ] web server logfile analyzer
  * [ ] covid19 scanner
  * [ ] system load

### Ideas which are unlikely to be implemented 

* [ ] Output format definition *- I think the way it is now is nice. YAGNI!*
* [ ] self defined column separator (text) *- and frames and ... maybe I'll die in options*
* [ ] Output filename, Output to stdout is default. *- Isn't stdout all you need? Writing `-o file` or `> file` isn't such a difference.*
* [x] HTML snippet output to integrate in a website (html)
  * [ ] inbound css (html) *- it's a snippet, that does not make sense*
  * [ ] choose color (background, text, bars) (html) *- you can/should modify the css template*
* [ ] autodetect if numbers are on the right or left of the text *- I must be bored to do so (Maybe it's easy: Just try to parse from left and right, count the successful ones. The most successful side wins.)*
* [x] Colored terminal output (text)
  * [ ] coloring bars from green to red or other way round *- To die in beauty. ;)*
  * [ ] choose coloring
* [ ] process two numbers per row for stacked or mixed charts *- Maybe this is overkill?*
* [ ] sort input *- there are other tools for this*

Elements marked with (text) are for text output only. \
Elements marked with (html) are for HTML output only.

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
  