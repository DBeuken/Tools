# Tool: listDB
This tool can be used to make lists that only contain numbers.\
You can also add brackets, commas, etc.

## Prerequirements
Linux OS\
Python3 installed

## Set this tool up
To setup this tool, you first have to install python3
If you have a Debian OS, you can install python3 by the following command:\
sudo apt-get install python3\

Now, you have to convert this code to an executable:\
chmod +x listDB.py

## Helpmenu
This tool is being used to make a range of numbers.\
You can make your own list with the tags below.\
    
    h : horizontal list\
    v : vertical list\
    s : list with skipping\
    c : list with commas\
    p : list with parentheses\
    b : list with brackets\

    -v : range without skipping\
        usage: listDB -h <first number>,<last number>\
        usage example: listDB -e 5,7\
            //OUTPUT:\
            5\
            6\
            7\
    -vs : range with skipping\
        usage: listDB -hs <first number>,<last number>,<step size>\
        usage example: listDB -vs 5,14,3\
            //OUTPUT:\
            5\
            8\
            12\
    -h : range without skipping\
        usage: listDB -h <first number>,<last number>\
        usage example: listDB -h 5,10\
            //OUTPUT:\
            5 6 7 8 9 10\
    -hs : range with skipping\
        usage: listDB -hs <first number>,<last number>,<step size>\
        usage example: listDB -hs 5,15,4\
            //OUTPUT:\
            5 9 13\
    -ch : range without skipping and with commas\
        usage: listDB -ch <first number>,<last number>\
        usage example: listDB -ch 3,5\
            //OUTPUT:\
            3,4,5\
    -hp : range with parentheses\
        usage: listDB -ph <first number>,<last number>\
        usage example: listDB -ph 3,6\
            //OUTPUT:\
            (3 4 5 6)\
    -hb : range with brackets\
        usage: listDB -hb <first number>,<last number>\
        usage example: listDB -hb 3,6\
            //OUTPUT:\
            [3 4 5 6]\
    -chs : range with skipping and with commas\
        usage: listDB -chs <first number>,<last number>,<step size>\
        usage example: listDB -chs 3,9,2\
            //OUTPUT:\
            3,5,7,9\
    -pch : range without skipping and with commas and parentheses\
        usage: listDB -pch <first number>,<last number>\
        usage example: listDB -pch 3,5\
            //OUTPUT:\
            (3,4,5)\
    -pchs : range with skipping and with commas and parentheses\
        usage: listDB -pchs <first number>,<last number>,<step size>\
        usage example: listDB -pchs 3,9,2\
            //OUTPUT:\
            (3,5,7,9)\
    -bch : range without skipping and with commas and brackets\
        usage: listDB -bch <first number>,<last number>\
        usage example: listDB -bch 3,5\
            //OUTPUT:\
            [3,4,5]\
    -bchs : range with skipping and with commas and brackets\
        usage: listDB -bchs <first number>,<last number>,<step size>\
        usage example: listDB -bchs 3,9,2\
            //OUTPUT:\
            [3,5,7,9]
