package gowatch

import (
	"errors"
	"net"
	"time"
)

const (
	OK = "OK"
)

type Watcher struct {
	HostAddress string
	Protocol    string
	Status      chan string

	interval   time.Duration
	ticker     *time.Ticker
	lastStatus string
	stop       chan bool
}

func NewWatcher(protocol, hostAddress string, duration time.Duration) (*Watcher, error) {

	if protocol != "tcp" && protocol != "udp" {
		return nil, errors.New("protocol must be tcp or udp")
	}

	watch := new(Watcher)

	watch.HostAddress = hostAddress
	watch.Protocol = protocol
	watch.Status = make(chan string)

	watch.ticker = time.NewTicker(duration)
	watch.interval = duration
	watch.stop = make(chan bool)
	watch.Run()
	return watch, nil
}

func (w *Watcher) Run() {
	go func() {
		for _ = range w.ticker.C {
			_, err := net.Dial(w.Protocol, w.HostAddress)
			if err != nil {
				w.lastStatus = err.Error()
				w.Status <- w.lastStatus
			} else if w.lastStatus != "" {
				w.lastStatus = ""
				w.Status <- OK
			}
			//fmt.Println("Tick at", t, w.HostAddress, con, err)
		}
	}()
}

func (w *Watcher) Stop() {
	w.ticker.Stop()
}
