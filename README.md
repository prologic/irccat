# irccat - cat for IRC

`irccat` is a simple tool that reads from standard input (*stdin*) and
posts the read message to a target IRC channel on a specified server.

Useful for simple command-line notifications, part of a pipeline, monotiring
or anything you can pipe to `ircat`. Example:

```#!bash
echo 'Hello World!' | irccat irc.freenode.net gonuts
```

## Installation

```#!bash
$ go get github.com/prologic/irccat
```

## License

`irccat` is licensed under the MIT License.
