package main

import (
	"fmt"

	"go_feed_export/exporter"

	"go_feed_export/facebook"
	"go_feed_export/linkedin"
	"go_feed_export/twitter"
)

func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)

	// Create text files for feed
	err := exporter.ExportTXT(fb, "fbdata.txt")
	err = exporter.ExportTXT(twtr, "twtrdata.txt")
	err = exporter.ExportTXT(lnkdin, "lnkdin.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------")

	// Create json files for feed
	err = exporter.ExportJSON(fb, "fbdata.json")
	err = exporter.ExportJSON(twtr, "twtrdata.json")
	err = exporter.ExportJSON(lnkdin, "lnkdin.json")

	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------")

	// Create xml files for feed
	err = exporter.ExportXML(fb, "fbdata.xml")
	err = exporter.ExportXML(twtr, "twtrdata.xml")
	err = exporter.ExportXML(lnkdin, "lnkdin.xml")

	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------")

	// Create yaml files for feed
	err = exporter.ExportYAML(fb, "fbdata.yaml")
	err = exporter.ExportYAML(twtr, "twtrdata.yaml")
	err = exporter.ExportYAML(lnkdin, "lnkdin.yaml")

	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------")

	// Print feed to screen
	exporter.ScrollFeeds(fb, twtr, lnkdin)
}
