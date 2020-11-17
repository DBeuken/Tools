# Tool: grabDB
This tool can be used to grab multiple strings of a file.\
This tool uses Linux's "grep"-command.\
The strength of this tool, is that is can put a wordlist in the grep's E-tag.

## Prerequirements
Linux OS

## Set this tool up
To setup this tool, all you have to do is enter the following command:\
  chmod +x grabDB.py


## Helpmenu
This is the grabDB helpmenu.\
\
  With this tool, you can grep multiple strings of a file simultaniously.\
  The difference with the -E tag of grep, is that you can simply filter based on the strings of a wordlist-file.
  
  
  TAGS------------------------------------------------------------------------------------------------\
  --help		:	Display this helpmenu\
  -v		:	Search for everything, except the inserted words\
  [Other tags of the grep-command]
  
  
  
  SYNTAX----------------------------------------------------------------------------------------------\
  cat file.txt | grabDB wordlist.txt\
  cat file.txt | grabDB wordlist.txt -v\
  cat file.txt | grabDB wordlist.txt [Other grep-tags] -v
  
