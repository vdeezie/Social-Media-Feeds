package main

import (
	"desire/class3"
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

func main() {
    fb := new(class3.Facebook)
	twtr := new(class3.Twitter)
	lnkdin := new(class3.Linkedin)
	
    feedf := make(map[string]string)
	for i := range fb.Feed() {
		feedf[strconv.Itoa(i+1)] = fb.Feed()[i]
	}
	f, _ := json.MarshalIndent(feedf, "", " ")
	_ = ioutil.WriteFile("fbdata.json", f, 0644)
	fcb, _ := xml.MarshalIndent(fb.Feed(), "", " ")
	_ = ioutil.WriteFile("fbdata.xml", fcb, 0644)

	feedt := make(map[string]string)
	for i := range twtr.Feed() {
		feedt[strconv.Itoa(i+1)] = twtr.Feed()[i]
	}
	t, _ := json.MarshalIndent(feedt, "", " ")
	_ = ioutil.WriteFile("twtr.json", t, 0644)
	tw, _ := xml.MarshalIndent(twtr, "", " ")
	_ = ioutil.WriteFile("twtr.xml", tw, 0644)

	feedl := make(map[string]string)
	for i := range lnkdin.Feed() {
		feedl[strconv.Itoa(i+1)] = lnkdin.Feed()[i]
	}
	in, _ := json.MarshalIndent(feedl, "", " ")
	_ = ioutil.WriteFile("lnkdin.json", in, 0644)
	lin, _ := xml.MarshalIndent(lnkdin, "", " ")
	_ = ioutil.WriteFile("lnkdin.xml", lin, 0644)

}