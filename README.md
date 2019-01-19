# irccat - cat for IRC

[![Build Status](https://cloud.drone.io/api/badges/prologic/irccat/status.svg)](https://cloud.drone.io/prologic/irccat)
[![Go Report Card](https://goreportcard.com/badge/prologic/irccat)](https://goreportcard.com/report/prologic/irccat)
[![GoDoc](https://godoc.org/github.com/prologic/irccat?status.svg)](https://godoc.org/github.com/prologic/irccat) 
[![Sourcegraph](https://sourcegraph.com/github.com/prologic/irccat/-/badge.svg)](https://sourcegraph.com/github.com/prologic/irccat?badge)

`irccat` is a simple tool that reads from standard input (*stdin*) and
posts the read message to a target IRC channel on a specified server.

Useful for simple command-line notifications, part of a pipeline, monitoring
or anything you can pipe to `irccat`. Example:

```#!bash
echo 'Hello World!' | irccat irc.freenode.net:6667 gonuts
```

## Installation

From Source:
```#!bash
$ go -u get github.com/prologic/irccat
```

Using Docker:
```#!bash
$ docker pull prologic/irccat
```

## Usage

From Source:
```#!bash
$ echo '<message>' | irccat <address>:<port> <channel>
```

Using Docker:
```#!bash
$ echo '<message>' | docker run -i -t prologic/irccat <address>:<port> <channel>
```

## License

`irccat` is licensed under the MIT License.
