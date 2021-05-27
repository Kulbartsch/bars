# bars Examples

> The output of the examples are slightly outdated due to the usage of UTF8 symbols. This plain ascii output now requires the `-ascii` parameter.
## Disk usage  

Example, display disk usage of this test files as bars:

    $  du -b *
    60      bill.txt
    237     performance.txt
    115     temperatur.txt
    
    $  du -b * | bars
    bill.txt         60 ###############
    performance.txt 237 ############################################################
    temperatur.txt  115 #############################

## A Bill

A bill: 

    [Your Restaurant Bill]
    [price    article]
     7,90€   Pizza Napoli
    19,90€   Wine, Bottle
     8,70€   Spaghetti
    14,90€   Fish of the day
     1,50€   Espresso
     2,50€   Café correto

With this command: 

    $ bars -comma -add-num-chars=€ -comment="[" -decimals=2 examples/bill.txt

The commas are used as the decimal separator.
The € sign is part of the number (we don't want it to be part of the label text).
The [ is a comment line to be ignored.
We want to keep the two decimals in the output. 

    Pizza Napoli     7.90 #######################
    Wine, Bottle    19.90 ##########################################################
    Spaghetti        8.70 #########################
    Fish of the day 14.90 ###########################################
    Espresso         1.50 ####
    Café correto     2.50 #######


## Temperatures in German cities

The File: 

    Berlin      6,1°C
    Dresden    -4,8°C
    Düsseldorf +7,5°C
    Hamburg     6,0°C
    München     0,0°C
    Stuttgart  -0,5°C

The Command: 

    bars -decimals 2 --comma -add-num-chars=°C --at-end -output-width=100 -label-header=City -value-header="°C" -chart-header="temperature in degree celsius" -no-ruler examples/temperatur.txt

You can use one or two dashes for a parameter. 

The Result:

    City______ ___°C temperature in degree celsius______________________________________________________
    Berlin      6.10                                │████████████████████████████████████████
    Dresden    -4.80 ███████████████████████████████│
    Düsseldorf  7.50                                │█████████████████████████████████████████████████
    Hamburg     6.00                                │███████████████████████████████████████
    München     0.00                                │
    Stuttgart  -0.50                             ███│



## Visualize your disk usage 

    df -x squashfs -x tmpfs | awk '{print $5 " " $6} END{print "100% 100%"}' | bars -ascii -add-num-chars=% -label-header=mount -value-header=% -chart-header=usage

This command may result in something like:

    mount                  % usage
    -------------------- --- -------------------------------------------
    /                     36 ###############
    /boot/efi              1
    /run/media/me/media   63 ###########################
    /run/media/me/backup  73 ###############################
    /run/media/me/data    13 #####
    100%                 100 ###########################################

I added a 100% bar with awk for better comparability. 
(There is a feature in the queue which makes this obsolete.)


## many flags combined

The File:

    -182.456; Methane
    -114; Ethanol
    -38.36; Mercury (quicksilver)
    *0; Water

The Command:

    bars --add-num-chars="*;" -output-width=20 -value-header="°C" -label-header=Element -chart-header=temperature -count -title="Melting Point" -ascii meltingpoint.csv

The Result: 

    Melting Point
    Element   °C temp...
    Methane -182 ######|
    Ethanol -114    ###|
    Merc...  -38      #|
    Water      0       |
    Count      4