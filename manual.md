# **bars** - Manual

*Easily generate bar charts from the command line as text or HTML*

## General Information

**bars** follows the Unix philosophy:
> Write programs that do one thing and do it well.  
> Write programs to work together.  
> Write programs to handle text streams, 
> because that is a universal interface.

**bars** takes a text file or *stdin* stream as the input data.
The result ist writen to *stdout*.
To save the result in a file use `> filename`.

The input text should have a numerical value at the beginning 
or end of a line. The not numerical data is the description
for this value.

There are many options to specify the parsing of the input,
as well as to control the formatting of the output.
This documentation shows how to use all this features.

If you are looking for more examples visit
https://github.com/Kulbartsch/bars/blob/main/examples/example.md

### About program flags (aka options)

Every flag starts with one dash "-" (or two). There is no difference
between short, one character option and long options.   

If an option expects a value, it can be given seperated by a 
white-space or a "=". 

Examples: 

    -decimals 2
    --decimals=2
    -title="My Chart"

### General flags

`-help` shows a short overview of the command line flags.  
The program ends without any further processing.

With the flag `-about` information about version, license of 
**bars** are displayed.  
The program ends without any further processing.

The flag `-manual` shows this documentation.  
The program ends without any further processing.

`-verbose` prints information about the parsing of the input
data to stdout. This is helpful for identifying parsing problems.


## Parsing of input data

For each line of the input data **bars** tries to split into a value
(number) part, and a label (text). Everything which is not part of 
the value is the label.  
Lines which don't match this pattern are ignored, as well as comment
line. Comments start by default with a "#". This can be changed with 
the flag `-comment`. For example, it can be defined that lines
beginning with a semicolon are comments using:

    --comment=";"

### Numeric value position 

The numbers for the chart are expected to be at the beginning of a
line. If the numbers are at the end use `-at-end`.

### Separating numeric value and label 

**bars** splits each line of the input data into two parts. One is 
the *value* the other is the *label* for the value in the chart.

To do this *bars* checks for characters which are part of number
from one side of the input line. As soon as non-numeric character is
found the rest of the line is label.

The default number characters are "0123456789+-.,E". You can define
additional number characters with the flag `-add-num-chars=`.
This is useful if there is a unit next to the number, which we won't
like to have as a part of the label.
The Additional number characters are removed before parsing the
number.

Example:

    $ echo "15° Outside temperature" | bars
    ° Outside temperature 15 ██████████████████████████████████
    $  echo "15° Outside temperature" | bars -add-num-chars=°
    Outside temperature 15 ████████████████████████████████████

Note: If you want the unit in the output it's reasonable to write it
in the header for the value with `-value-header="°C"`. 
For more information see below in the "Title, header and label texts"
section.

The non-numerical rest of the line is the label. It may still contain
characters which are unwanted as part of the label. By default, all 
white-space (space and tab) is trimmed from the label. Additional 
characters which should be removed, like a semicolon, can be added 
with the label `-trim-chars=`. 
Here is difference to the *-add-num-chars* flag, which is been made
clear in the next example, parsing the average distance of Uranus to
the sun in Astronomical Units (AU).

Without any parsing flags the unit and semicolon is part of the 
label:

    $ echo "20AU;Uranus" | bars -ascii
    AU;Uranus 20 ##############################################

Now we add the unit "AU" and the semicolon to the additional
number characters. Now the label misses the first "U" of Uranus,
because if became part of the number:

    $ echo "20AU;Uranus" | bars -ascii -add-num-chars="AU;"
    ranus 20 ##################################################

The correct way is to add the semicolon to the trim characters:

    $ echo "20AU;Uranus" | bars -ascii -add-num-chars="AU" \
      -trim-chars=";"
    Uranus 20 #################################################

### Decimal comma

By default, numbers are parsed with a dot "." as the decimal
separator. If you want to use a comma "," just use the flag
`-comma`.

## Control Output Formatting

**bars** are either displayes the input data as text or as HTML
charts.

The test version is convenient for visualization of data int the 
terminal or to send is as an email. You may have more ideas of how
to use it.

The two major HTML versions can be used as a standalone page or 
for including a chart snippet into an existing website.

To control the output the basic option is `-mode`, which is described
in the following.

### Text modes

`-mode=color` (or 'colour') is the default mode. It displays the
information in the terminal using ANSI codes for colorizing and 
formatting the output. This is the default if no mode flag is given.

Alternatively the `-mode=plain` (or `-mode=text`) displays plain
text without any colors and formatting.

Don't mix up "*ANSI color* vs *plain text*" with "*ASCII* vs *UTF-8*"
as the first is for colorizing, and the second for the used character
encoding.

### HTML modes 

You can either generate a complete HTML-page an HTML-snippet to 
include it in a webpage.  

With `-mode=page` (or `-mode=html`) a complete webpage with included
stylesheet is generated. 

To get an HTML-snippet use `-mode=snippet`. There is no CSS included
because you probably want to include it as a seperate file.  To get 
the style sheet use `-mode=css`.  In this case no input data is 
parsed. 

Remember to use the `>` to write the HTML and CSS to a file:

    bars --mode=page test.dat > test.html

For both ways - text and HTML - are several options to control the
layout.

### Number format

By default, the output of a value has no decimals. If you would like
to have decimals use the option, ``--decimals`` with the number of 
decimals you like to see.

Example:

    bars --decimals=2

### Color and plain text output (only valid for text modes)

By default, the output mode is set to "color" which gives you a nice
colored output with some text formatting. 

If you want to avoid ANSI based formatting you can set `-mode=plain`
to get a plain text output. This is useful if your terminal does not
support formatting, or you want to store charts in a text file or 
send it by email.

When switching to plain text output, optional available headers can
not use the underline formatting. In this case an extra ruler line 
between header and data is shown.  
This can be prevented with the `-no-ruler` flag.

### Output width (only valid for text modes)

By default, the output width is set to 80 characters.
With flag `-output-width` this can be changed.

The calculation of the width of the output elements is roughly done
as follows: The maximum width of the text labels and the values, plus
the spacer in between are added up. The rest of the output width is
used for the bar charts. If there is not enough space for a 
reasonable chart (minimum 7 characters) the label text is shortened.
This is displayed with an ellipses "…".

There is a lower limit around 25 characters, depending on the values,
for the output width.  

### UTF8 or ASCII  (only valid for text modes)

By default, bars uses UTF8 symbols in it's output. I.e., for the bars 
or to shorten text with an ellipsis "…". In case your output device 
can't handle this, the parameter `-ascii` will fall back to pure 
ASCII symbols. (This does not apply to input data or options using
UTF8 characters.)

### Self defined symbols (only valid for text modes)

in case there are negative numbers in the data a zero axis is 
displayed. By default, this is printed using the pipe "|" symbol
or a vertical line.  
With the flag `-zero-symbol` this symbol can be changed. 

The default bar chart symbol "#" can easily be replaced with the
`-chart-symbol` flag. Of course, it's possible to use any UTF8 symbol
like "🮱".  

Hint: The output of extra wide UTF8 symbols may look weird, 
depending on your font and terminal program. 

### Summary lines (all modes)

Below the data charts, summary lines can be displayed: 

* The option `-sum` shows the sum of values.
* `-count` displays the number of values in the chart.
* With `-average`, the average of the values are show.

The lines show up in the above order.
Before each value an appropriate label is printed, depending on the 
space available for the labels above:

* "Sum", "Σ"
* "Count", "Cnt", "#"
* "Average", "Avg", "⦵"

If a corresponding label oder -label or -text flag is set, the above
flags are set implicitly.

### Title, header and label texts

There are several options to define a title, headings or custom texts
for the sum, count and average summary lines.

Using the `-[sum|count|average]-label=` flag override the default 
text elements for the summary lines.

The following examples shows the flags at there corresponding
position in the chart:

    -title
    -label-header      -value-header -chart-header
    ───────────────── ────────────── ───────────────────────────────────────────────
    Label A from File 11111111111111 ███████████████
    Label B from File 22222222222222 ███████████████████████████████
    Label C from File 33333333333333 ███████████████████████████████████████████████
    ───────────────── ────────────── ───────────────────────────────────────────────
    -sum-label        66666666666666 -sum-text
    -count-label                   3 -count-text
    -average-label    22222222222222 -average-text

The command used to display this is:

    bars -mode=plain \
         -title=-title \
         -label-header=-label-header \
         -value-header=-value-header \
         -chart-header=-chart-header \
         -sum-label=-sum-label \
         -sum-text=-sum-text \
         -count-label=-count-label \
         -count-text=-count-text \
         -average-label=-average-label \
         -average-text=-average-text \
         mini.txt

Remember that the example uses plain UTF-8 text. In color and HTML
mode all text elements are reasonably formatted.

## About
