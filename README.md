# GoBot - A Go IRC Bot

GoBot is an IRC bot written in Go utlizing a Go IRC framework that listens in on an IRC channel and performs an event (post to Anthracite) when specific word(s) are seen. 

## Compiling ##

```
go build -v -o botname
```

Go build will create a binary based on OS architecture - i386 or x86_64

## Usage ##

```
./gobot -c [file.json]
```

GoBot needs a config file in json format. Make sure a proper config.json is in the same directory as your executable. See [config.json](https://gitlab.internal.tinyprints.com/iops/gobot/blob/master/config.json) and make modifications as needed. The program isn't ready for additions or subtractions to the current config.json without making modification to main.go. Hopefully in future versions, modications will be handled with grace.
The fluffle/GoIRC framework only supports connections to a single channel. You will need to run multiple GoBots and passing in a different config json with modifications to point to the correct channel.

## Package in a Nutshell ##

Reading the source will help you understand the package, but below is a nutshell on what the bot is doing.

1. Reads in a config json
2. Sets IRC config, utilizing the GoIRC package
3. Connects to IRC server and connects to channel
4. Listens for 'privmsg' and searches for CCP
5. If it sees CCP, sets 'start' or 'end' tag
6. Posts CCP ticket and user to Anthracite 

## Wants ##

Make the bot more general
- see the words and do an event, instead of specifically posting to Anthracite
- handle additions and subtractions to config.json
- pass in a config file via a flag
Logging preferrably.
