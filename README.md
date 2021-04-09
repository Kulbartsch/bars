# bars

A CLI tool to generate bar charts in the terminal or as an HTML snippet

Just started. Work in progress. Stay tuned! 

Basic version is working

Example, display disk usage as bars:

    $  du -b *
    60      bill.txt
    237     performance.txt
    115     temperatur.txt
    
    $  du -b * | ../bars                                                                                                                                             127 ↵
    bill.txt         60 ###############
    performance.txt 237 ############################################################
    temperatur.txt  115 #############################
