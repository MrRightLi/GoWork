package main

import (
	lllog "GoWork/log"
)

func main() {
	// log
	lllog.Trace.Println("I have something standard to say")
	lllog.Info.Println("Special Information")
	lllog.Warning.Println("There is something you need to know about")
	lllog.Error.Println("Something has failed")
}
