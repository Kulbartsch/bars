# bars

**A CLI tool to generate bar charts in the terminal (or as an HTML snippet [WIP]).**

**bars** follows the Unix philosophy:
> Write programs that do one thing and do it well. 
> Write programs to work together. 
> Write programs to handle text streams, because that is a universal interface.

Features: 

* [X] Handles UTF8 characters
* [X] Values can be on the left or right side
* [X] Accept units as part of the value like ``12.90$``
* [X] Alternatively accept a comma as a decimal separator
* [X] Define the number of decimals to display 

I just started this tool. The basic version is working, 
but there will be more features:

* [ ] HTML snippet output to integrate in a website
* [ ] Colored terminal output
* [ ] and more ...

So stay tuned!

## Example 

Displaying the disk usage as bars:

    $  du -b *
    60      bill.txt
    237     performance.txt
    115     temperatur.txt

Each input line is split up into an _value_ and _label_ part. 
By default, the output is _label_, _value_ and _bars_:

    $  du -b * | ../bars                                                                                                                                             127 ↵
    bill.txt         60 ###############
    performance.txt 237 ############################################################
    temperatur.txt  115 #############################

There are parameters to accept units in the values, 
use a comma instead of a dot as a decimal sepeaator 
and more. 

**See here for [more examples](examples/example.md).** 