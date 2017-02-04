package main

import (
	"flag"
)

func main() {
	feed := flag.String("feed", "https://vel.myjoomla.io/", "URL or full path to file containing a Joomla! Extensions Directory JSON feed")
	vhost := flag.String("vhost", "/mnt/data/vhosts/", "Location of the folder that contains websites")
	docroot := flag.String("docroot", "httpdocs", "Name of the folder (like httpdocs, htdocs, or public_html) containing an installation of Joomla!")

	flag.Parse()

	v := &Vel{}
	v.New(*feed)
	v.Scan(*vhost, *docroot)
}
