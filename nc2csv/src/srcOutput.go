package main

import (
	"fmt"
	"log"
)

func srcOutput(in chan string, out chan string, done chan bool) {
	output := ""
	for {
		text, flag := <-in
		if flag {
			output += fmt.Sprintf("%s\n", text)
		} else {
			output += fmt.Sprintf("\n")
			break
		}
	}
	log.Println(output)
	out <- output
	done <- true
	return
}
