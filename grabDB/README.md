# Tool: aliveDB
This tool can be used to see if sites are up, down or unreachable.
The strength of this tool, is that is uses threading.


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
  go build aliveDB.go

## Helpmenu
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
