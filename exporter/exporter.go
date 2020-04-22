package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	week3 "go_feed_export"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// ScrollFeeds prints all social media feeds
func ScrollFeeds(platforms ...week3.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}

// ExportTXT writes all feed into corresponding txt files
func ExportTXT(u week3.SocialMedia, filename string) error {
	f, err := os.OpenFile("./files_txt/"+filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	fmt.Println()
	return nil
}

// FeedPresentation struct holds an array for each feed
type FeedPresentation struct {
	FeedStream []string
	FeedCount  int
}

// ExportJSON writes all feed into corresponding json files
func ExportJSON(u week3.SocialMedia, filename string) error {
	data := FeedPresentation{}
	f, err := os.OpenFile("./files_json/"+filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		data.FeedStream = append(data.FeedStream, fd)
	}
	data.FeedCount = len((data.FeedStream))
	file, _ := json.MarshalIndent(data, "", "  ")
	n, err := f.Write([]byte(file))
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes of json\n", n)
	fmt.Println()
	return nil
}

// ExportXML writes all feed into corresponding xml files
func ExportXML(u week3.SocialMedia, filename string) error {
	data := FeedPresentation{}
	f, err := os.OpenFile("./files_xml/"+filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		data.FeedStream = append(data.FeedStream, fd)
	}
	data.FeedCount = len((data.FeedStream))
	file, _ := xml.MarshalIndent(data, "", "  ")
	n, err := f.Write([]byte(file))
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes of xml\n", n)
	fmt.Println()
	return nil
}

// ExportYAML writes all feed into corresponding xml files
func ExportYAML(u week3.SocialMedia, filename string) error {
	data := FeedPresentation{}
	f, err := os.OpenFile("./files_yaml/"+filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		data.FeedStream = append(data.FeedStream, fd+"\n")
	}
	data.FeedCount = len((data.FeedStream))
	file, _ := yaml.Marshal(data)
	n, err := f.Write([]byte(file))
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes of yaml\n", n)
	fmt.Println()
	return nil
}
