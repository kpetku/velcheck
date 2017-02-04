package main

import (
	"io/ioutil"
	"log"
	"os"
)

func (v *Vel) Scan(virtualHostLocation, documentRoot string) {
	var results int
	var suffix, pwd string

	directories, err := ioutil.ReadDir(virtualHostLocation)
	if err != nil {
		log.Fatalf("Unable to read %s: %s", virtualHostLocation, err.Error())
	}
	for _, vhost := range directories {
		for extension := range VelMap {
			switch extension[0:3] {
			case "com":
				suffix = extension[4:] + ".xml"
				pwd = virtualHostLocation + vhost.Name() + "/" + documentRoot + "/administrator/components/" + extension + "/" + suffix
			case "mod":
				suffix = extension[4:] + ".xml"
				pwd = virtualHostLocation + vhost.Name() + "/" + documentRoot + "/modules/" + extension + "/" + suffix
			case "plg":
				suffix = extension[4:] + ".xml"
				pwd = virtualHostLocation + vhost.Name() + "/" + documentRoot + "/plugins/" + extension + "/" + suffix
			}
			if _, err := os.Stat(pwd); err == nil {
				installedVersion, err := obtainVersion(pwd)
				if err != nil {
					log.Fatalf("%s", err)
				}
				found, info, err := v.check(extension, installedVersion)
				if err != nil {
					log.Fatalf("%s", err)
				}
				if found {
					log.Printf("WARNING: %s has %s version %s installed: %s", vhost.Name(), extension, installedVersion, info)
					results++
				}
			}
		}
	}
	log.Printf("Found: %d result(s)", results)
}
