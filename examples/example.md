# bars Examples

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

    $ bars -comma -addNumChars=€ -comment="[" -decimals=2 examples/bill.txt

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

    bars -decimals 2 --comma -addNumChars=°C --atEnd --zero="|" examples/temperatur.txt 

You can use one or two dashes for a parameter. 

The Result:

    Berlin      6.10                         |##############################
    Dresden    -4.80 ########################|
    Düsseldorf  7.50                         |#####################################
    Hamburg     6.00                         |##############################
    München     0.00                         |
    Stuttgart  -0.50                       ##|



## Visualize your disk usage 

    df -x squashfs -x tmpfs | awk '{print $5 " " $6} END{print "100% 100%"}' | bars -addNumChars=% 

This command may result in something like: 

    /                     35 ###############
    /boot/efi              1
    /run/media/me/media   63 ###########################
    /run/media/me/backup  72 ##############################
    /run/media/me/data    13 #####
    100%                 100 ###########################################

I added a 100% bar with awk for better comparability.

