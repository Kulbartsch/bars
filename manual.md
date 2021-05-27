# bars - Manual

*Easily generate bar charts from the command line as text or HTML*

## General Information

**bars** follows the Unix philosophy:
> Write programs that do one thing and do it well.
> Write programs to work together.
> Write programs to handle text streams, because that is a universal interface.

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

`-v` or `-verbose` prints information about the parsing of the input
data to stdout. This is helpful for identifying parsing problems.


## Flags for parsing of input data

	myParam.comment	 	= flag.String("comment", "#", "comment line start")

### Where is the numeric value 

	myParam.valueAtEnd	= flag.Bool("at-end", false, "values are at the end of a line")

### What is the numeric value 

	myParam.addNumChars	= flag.String("add-num-chars", "", "additional characters representing a number")

### Decimal comma

	myParam.comma		= flag.Bool("comma", false, "use comma as decimal separator")

## Control Output Formatting

	myParam.mode 		= flag.String("mode", "color", "display mode, one of 'plain', 'color', 'snippet', 'css', 'page'")


## Flags for controlling output 

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


	myParam.zero		= flag.String("zero", "", "symbol to represent the 0 line in text chart")
	myParam.zero		= flag.String("zero-symbol", "", "symbol to represent the 0 line in text chart")
	myParam.chartSymbol = flag.String("chart-symbol", "", "alternative symbol for text-mode bars")	

### HTML Output (only valid for html modes)

	myParam.title 		= flag.String("title", "", "Title of the chart")

### Header and summary (all modes)

	myParam.labelHeader = flag.String("label-header", "", "header text for the label")
	myParam.valueHeader = flag.String("value-header", "", "header text for the value")
	myParam.chartHeader	= flag.String("chart-header", "", "header text for the chart")

	myParam.sum 		= flag.Bool("sum", false, "display sum of values")
	myParam.count 		= flag.Bool("count", false, "display count of values")
	myParam.average		= flag.Bool("average", false, "display average of values")


