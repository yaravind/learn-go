package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	go PrintStats()
	time.Sleep(1 * time.Second)

	if hn, err := os.Hostname(); err == nil {
		fmt.Printf("Host: %s\n", hn)

		if addrs, err := net.LookupHost(hn); err == nil {
			for _, addr := range addrs {
				fmt.Printf("IP Address: %s\n", addr)
			}
		}
	}
	time.Sleep(1 * time.Second)
	_, pwd := os.Getwd()
	fmt.Println(pwd)
	fmt.Printf("PID: %d\n", os.Getpid())

	checkDep("cmd.exe")
}

func checkDep(name string) error {
	if p, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("coudln't find %s in PATH:%s", name, err)
	} else {
		fmt.Printf("found %s in PATH: %s", name, p)
		return nil
	}
}

func PrintStats() {
	log.Println("Num of CPU's: ", runtime.NumCPU())

	m := &runtime.MemStats{}

	for {
		r := runtime.NumGoroutine()
		log.Println("Num of goroutines: ", r)

		runtime.ReadMemStats(m)
		log.Println("Allocated memory: ", m.Alloc)
		time.Sleep(1 * time.Second)
	}
}