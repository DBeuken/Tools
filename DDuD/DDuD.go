package main

import (
	"fmt"
	"net"
	"os"
	"flag"
	"strings"
	"strconv"
	"time"
	"sync"
	"io/ioutil"
	"sort"
)

var banner = `  _____   _____            _____     
 /\  _ \ /\  _ \          /\  _ \
 \ \ \/\ \ \ \/\ \  __  __\ \ \/\ \
  \ \ \ \ \ \ \ \ \/\ \/\ \\ \ \ \ \
   \ \ \_\ \ \ \_\ \ \ \_\ \\ \ \_\ \
    \ \____/\ \____/\ \____/ \ \____/
     \/___/  \/___/  \/___/   \/___/

         v. 1.0.0
_____________________________________
`

var helpmenu = `_________________________________________________________
  _    _      _                                  
 | |  | |    | |                                 
 | |__| | ___| |_ __  _ __ ___   ___ _ __  _   _ 
 |  __  |/ _ \ | '_ \| '_ ' _ \ / _ \ '_ \| | | |
 | |  | |  __/ | |_) | | | | | |  __/ | | | |_| |
 |_|  |_|\___|_| .__/|_| |_| |_|\___|_| |_|\__,_|
               | |                               
               |_|                               
_________________________________________________________

This tool is used to scan ports. You can scan one target as well as multiple targets.

TAGS_____________________________________________________

   -t		Target. You can specify an IP-address or a domain name.
   -ft		File containing targets. If you select this file, every line will be seen as a target.
   -d		Delay. Time to wait when trying to connect to a specific port (specified in seconds).
 		Default = 5 seconds
   -th		Threads. Amount of threads to make connections.
 		Default = 200
   -p		Ports. Specify the ports to scan. This will count for every target.
 		If you specify it like '20-30', every port from 20 to 30 will be scanned.
		If you specify it like '20,30', only port 20 and 30 will be scanned.
  		Default = top 1000 most-commonly used ports
   -nobanner	Dont show the banner.
   -open	Only show open ports
 		Default = show both closed and open ports.
   -sequence	Show ports sequentially.
   -help	Shows this helpmenu.



EXAMPLE_SYNTAX___________________________________________

   DDuD -t youtube.com
   DDuD -ft sites.txt -th 20 -d 10
   DDuD -t localhost -p 20-30,443 -open
   DDUD -t google.com -sequence

`

var wg sync.WaitGroup

var scanned []string
var scannedall []string

func ScanPort(hostname string, port int, delay int) string {
	conn, err := net.DialTimeout("tcp", hostname + ":" + strconv.Itoa(port), time.Duration(delay) * time.Second)
	if (err != nil) {
		return "CLOSED"
	}
	defer conn.Close()
	return "OPEN  "
}

func worker(portsarr []int, target string, delay int) {
	for _,port := range portsarr {
		open := ScanPort(target, port, delay)
		fmt.Printf("Port %v: %v\n", port, open)
	}
	wg.Done()
}

func worker2(portsarr []int, target string, delay int) {
	for _,port := range portsarr {
		open := ScanPort(target, port, delay)
		if (open == "OPEN  ") {
			line := fmt.Sprintf("Port %v: %v", port, open)
			scanned = append(scanned, line)
		}
	}
	wg.Done()
}

func worker3(portsarr []int, target string, delay int) {
	for _,port := range portsarr {
		open := ScanPort(target, port, delay)
		line := fmt.Sprintf("Port %v: %v", port, open)
		scannedall = append(scannedall, line)
	}
	wg.Done()
}

func main() {
	//MAKE FLAGS
	var target string
	flag.StringVar(&target, "t", "-", "Specify the target")
	var filetarget string
	flag.StringVar(&filetarget, "ft", "-", "Specify a file containing targets")
	var delay int
	flag.IntVar(&delay, "d", 5, "Specify the time to look for a connection\nDefault = 5")
	var threads int
	flag.IntVar(&threads, "th", 200, "Specify amount of threads\nDefault: 200")
	var ports string
	flag.StringVar(&ports, "p", "-", "Specify the port(s)\nDefault: 5-9 == 5,6,7,8,9")
	sequence := flag.Bool("sequence", false, "Show the ports in sequence")
	open := flag.Bool("open", false, "Only show open ports")
	nobanner := flag.Bool("nobanner", false, "Don't show the banner")
	help := flag.Bool("help", false, "Show the helpmenu")
	flag.Parse()

	//HELPMENU
	if (*help == true) {
		fmt.Println(helpmenu)
		os.Exit(0)
	}

	time_start := time.Now()
	var targetarr []string

	//CHECK TARGET
	if (target == "-" && filetarget == "-") {
		fmt.Println("You need to specify a target")
		os.Exit(0)
	}
	if (target != "-") {
		targetarr = append(targetarr, target)
	}
	if (filetarget != "-") {
		file, err := ioutil.ReadFile(filetarget)
		if (err != nil) {
			fmt.Println("ERROR:\nCannot open file with threads")
			os.Exit(0)
		}
		file_str := string(file)
		file_split := strings.Split(file_str, "\n")
		for i:= 0; i < len(file_split) -1; i++ {
			targetarr = append(targetarr, file_split[i])
		}
	}

	//GET PORTS
	var portsarr []int

	if (ports == "-") {
		portsarr = []int{22, 80, 111, 122}
	} else {
		var portsarrcomm = strings.Split(ports, ",")
		for _,element := range portsarrcomm {
			split := strings.Split(element, "-")
			if len(split) == 2 {
				minport, err := strconv.Atoi(split[0])
				if (err != nil) {
					fmt.Println("You can only use numbers to specify the ports")
					os.Exit(0)
				}
				maxport, err := strconv.Atoi(split[1])
				if (err != nil) {
					fmt.Println("You can only use numbers to specify the ports")
					os.Exit(0)
				}
				for i := minport; i <= maxport; i++ {
					portsarr = append(portsarr, i)
				}
			} else {
				make_int, err := strconv.Atoi(split[0])
				if (err != nil) {
					fmt.Println("You can only use numbers to specify the ports")
					os.Exit(0)
				}
				portsarr = append(portsarr, make_int)
			}
		}
	}

	//PRINT BANNER
	if (*nobanner != true) {
		fmt.Println(banner)
		fmt.Printf(" :: Target 	:   %v\n", target)
		fmt.Printf(" :: Filetarget  :   %v\n", filetarget)
		fmt.Printf(" :: Threads	:   %v\n", threads)
		fmt.Printf(" :: Delay	:   %v\n", delay)
		fmt.Printf("_____________________________________\n")
	}

	//START SCANNING
	if (*sequence == false && *open == false) {
		for _,targetslice := range targetarr {
			fmt.Printf("\nStarting with target %v...\n\n", targetslice)
			for number := 0; number < threads; number++ {
				var thread_ports []int
				for i := number; i < len(portsarr); i = i + threads {
					thread_ports = append(thread_ports, portsarr[i])
				}
				wg.Add(1)
				go worker(thread_ports, targetslice, delay)
			}
			wg.Wait()
		}
	} else if (*sequence == false && *open == true) {
		for _,targetslice := range targetarr {
			fmt.Printf("\nStarting with target %v...\n\n", targetslice)
			for number := 0; number < threads; number++ {
				var thread_ports []int
				for i := number; i < len(portsarr); i = i + threads {
					thread_ports = append(thread_ports, portsarr[i])
				}
				wg.Add(1)
				go worker2(thread_ports, targetslice, delay)
			}
			wg.Wait()
		}
		for _,element := range scanned {
			fmt.Println(element)
		}
	} else if (*sequence == true && *open == false) {
		for _,targetslice := range targetarr {
			fmt.Printf("\nStarting with target %v...\n\n", targetslice)
			for number := 0; number < threads; number++ {
				var thread_ports []int
				for i := number; i < len(portsarr); i = i + threads {
					thread_ports = append(thread_ports, portsarr[i])
				}
				wg.Add(1)
				go worker3(thread_ports, targetslice, delay)
			}
			wg.Wait()
		}
		sort.Strings(scannedall)
		for _,element := range scannedall {
			if (len(element) == 14) {
				fmt.Println(element)
			}
		}
		for _,element := range scannedall {
			if (len(element) == 15) {
				fmt.Println(element)
			}
		}
		for _,element := range scannedall {
			if (len(element) == 16) {
				fmt.Println(element)
			}
		}
		for _,element := range scannedall {
			if (len(element) == 17) {
				fmt.Println(element)
			}
		}
		for _,element := range scannedall {
			if (len(element) == 18) {
				fmt.Println(element)
			}
		}
	}
	fmt.Printf("\nCode finished in %v\n", time.Since(time_start))
}
