package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// check if user in blacklist
func CheckBlacklist(u string, c *Configuration) bool {
	for _, i := range c.Blacklist {
		if strings.Contains(u, i) {
			return false
		}
	}

	return true
}

// check if text from IRC line contains search string from config
func (c *Configuration) HasText(line string) bool {

	// loop through slice to find a match
	for _, i := range c.SearchKeywords {
		if strings.Contains(line, i) {
			return true
		}
	}

	return false
}

// Parse line for user, CCP url, and CCP number
func ParseText(line string) string {
	rgx := regexp.MustCompile(`(?i)ccp-\d+`)
	ccp := rgx.FindString(line)
	ccp = strings.ToUpper(ccp)

	return ccp
}

// Loop through slice to find word
func FindKeyword(line string, s []string) bool {
	line = strings.ToLower(line)

	for _, i := range s {
		if strings.Contains(line, i) {
			return true
		}
	}

	return false
}

// Post event to Anthracite
func (c *Configuration) PostEvent(ts int64, user string, m string, tag string) bool {

	// get CCP ticket
	id := ParseText(m)

	// convert timestamp from int64 to string
	string_ts := strconv.FormatInt(ts, 10)

	// set query params
	data := url.Values{}
	data.Add("event_timestamp", string_ts)
	data.Add("event_desc", id+"\n"+user)
	data.Add("event_tags", "CCP"+" "+tag)

	uri, _ := url.ParseRequestURI(c.Anthracite.URL)
	uri.Path = c.Anthracite.Resource
	urlStr := fmt.Sprintf("%v", uri)

	// use for SSL
	tlsConfig := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}

	client := &http.Client{Transport: tlsConfig}
	req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	return true
}
