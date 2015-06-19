// author: Paul Clata
// gitlab.internal.tinyprints.com/iops/gobot
package main

import (
	"flag"
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"time"
)

var conf *Configuration

func main() {

	var config string
	flag.StringVar(&config, "c", "", "config file in json format.")
	flag.Parse()

	conf, err := GetConfig(config)

	if err != nil {
		fmt.Println("Config file does not exist. Please pass in a valid config file")
	}

	cfg := irc.NewConfig(conf.Bot.Name)
	cfg.SSL = false
	cfg.Server = conf.IRCserver.Server
	cfg.NewNick = func(n string) string { return n + "^" }
	c := irc.Client(cfg)

	quit := make(chan bool)

	c.HandleFunc("connected", func(conn *irc.Conn, line *irc.Line) {
		conn.Join(conf.IRCserver.Channels)
		fmt.Println("Connected to server")
	})

	c.HandleFunc("disconnected", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("Disconnected to server")
		quit <- true
	})

	c.HandleFunc("privmsg", func(conn *irc.Conn, line *irc.Line) {
		now := time.Now()
		timestamp := now.Unix()
		msg := line.Text()
		user := line.Nick

		// if user is another IRC bot, it shall not pass
		if CheckBlacklist(user, conf) {

			if conf.HasText(msg) {

				if FindKeyword(msg, conf.StartKeywords) {
					tag := "start"

					// send event to anthracite
					conf.PostEvent(timestamp, user, msg, tag)
				} else if FindKeyword(msg, conf.EndKeywords) {
					tag := "end"

					// send event to anthracite
					conf.PostEvent(timestamp, user, msg, tag)
				}
			}
		}
	})

	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}

	<-quit
}
