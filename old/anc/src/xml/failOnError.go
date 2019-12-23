package main

import "log"

func FailOnError(err error, file string) {
	if err != nil {
		log.Fatal("Error ", file, " : ", err)
	}
}
