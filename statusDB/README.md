# Tool: statusDB
This tool can be used to get the status codes of websites.\
The strength of this tool, is that is uses threading and that you can use your own wordlist.


## Prerequirements
Linux OS\
Golang installed


## Set this tool up
To setup this tool, you first have to install golang. \
If you have a Debian OS, you can simply install golang by the following command:\
  sudo apt-get install golang-go\
  \
Otherwise, you can install golang on www.golang.org
\
To check if golang is installed properly, type in the following command:\
  go version\
\
Now, you can make this code an executable by the following command:\
  go build statusDB.go

## Helpmenu
_________________________
 
       HELPMENU
_________________________

This tool can be used to get the response codes of websites.\
To do that, you can use your own wordlist.

TAGS--------------------------------------\
 -th	: Amount of Threads.\
 	  Default = 4\
 -f	: File\
 -d	: Delay in seconds\
 	  Default = 0

EXAMPLE SYNTAX----------------------------\
statusDB -f sites.txt			//File: sites.txt; Threads: 4; Delay: 4\
statusDB -f sites.txt -th 10		//File: sites.txt; Threads: 10; Delay: 0\
statusDB -f sites.txt -th 2 -d 5	//File: sites.txt; Threads: 2; Delay: 5
