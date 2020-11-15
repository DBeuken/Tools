# Tool: aliveDB
This tool can be used to see if sites are up, down or unreachable.
The strength of this tool, is that is uses threading.


## Prerequirements
Linux OS
Golang installed


## Set this tool up
To setup this tool, you first have to install golang. 
If you have a Debian OS, you can simply install golang by the following command:
  sudo apt-get install golang-go
  
Otherwise, you can install golang on www.golang.org

To check if golang is installed properly, type in the following command:
  go version



Now, you can make this code an executable by the following command:
  go build aliveDB.go


## Helpmenu

 __  __  ______  __      ______     __    __  ______  __   __  __  __
/\ \_\ \/\  ___\/\ \    /\  == \   /\ "-./  \/\  ___\/\ "-.\ \/\ \/\ \
\ \  __ \ \  __\\ \ \___\ \  _-/   \ \ \-./\ \ \  __\\ \ \-.  \ \ \_\ \
 \ \_\ \_\ \_____\ \_____\ \_\      \ \_\ \ \_\ \_____\ \_\\"\_\ \_____\
  \/_/\/_/\/_____/\/_____/\/_/       \/_/  \/_/\/_____/\/_/ \/_/\/_____/

This tool can be used to check if websites are alive. 
To do that, this tool uses the "fping"-command.
The strength of this tool however, is that threading is supported.

To use this tool, you have to specify a file that contains sites. 

TAGS:
The underlying tags can be used.
	-th	: Amount of threads. 
		  Default = 4
	-t	: Time in seconds to wait after every request. This slows down the amount of requests
		  per second.
		  Default = 0 seconds
	-f	: File to use
		  The file should contain 1 site per line so that this tool can handle each site one
		  by one.


EXAMPLE SYNTAX:
	aliveDB -th 8 -f sites.txt -t 2	  //Threads: 8; File: sites.txt; Intermediate time: 2 seconds
	aliveDB -f sites.txt		  //Threads: 4; File: sites.txt; Intermediate time: 0 seconds
	aliveDB -f domains.txt -t 1	  //Threads: 4; File: domains.txt; Intermediate time: 1 second
