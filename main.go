package main

import (
	"log"
)

var TOKEN string
var CHANNEL_ID string
var STUDENT_ID string
var PASSWORD string

func main() {
	if err := Process(); err != nil {
		log.Print(err)
	}
}
