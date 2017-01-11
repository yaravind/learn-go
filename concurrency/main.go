package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	nBytes, nChunks := int64(0), int64(0)

	f, err := os.Open("c:/aravind/resume.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	buf := make([]byte, 0, 1024)
	log.Println("opend reader for the file")
	textCh := make(chan []byte, 2)

	go func() {
		defer close(textCh) //close before returning from goroutine
		for {
			n, err := r.Read(buf[:cap(buf)])
			buf = buf[:n]
			if n == 0 {
				if err == nil {
					continue
				}
				if err == io.EOF {
					break
				}
			}
			nChunks++
			nBytes += int64(len(buf))
			log.Printf("nChunks=%d, nBytes=%d\n", nChunks, nBytes)
			// process buf
			textCh <- buf

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
		}
		log.Printf("done from goroutine len=%d\n", len(textCh))
	}()

	histogram := make(map[string]int)
	for {
		if bytes, open := <-textCh; open {
			words := strings.Split(string(bytes), " ")
			log.Printf("waiting for data on channel len=%d", len(textCh))
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		} else {
			break
		}
	}

	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
