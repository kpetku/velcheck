package main

import (
	"log"
	"os"
)

func main() {
	v := &Vel{}
	if len(os.Args) > 1 {
		v.New(os.Args[1])
	} else {
		log.Printf("not enough arguments: ./velcheck http://example.org/file.json or ./velcheck file.json -- using https://vel.myjoomla.io as default.")
		v.New("https://vel.myjoomla.io")
	}
	v.Marshall()
	v.ScanVhosts()
}
