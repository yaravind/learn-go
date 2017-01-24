package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" //import to register handler for debug/pprof
	"time"
)

//Access the profile details at http://localhost:6060/debug/pprof/
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	time.Sleep(1 * time.Minute)
}
