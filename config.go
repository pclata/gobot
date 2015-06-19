package main

import (
	"encoding/json"
	"os"
)

// configuration contains the ccpbot configuration
type Configuration struct {
	Bot struct {
		Name string `json:"name"`
	} `json:"bot"`

	IRCserver struct {
		Server   string `json:"server"`
		Port     string `json:"port"`
		Channels string `json:"channels"`
	} `json:"irc_server"`

	Anthracite struct {
		URL      string `json:"URL"`
		Resource string `json:"resource"`
	} `json:"anthracite"`

	MatchStr struct {
		searchStr string `json:"search_string"`
	} `json:"match_str"`

	SearchKeywords []string `json:"search_keywords"`

	StartKeywords []string `json:"start_keywords"`

	EndKeywords []string `json:"end_keywords"`

	Blacklist []string `json:"blacklist"`
}

// GetConfig builds a config obj
func GetConfig(cf string) (*Configuration, error) {
	confFile, err := os.Open(cf)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(confFile)
	conf := &Configuration{}
	err = decoder.Decode(conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
