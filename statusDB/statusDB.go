package main

import (
	"fmt"
	"flag"
	"strconv"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"sync"
)

var banner = `         __             __                     ____    ____
        /\ \__         /\ \__                 /\  _ \ /\  _ \
    ____\ \  _\    __  \ \  _\  __  __    ____\ \ \/\ \ \ \_\ \
   /  __\\ \ \/  / __ \ \ \ \/ /\ \/\ \  /  __\\ \ \ \ \ \  _ <
  /\__   \\ \ \_/\ \_\ \_\ \ \_\ \ \_\ \/\__   \\ \ \_\ \ \ \_\ \
  \/\____/ \ \__\ \__/ \_\\ \__\\ \____/\/\____/ \ \____/\ \____/
   \/___/   \/__/\/__/\/_/ \/__/ \/___/  \/___/   \/___/  \/___/
     
         
       v.1.2.1`

var helpmenu = `_______________

   HELPMENU
_______________

This tool can be used to get the status code of a list containing sites.

TAGS______________________________________

  -f		File
  -th		Threads. Default=10
  -http		Use http-protocol
  -https	Use https-protocol
  		If http and https are not specified, http will be used
  -nobanner	Don't show the banner
  -help		Show this helpmenu


EXAMPLE SYNTAX____________________________
  statusDB -f list.txt -th 4 -http -https --nobanner
  statusDB -f list.txt -https	
`

var wg sync.WaitGroup

func worker(sites []string) {
	for i := 0; i < len(sites); i++ {
		response, err := http.Get(sites[i])
		if (err != nil) {
			fmt.Println("-	" + string(sites[i]))
		} else {
			fmt.Println(strconv.Itoa(response.StatusCode) + "	" + string(sites[i]))
		}
	}
	wg.Done()
}

func main() {
	var threads int
	flag.IntVar(&threads, "th", 10, "Amount of threads. Specified in integers")
	var file_name string
	flag.StringVar(&file_name, "f", "not selected", "File not selected")
	http := flag.Bool("http", false, "Use http")
	https := flag.Bool("https", false, "Use https")
	nobanner := flag.Bool("nobanner", false, "Don't show the banner")
	help := flag.Bool("help", false, "stdout helpmenu")

	flag.Parse()

	if (*help == true) {
		fmt.Println(helpmenu)
		os.Exit(3)
	} else if (file_name == "not selected") {
		fmt.Println("ERROR: File not selected\n")
		fmt.Println(helpmenu)
		os.Exit(3)
	}

	file_byte, err := ioutil.ReadFile(file_name)
	if (err != nil) {
		fmt.Println("ERROR: Problem with reading the file\n")
		fmt.Println(helpmenu)
		os.Exit(3)
	}
	file := string(file_byte)

	file1 := strings.Replace(file, "http://", "", -1)
	file2 := strings.Replace(file1, "https://", "", -1)
	file_split := strings.Split(file2, "\n")

	var sites []string

	if (*http == true || *http == false && *https == false) {
		for number := 0; number < len(file_split) -1; number++ {
			sites = append(sites, "http://" + string(file_split[number]))
		}
	}
	if (*https == true) {
		for number := 0; number < len(file_split) -1; number++ {
			sites = append(sites, "https://" + string(file_split[number]))
		}
	}

	if (*nobanner == false) {
		fmt.Println(banner)
		fmt.Println("_________________________________\n")
		fmt.Println(" :: Theads:	" + strconv.Itoa(threads))
		fmt.Println(" :: File:	" + file_name)
		if (*https == true && *http == true) {
			fmt.Println(" :: Protocol:	https, http")
		} else if (*https == true && *http == false) {
			fmt.Println(" :: Protocol:	https")
		} else {
			fmt.Println(" :: Protocol:	http")
		}
		fmt.Println("_________________________________\n")
	}

	for number := 0; number < threads; number++ {
		var threads_sites []string
		for i := number; i < len(sites); i = i + threads {
			threads_sites = append(threads_sites, sites[i])
		}
		wg.Add(1)
		go worker(threads_sites)
	}
	wg.Wait()
}
