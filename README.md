# bars

_Visualize numbers with ease! In the terminal or for web pages._

**A CLI tool to generate bar charts as (colored) text or HTML, written in Go**

**bars** follows the Unix philosophy:
> Write programs that do one thing and do it well.  
> Write programs to work together.  
> Write programs to handle text streams, because that is a universal interface.

## Main Features 

* [X] Handles UTF8 characters
* [X] Values can be on the left or right side
* [X] Accept units as part of the value like ``12.90$``
* [X] Alternatively accept a comma as a decimal separator
* [X] Define the number of decimals to display 
* [X] Colored terminal output
* [X] Optional print the sum, count and average of the give values
* [X] HTML output to integrate in a website or as a standalone page
* [X] No runtime dependencies, just one binary!
* [X] No compile time dependencies, just one repo!
* [X] Manual included 

... and more to come.\
Checkout the detailed [features](features.md) page.\
So stay tuned!

## Example 

Displaying the disk usage as bars:

    $  du -b *
    60      bill.txt
    237     performance.txt
    115     temperatur.txt

Each input line is split up into an _value_ and _label_ part. 
By default, the output is _label_, _value_ and _bars_:

    $  du -b * | ../bars --ascii
    bill.txt         60 ###############
    performance.txt 237 ############################################################
    temperatur.txt  115 #############################

There are parameters to accept units in the values, 
use a comma instead of a dot as a decimal separator 
and more. 

**See here for**
* [more examples](examples/example.md)
* [an exhaustive manual](manual.md)

## License

Copyright © 2021 Alexander Kulbartsch   
License AGPL-3.0-or-later (GNU Affero General Public License 3 or later)

You find a copy of the [GNU Affero General Public License](LICENSE) in this repo.  
If not, see <http://www.gnu.org/licenses/>.