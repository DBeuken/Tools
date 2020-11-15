#!/usr/bin/bash

if [ $1 == "--help" ]; then
	echo "  _    _      _                                  
 | |  | |    | |                                 
 | |__| | ___| |_ __  _ ______   ___ _ __  _   _ 
 |  __  |/ _ \ |  _ \|  _   _ \ / _ \  _ \| | | |
 | |  | |  __/ | |_) | | | | | |  __/ | | | |_| |
 |_|  |_|\___|_| .__/|_| |_| |_|\___|_| |_|\__,_|
               | |                               
               |_|                               

This is the grabDB helpmenu.

With this tool, you can grep multiple strings of a file simultaniously.
The difference with the -E tag of grep, is that you can simply filter based on the strings of a wordlist-file.


TAGS------------------------------------------------------------------------------------------------
--help		:	Display this helpmenu
-v		:	Search for everything, except the inserted words
[Other tags of the grep-command]



SYNTAX----------------------------------------------------------------------------------------------
cat file.txt | grabDB wordlist.txt
cat file.txt | grabDB wordlist.txt -v
cat file.txt | grabDB wordlist.txt [Other grep-tags] -v

"


elif [ "$5" != "" ]; then
	Jow=$(for word in $(cat $1) ;do echo -n $word"|" ;done | awk '{print substr($1, 1, length($0) -1)}' | tr -d '\n')
	grep $2 $3 $4 $5 -E $Jow


elif [ "$4" != "" ]; then
	Jow=$(for word in $(cat $1) ;do echo -n $word"|" ;done | awk '{print substr($1, 1, length($0) -1)}' | tr -d '\n')
	grep $2 $3 $4 -E $Jow


elif [ "$3" != "" ]; then
	Jow=$(for word in $(cat $1) ;do echo -n $word"|" ;done | awk '{print substr($1, 1, length($0) -1)}' | tr -d '\n')
	grep $2 $3 -E $Jow


elif [ "$2" != "" ]; then
	Jow=$(for word in $(cat $1) ;do echo -n $word"|" ;done | awk '{print substr($1, 1, length($0) -1)}' | tr -d '\n')
	grep $2 -E $Jow


else
	Jow=$(for word in $(cat $1) ;do echo -n $word"|" ;done | awk '{print substr($1, 1, length($0) -1)}' | tr -d '\n')
	grep -E $Jow

fi
