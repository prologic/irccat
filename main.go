package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/namsral/flag"
	"github.com/thoj/go-ircevent"
)

const (
	username = "irccat"
	realname = "IRC Cat, cat for IRC"
)

type Addr struct {
	Host   string
	Port   int
	UseTLS bool
}

func (a *Addr) String() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func ParseAddr(s string) (addr *Addr, err error) {
	addr = &Addr{}

	parts := strings.Split(s, ":")
	fmt.Printf("%v", parts)
	if len(parts) != 2 {
		return nil, fmt.Errorf("malformed address: %s", s)
	}

	addr.Host = parts[0]

	if parts[1][0] == '+' {
		port, err := strconv.Atoi(parts[1][1:])
		if err != nil {
			return nil, fmt.Errorf("invalid port: %s", parts[1])
		}
		addr.Port = port
		addr.UseTLS = true
	} else {
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid port: %s", parts[1])
		}
		addr.Port = port
	}

	if addr.Port < 1 || addr.Port > 65535 {
		return nil, fmt.Errorf("invalid port: %d", addr.Port)
	}

	return addr, nil
}

func main() {
	var (
		err error

		version bool
		config  string
		debug   bool

		nickname string
	)

	flag.BoolVar(&version, "v", false, "display version information")
	flag.StringVar(&config, "c", "", "config file")
	flag.BoolVar(&debug, "d", false, "debug logging")

	flag.StringVar(&nickname, "n", "irccat", "nick to use")

	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if version {
		fmt.Printf("irccat v%s", FullVersion())
		os.Exit(0)
	}

	if flag.NArg() < 2 {
		log.Fatalf("Ussage: %s <address>[:port] <channel>", os.Args[0])
	}

	addr, err := ParseAddr(flag.Arg(0))
	if err != nil {
		log.Fatalf("error parsing addr: %s", err)
	}

	channel := flag.Arg(1)
	if channel[0] != '#' {
		log.Warnf("assuming channel #%s", channel)
		channel = "#" + channel
	}

	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	if isPrefix {
		log.Warn("message may be too long")
	}
	message := string(line)

	conn := irc.IRC(nickname, username)
	conn.RealName = realname

	conn.VerboseCallbackHandler = debug
	conn.Debug = debug

	conn.UseTLS = addr.UseTLS
	conn.KeepAlive = 30 * time.Second
	conn.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	conn.AddCallback("001", func(e *irc.Event) {
		log.Info("Connected!")
		conn.Join(channel)
		conn.Privmsg(channel, message)
		conn.Quit()
	})

	err = conn.Connect(addr.String())
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	conn.Loop()
}
