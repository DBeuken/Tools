#!/usr/bin/python3

import sys

help_prescription = r'''
--------------------------------------------------------------
        ___                 __    ____    ____      
       /\_ \    __         /\ \__/\  _`\ /\  _`\    
       \//\ \  /\_\    ____\ \ ,_\ \ \/\ \ \ \L\ \  
         \ \ \ \/\ \  /',__\\ \ \/\ \ \ \ \ \  _ <' 
          \_\ \_\ \ \/\__, `\\ \ \_\ \ \_\ \ \ \L\ \
          /\____\\ \_\/\____/ \ \__\\ \____/\ \____/
          \/____/ \/_/\/___/   \/__/ \/___/  \/___/ 

---------------------------------------------------------------                                                                                          
PRESCRIPTION:
This tool is being used to make a range of numbers.
You can make your own list with the tags below.
    
    h : horizontal list
    v : vertical list
    s : list with skipping
    c : list with commas
    p : list with parentheses
    b : list with brackets

    -v : range without skipping
        usage: listDB -h <first number>,<last number>
        usage example: listDB -e 5,7
            //OUTPUT:
            5
            6
            7
    -vs : range with skipping
        usage: listDB -hs <first number>,<last number>,<step size>
        usage example: listDB -vs 5,14,3
            //OUTPUT:
            5
            8
            12
    -h : range without skipping
        usage: listDB -h <first number>,<last number>
        usage example: listDB -h 5,10
            //OUTPUT:
            5 6 7 8 9 10
    -hs : range with skipping
        usage: listDB -hs <first number>,<last number>,<step size>
        usage example: listDB -hs 5,15,4
            //OUTPUT:
            5 9 13
    -ch : range without skipping and with commas
        usage: listDB -ch <first number>,<last number>
        usage example: listDB -ch 3,5
            //OUTPUT:
            3,4,5
    -hp : range with parentheses
        usage: listDB -ph <first number>,<last number>
        usage example: listDB -ph 3,6
            //OUTPUT:
            (3 4 5 6)
    -hb : range with brackets
        usage: listDB -hb <first number>,<last number>
        usage example: listDB -hb 3,6
            //OUTPUT:
            [3 4 5 6]
    -chs : range with skipping and with commas
        usage: listDB -chs <first number>,<last number>,<step size>
        usage example: listDB -chs 3,9,2
            //OUTPUT:
            3,5,7,9
    -pch : range without skipping and with commas and parentheses
        usage: listDB -pch <first number>,<last number>
        usage example: listDB -pch 3,5
            //OUTPUT:
            (3,4,5)
    -pchs : range with skipping and with commas and parentheses
        usage: listDB -pchs <first number>,<last number>,<step size>
        usage example: listDB -pchs 3,9,2
            //OUTPUT:
            (3,5,7,9)
    -bch : range without skipping and with commas and brackets
        usage: listDB -bch <first number>,<last number>
        usage example: listDB -bch 3,5
            //OUTPUT:
            [3,4,5]
    -bchs : range with skipping and with commas and brackets
        usage: listDB -bchs <first number>,<last number>,<step size>
        usage example: listDB -bchs 3,9,2
            //OUTPUT:
            [3,5,7,9]
'''



try:
    word1 = sys.argv[1]
except:
    print(help_prescription)
    sys.exit()

try: 
    word2 = sys.argv[2]
except:
    print(help_prescription)
    sys.exit()

if word1 == "--help" or word1 == "help" or word1 == "--h":
    print(help_prescription)
    sys.exit()



word1_lijst = []
try:
    word2_lijst = word2.split(",")
except:
    print(help_prescription)
    sys.exit()

for letter in range(0, len(word1)):
    word1_lijst.append(word1[letter])




#VERTICAL LIST
if "v" in word1_lijst:
    if len(word1_lijst) == 2:
        for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
            print(number)
    if "s" in word1_lijst:
        for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1, int(word2_lijst[2])):
            print(number)




#HORIZONTAL LIST
if "h" in word1_lijst:
    if len(word1_lijst) == 2:
        for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
            print(str(number) + " ", end='')
        print(" ")

    if len(word1_lijst) == 3:
        if "c" in word1_lijst:
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
                if number != int(word2_lijst[1]):
                    print(str(number) + "," , end='')
                else:
                    print(str(number) , end='')
            print(" ")

        if "s" in word1_lijst:
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1, int(word2_lijst[2])):
                print(str(number) + " " , end='')
            print(" ")

        if "p" in word1_lijst:
            print("(" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
                if number != int(word2_lijst[1]):
                    print(str(number) + " " , end='')
                else:
                    print(str(number) , end='')
            print(")")

        if "b" in word1_lijst:
            print("[" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
                if number != int(word2_lijst[1]):
                    print(str(number) + " " , end='')
                else:
                    print(str(number) , end='')
            print("]")


    if len(word1_lijst) == 4:
        if "s" in word1_lijst and "c" in word1_lijst:
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]), int(word2_lijst[2])):
                if number != int(word2_lijst[1]) -1 and number + int(word2_lijst[2]) < int(word2_lijst[1]):
                    print(str(number) + "," , end='')
                else:
                    print(str(number) , end='')
            print(" ")

        if "p" in word1_lijst and "c" in word1_lijst:
            print("(" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
                if number != int(word2_lijst[1]):
                    print(str(number) + "," , end='')
                else:
                    print(str(number) , end='')
            print(")")

        if "b" in word1_lijst and "c" in word1_lijst:
            print("[" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1):
                if number != int(word2_lijst[1]):
                    print(str(number) + "," , end='')
                else:
                    print(str(number) , end='')
            print("]")

        if "b" in word1_lijst and "s" in word1_lijst:
            print("[" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1 , int(word2_lijst[2])):
                if number != int(word2_lijst[1]) -1 and number + int(word2_lijst[2]) < int(word2_lijst[1]):
                    print(str(number) + " " , end='')
                else:
                    print(str(number) , end='')
            print("]")

        if "p" in word1_lijst and "s" in word1_lijst:
            print("(" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]) + 1 , int(word2_lijst[2])):
                if number != int(word2_lijst[1]) -1 and number + int(word2_lijst[2]) < int(word2_lijst[1]):
                    print(str(number) + " " , end='')
                else:
                    print(str(number) , end='')
            print(")")

    
    if len(word1_lijst) == 5:
        if "p" in word1_lijst:
            print("(" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]), int(word2_lijst[2])):
                if number != int(word2_lijst[1]) -1 and number + int(word2_lijst[2]) < int(word2_lijst[1]):
                    print(str(number) + "," , end='')
                else:
                    print(str(number) , end='')
            print(")")

        if "b" in word1_lijst:
            print("[" , end='')
            for number in range(int(word2_lijst[0]), int(word2_lijst[1]), int(word2_lijst[2])):
                if number != int(word2_lijst[1]) -1 and number + int(word2_lijst[2]) < int(word2_lijst[1]):
                    print(str(number) + "," , end='')
                else:
                    print(str(number) , end='')
            print("]")
