package main

import (
	"fmt"
	"os/exec"
	"flag"
	"strings"
	"os"
	"strconv"
	"time"
	"sync"
)

var logo = `
------------------------------------------------------------------------------------------------

          __          ___             ______ _ ____
              ___                        ____    ____          
             /\_ \    __                /\  _ \ /\  _ \               ___                
             /\_ \    __          ______  /\  _ \ /\  _\    
         __  \//\ \  /\_\  __  __     __\ \ \/\ \ \ \_\ \       x  x
       /'__ \  \ \ \ \/\ \/\ \/\ \  /'__ \ \ \ \ \ \  _ <      \____/
      /\ \_\.\_ \_\ \_\ \ \ \ \_/ |/\  __/\ \ \_\ \ \ \_\ \
      \ \__/.\_\/\____\\ \_\ \___/ \ \____\\ \____/\ \____/
       \/__/\/_/\/____/ \/_/\/__/   \/____/ \/___/  \/___/          __                      
      \ \__/.\_\/\____\\ \_\ \___/ \ \____\\ \____/\ \____/   __   __
       \/__/\/_/\/____/ \/_/\/__/   \/____/ \/___/  \/___
                         __              __

-------------------------------------------------------------------------------------------------
`

var helpmenu = `-------------------------------------------------------------------------
 __  __  ______  __      ______     __    __  ______  __   __  __  __
/\ \_\ \/\  ___\/\ \    /\  == \   /\ "-./  \/\  ___\/\ "-.\ \/\ \/\ \
\ \  __ \ \  __\\ \ \___\ \  _-/   \ \ \-./\ \ \  __\\ \ \-.  \ \ \_\ \
 \ \_\ \_\ \_____\ \_____\ \_\      \ \_\ \ \_\ \_____\ \_\\"\_\ \_____\
  \/_/\/_/\/_____/\/_____/\/_/       \/_/  \/_/\/_____/\/_/ \/_/\/_____/
-------------------------------------------------------------------------

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
`

var wg sync.WaitGroup


func worker(command string) {
	commandinput, _ := exec.Command("sh", "-c", command).Output()
	if (string(commandinput) != "") {
		fmt.Println(string(commandinput))
	}
}

func before_worker(sites_amount int, threads_int int, file_split []string, thread_loop int, time_int int) {
	for command_loop := thread_loop; command_loop < sites_amount; command_loop += threads_int {
		command := "fping " + file_split[command_loop] + " | tr -d '\n'"
		worker(command)
		time.Sleep(time.Duration(time_int) * time.Second)
	}
	wg.Done()
}

func start(threads_int int, sites_amount int, file_split []string, time_int int) {
	for thread_loop := 0; thread_loop < threads_int; thread_loop += 1 {
		wg.Add(1)
		go before_worker(sites_amount, threads_int, file_split, thread_loop, time_int)
	}
	wg.Wait()
}

func file_opener(file string) string {
	commandinput, err := exec.Command("bash", "-c", file).Output()
	if (err != nil) {
		fmt.Println("ERROR: File doesn't exist/cannot be used\n")
		fmt.Println(helpmenu)
		os.Exit(1)
		return "This won't be used"
	} else {
		file_str := string(commandinput[:])
		return file_str
	}
}

func main() {

	//MAKING FLAGS
	var file string
	flag.StringVar(&file, "f", "file not selected", "-")

	var threads string
	flag.StringVar(&threads, "th", "threads not selected", "-")

	var time_input string
	flag.StringVar(&time_input, "t", "time not selected", "-")

	help_tag := flag.Bool("help", false, "-")

	flag.Parse()

	//PRINT OUT HELPMENU
	if (strconv.FormatBool(*help_tag) == "true") {
		fmt.Println(helpmenu)
		os.Exit(1)
	}

	//DETERMINE AMOUNT OF THREADS
	threads_int, err := strconv.Atoi(threads)
	if (err != nil) {
		threads_int = 4
	}

	//PREPARING FILE
	file_output := file_opener("cat " + file)
	file_split := strings.Split(file_output, "\n")
	sites_amount := len(file_split)

	//DETERMINE TIME
	time_int, err := strconv.Atoi(time_input)
	if (err != nil) {
		time_int = 0
	}

	fmt.Println(logo)

	fmt.Println("Threads: " + strconv.Itoa(threads_int))
	fmt.Println("Intermediate time: " + strconv.Itoa(time_int))
	fmt.Println("File: " + string(file) + "\n")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("Starting...\n")

	//START SCANNING
	start(threads_int, sites_amount, file_split, time_int)
}
