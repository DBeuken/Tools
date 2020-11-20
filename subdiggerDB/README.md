# Tool: subdiggerDB
This tool can be used to brute-force for possible subdomains.\
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
  go build subdiggerDB.go

## Helpmenu

This tool can be used to search for subdomains. \
To do that, you can use your own wordlist.\

TAGS-----------------------------------------\
-t	: Target. This is the site you aim at.\
-f	: File. The file you want to use to discover subdomains.\
-th	: Threads. Amount of workers.\
	  default = 4\
-d	: Delay. The amount of time in seconds to wait before making the next request for \
	  every thread.\
	  default = 0\
-s	: HTTPS. Use https instead of http.\
	  default = http


EXAMPLE SYNTAX-------------------------------\
subdiggerDB -f subdoms.txt -t youtube.com	  	  //file: subdoms.txt; target: youtube.com;\
							    threads: 4; delay: 0 seconds; http\
subdiggerDB -f subdoms.txt -t youtube.com -th 10 -d 2 -s  //file: subdoms.txt; target: youtube.com; \
							    threads: 10; Delay: 2 seconds; https
