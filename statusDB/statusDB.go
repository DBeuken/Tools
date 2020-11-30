package main

import (
	"fmt"
	"strconv"
	"flag"
	"os"
	"io/ioutil"
	"strings"
	"net/http"
	"time"
	"sync"
)

var banner = `       __             __                     ____    ____      
      /\ \__         /\ \__                 /\  _ \ /\  _ \    
  ____\ \  _\   ___  \ \  _\  __  __    ____\ \ \/\ \ \ \_\ \  
 /  __\\ \ \/  / __ \ \ \ \/ /\ \/\ \  /  __\\ \ \ \ \ \  _ < 
/\__   \\ \ \_/\ \_\ \_\ \ \_\ \ \_\ \/\__   \\ \ \_\ \ \ \_\ \
\/\____/ \ \__\ \__/ \_\\ \__\\ \____/\/\____/ \ \____/\ \____/
 \/___/   \/__/\/__/\/_/ \/__/ \/___/  \/___/   \/___/  \/___/ 
                                                               
`

var helpmenu = `_________________________
 
       HELPMENU
_________________________

This tool can be used to get the response codes of websites.
To do that, you can use your own wordlist.

TAGS--------------------------------------
 -th	: Amount of Threads.
 	  Default = 4
 -f	: File
 -d	: Delay in seconds
 	  Default = 0

EXAMPLE SYNTAX----------------------------
statusDB -f sites.txt			//File: sites.txt; Threads: 4; Delay: 4
statusDB -f sites.txt -th 10		//File: sites.txt; Threads: 10; Delay: 0
statusDB -f sites.txt -th 2 -d 5	//File: sites.txt; Threads: 2; Delay: 5
`

var wg sync.WaitGroup

func get_status(file_split []string, delay int) {
	for i := 0; i < len(file_split); i++ {
		response, err := http.Get(file_split[i])
		if (err != nil) {
			fmt.Println("error			" + file_split[i])
		} else if (len(http.StatusText(response.StatusCode)) < 5) {
			fmt.Println(strconv.Itoa(response.StatusCode) + "	" + http.StatusText(response.StatusCode) + "		" + file_split[i])
		} else {
			fmt.Println(strconv.Itoa(response.StatusCode) + "	" + http.StatusText(response.StatusCode) + "	" + file_split[i])
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
	wg.Done()
}

func start(file_split []string, delay int, threads int) {
	for t := 0; t < threads; t++ {
		var sites []string
		for site_loop := t; site_loop < len(file_split); site_loop = site_loop + threads {
			sites = append(sites, file_split[site_loop])
		}
		wg.Add(1)
		go get_status(sites, delay)
	}
	wg.Wait()
}

func fileread(file string) string {
	filetext, err := ioutil.ReadFile(file)
	if (err != nil) {
		fmt.Println("ERROR: file doesn't exist\n")
		fmt.Println(helpmenu)
		os.Exit(3)
		return "-"
	} else {
		return string(filetext)
	}
}

func main() {
	var threads_str string
	var file string
	var delay_str string
	flag.StringVar(&threads_str, "th", "Threads not selected", "-")
	flag.StringVar(&file, "f", "File not selected", "-")
	flag.StringVar(&delay_str, "d", "Delay not selected", "-")
	help_tag := flag.Bool("help", false, "-")
	flag.Parse()

	if (strconv.FormatBool(*help_tag) == "true") {
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	filetext := fileread(file)

	threads, err := strconv.Atoi(threads_str)
	if (err != nil) {
		threads = 4
	}

	delay, err := strconv.Atoi(delay_str)
	if (err != nil) {
		delay = 0
	}

	file_split := strings.Split(filetext, "\n")

	fmt.Println(banner)
	fmt.Println("       v1.0.0")
	fmt.Println("___________________________________________________________________________\n")
	fmt.Println(" :: Threads:	" + strconv.Itoa(threads))
	fmt.Println(" :: File:	" + file)
	fmt.Println(" :: Delay:	" + strconv.Itoa(delay))
	fmt.Println("___________________________________________________________________________\n")
	start(file_split, delay, threads)

}
