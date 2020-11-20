package main

import (
	"fmt"
	"strconv"
	"flag"
	"os"
	"os/exec"
	"strings"
	"net/http"
	"sync"
	"time"
)

var banner = `______________________________________________________________________________________________________
                 __          __       __    _____                 ____    ____      __ _           _
 __     __ __   /\ \        /\ \  __                             /\  _ \ /\  _ \           __ __   
 _  ____  __  __\ \ \____   \_\ \/\_\     __      __      __   __\ \ \/\ \ \ \_\ \                 
    ___\ \ \  __\ \ \____   \_\ \/\_\     __      __      __   _ __\ \ \/\ \ \ \_\ \              _ 
   /  __\/\ \/\ \\ \  __ \  / _  \/\ \  / _  \  / _  \  / __ \/\  __\ \ \ \ \ \  _ <     x  x      
  /\__   \ \ \_\ \\ \ \_\ \/\ \_\ \ \ \/\ \_\ \/\ \_\ \/\  __/\ \ \/ \ \ \_\ \ \ \_\ \  \____/   _  
  \/\____/\ \____/ \ \_ __/\ \___ _\ \_\ \____ \ \____ \ \____\\ \_\  \ \____/\ \____/             
   \/___/  \/___/   \/___/  \/__ _ /\/_/\/____\ \/____\ \/____/ \/_/  \\/___/  \/___/ __ _         
   \______\   ___ \\     \  \  \    \  \   /\____/ /\____/\    \ \ \    \ \__/ \___/     _     __  
  \/\____/\ \____/ \ \_ __/\ \___ _\ \_\ \____ \ \____ \ \____\\ \_\  \ \____/\ \/                 _
      ____              __ _              /\____/ /\____/     __  _ __              __     __ _ _   
   __        _                      __    \_/__/  \_/__/   ___         ___          __ __          
______________________________________________________________________________________________________
`

var helpmenu = ` _    _      _
| |  | |    | |
| |__| | ___| |_ __  _ ______   ___ _ __  _   _
|  __  |/ _ \ | '_ \| '_   _ \ / _ \  _ \| | | |
| |  | |  __/ | |_) | | | | | |  __/ | | | |_| |
|_|  |_|\___|_|  __/|_| |_| |_|\___|_| |_|\__,_|
              | |
              |_|

This tool can be used to search for subdomains. 
To do that, you can use your own wordlist.

TAGS-----------------------------------------
-t	: Target. This is the site you aim at.
-f	: File. The file you want to use to discover subdomains.
-th	: Threads. Amount of workers.
	  default = 4
-d	: Delay. The amount of time in seconds to wait before making the next request for 
	  every thread.
	  default = 0
-s	: HTTPS. Use https instead of http.
	  default = http


EXAMPLE SYNTAX-------------------------------
subdiggerDB -f subdoms.txt -t youtube.com	  	  //file: subdoms.txt; target: youtube.com;
							    threads: 4; delay: 0 seconds; http
subdiggerDB -f subdoms.txt -t youtube.com -th 10 -d 2 -s  //file: subdoms.txt; target: youtube.com; 
							    threads: 10; Delay: 2 seconds; https

`
var wg sync.WaitGroup

func worker(sites []string, delay int) {
	for i := 0; i < len(sites); i++ {
		response, err := http.Get(sites[i])
		if (err != nil) {
			fmt.Println("error			" + sites[i])
		} else if (len(http.StatusText(response.StatusCode)) < 5) {
			fmt.Println(strconv.Itoa(response.StatusCode) + "	" + http.StatusText(response.StatusCode) + "		" + sites[i])
		} else {
			fmt.Println(strconv.Itoa(response.StatusCode) + "	" + http.StatusText(response.StatusCode) + "	" + sites[i])
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
	wg.Done()
}

func start(threads int, sites []string, delay int) {
	for t := 0; t < threads; t++ {
		var site_filter []string
		for site_loop := t; site_loop < len(sites); site_loop = site_loop + threads {
			site_filter = append(site_filter, sites[site_loop])
		}
		wg.Add(1)
		go worker(site_filter, delay)
	}
	wg.Wait()
}

func make_sites(file_output string, target string, first_part string) []string {
	file_split := strings.Split(file_output, "\n")
	var sites []string

	for i := 0; i < len(file_split) - 1; i++ {
		sites = append(sites, first_part + file_split[i] + "." + target)
	}
	return sites
}

func file_opener(file string) string {
	commandinput, err := exec.Command("bash", "-c", file).Output()
	if (err != nil) {
		fmt.Println("ERROR: File doesn't exist/cannot be used\n")
		fmt.Println(helpmenu)
		os.Exit(3)
		return("This won't be used")
	} else {
		file_str := string(commandinput[:])
		return file_str
	}
}

func main() {
	var target string
	var file string
	var threads_str string
	var delay_str string
	flag.StringVar(&target, "t", "false", "0")
	flag.StringVar(&file, "f", "false", "0")
	flag.StringVar(&threads_str, "th", "false", "0")
	flag.StringVar(&delay_str, "d", "false", "0")
	https := flag.Bool("s", false, "-")
	help_tag := flag.Bool("help", false, "-")
	flag.Parse()

	if (strconv.FormatBool(*help_tag) == "true") {
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	if (target == "false") {
		fmt.Println("Site not selected\n")
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	if (file == "false") {
		fmt.Println("File not selected\n")
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	if (threads_str == "false") {
		threads_str = "4"
	}

	threads, err := strconv.Atoi(threads_str)
	if (err != nil) {
		fmt.Println("You should give the threads-tag a number\n")
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	if (delay_str == "false") {
		delay_str = "0"
	}

	delay, err := strconv.Atoi(delay_str)
	if (err != nil) {
		fmt.Println("You should give the delay-tag a number\n")
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	var first_part string
	if (strconv.FormatBool(*https) == "true") {
		first_part = "https://"
	} else {
		first_part = "http://"
	}

	file_output := file_opener("cat " + file)
	sites := make_sites(file_output, target, first_part)

	//THE START
	fmt.Println(banner)
	fmt.Println("Site:		" + target)
	fmt.Println("File: 		" + file)
	fmt.Println("Threads: 	" + threads_str)
	fmt.Println("Delay: 		" + delay_str + " seconds")
	fmt.Println("______________________________________________________________________________________________________\n")
	fmt.Println("Starting...\n")
	fmt.Println("Status:	Response:	Site:")

	start(threads, sites, delay)

}
