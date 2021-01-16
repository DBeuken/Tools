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
)

var wg sync.WaitGroup

func ScanPort(hostname string, port int, delay int) string {
	conn, err := net.DialTimeout("tcp", hostname + ":" + strconv.Itoa(port), time.Duration(delay) * time.Second)
	if (err != nil) {
		return "CLOSED"
	}

	defer conn.Close()

	return "OPEN"
}

func worker(poortenarr []int, target string, delay int) {
	for _,poort := range poortenarr {
		open := ScanPort(target, poort, delay)
		fmt.Printf("Port %v: %v\n", poort, open)
	}
	wg.Done()
}

func main() {
	var target string
	flag.StringVar(&target, "t", "-", "Specify the target")
	var filetarget string
	flag.StringVar(&filetarget, "ft", "-", "Specify a file containing targets")
	var delay int
	flag.IntVar(&delay, "d", 5, "Specify the time to look for a connection\nDefault = 5")
	var threads int
	flag.IntVar(&threads, "th", 200, "Specify amount of threads\nDefault: 200")
	var poorten string
	flag.StringVar(&poorten, "p", "-", "Specify the port(s)\nDefault: 5-9 == 5,6,7,8,9")
	nobanner := flag.Bool("nobanner", false, "Don't show the banner")
	flag.Parse()


	//CHECK TARGET
	if (target == "-") {
		fmt.Println("You need to specify a target")
		os.Exit(0)
	}

	//GET PORTS
	var poortenarr []int

	if (poorten == "-") {
		poortenarr = []int{22, 80, 111, 122}
	} else {
		var poortenarrcomm = strings.Split(poorten, ",")
		for _,element := range poortenarrcomm {
			split := strings.Split(element, "-")
			if len(split) == 2 {
				minpoort, err := strconv.Atoi(split[0])
				if (err != nil) {
					fmt.Println("You can only use numbers to specify the ports")
					os.Exit(0)
				}
				maxpoort, err := strconv.Atoi(split[1])
				if (err != nil) {
					fmt.Println("You can only use numbers to specify the ports")
					os.Exit(0)
				}
				for i := minpoort; i <= maxpoort; i++ {
					poortenarr = append(poortenarr, i)
				}
			} else {
				make_int, err := strconv.Atoi(split[0])
				if (err != nil) {
					fmt.Println("You can only use numbers to specify the ports")
					os.Exit(0)
				}
				poortenarr = append(poortenarr, make_int)
			}
		}
	}

	var banner = `______________________________________________
                              ____    ____      
                             /\  _ \ /\  _ \    
  ____    ___     __      ___\ \ \/\ \ \ \L\ \  
 /',__\  /'___\ /'__ \  /' _  \ \ \ \ \ \  _ <'
/\__,  \/\ \__//\ \L\.\_/\ \/\ \ \ \_\ \ \ \L\ \
\/\____/\ \____\ \__/.\_\ \_\ \_\ \____/\ \____/
 \/___/  \/____/\/__/\/_/\/_/\/_/\/___/  \/___/ 
 _______________________________________________
                                                
	`

	if (*nobanner != true) {
		fmt.Println(banner)
	}

	//START SCANNING
	for number := 0; number < threads; number++ {
		var thread_ports []int
		for i := number; i < len(poortenarr); i = i + threads {
			thread_ports = append(thread_ports, poortenarr[i])
		}
		wg.Add(1)
		go worker(thread_ports, target, delay)
	}
	wg.Wait()



	/*
	for _,poort := range poortenarr {
		open := ScanPort(target, poort, delay)
		fmt.Printf("Port %v: %v\n", poort, open)
	}
	*/
}
