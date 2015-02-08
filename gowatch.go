package gowatch

import (
	"errors"
	"net"
	"time"
)

// Watcher able to watch a lot
type Watcher struct {
	WatchChan chan *WatchNode
	Nodes     []*WatchNode
}

// WatchNode for Watcher
type WatchNode struct {
	Watcher     *Watcher
	Name        string
	Group       string
	HostAddress string
	Protocol    string

	interval time.Duration
	ticker   *time.Ticker
	Status   string
	stop     bool
}

// NewWatcher return the new Watcher
func NewWatcher() *Watcher {
	w := new(Watcher)
	w.WatchChan = make(chan *WatchNode)
	return w
}

// Append append the new node to Watcher
// protocol support tcp or udp
// hostAddress is like localhost:3000
// duration simple put time.Duration you would like to pool
func (w *Watcher) Append(protocol, hostAddress string, duration time.Duration) error {
	_, err := NewWatchNode(protocol, hostAddress, duration, w)
	return err
}

//NewWatchNode you could specific the Watcher
func NewWatchNode(protocol, hostAddress string, duration time.Duration, w *Watcher) (*WatchNode, error) {

	if protocol != "tcp" && protocol != "udp" {
		return nil, errors.New("protocol must be tcp or udp")
	}

	node := new(WatchNode)

	node.HostAddress = hostAddress
	node.Protocol = protocol
	if w != nil {
		node.Watcher = w
	}

	node.ticker = time.NewTicker(duration)
	node.interval = duration
	node.stop = true
	node.Start()

	w.Nodes = append(w.Nodes, node)
	return node, nil
}

// Start will kick the time.Ticker start to watch the node
// when the status of the node change wouldi send  notify to Watcher.WatchChan
func (node *WatchNode) Start() {
	go func() {
		for _ = range node.ticker.C {
			_, err := net.Dial(node.Protocol, node.HostAddress)
			if err != nil {
				// bad news
				// this bad news different from previous one
				if node.Status != err.Error() {
					node.Status = err.Error()
					node.Watcher.WatchChan <- node
				}
			} else if node.Status != "" {
				// no news is good news
				node.Status = ""
				node.Watcher.WatchChan <- node
			}
			//fmt.Println("Tick at", time.Now(), node.HostAddress, err)
		}
	}()
}

// Stop ask the WatchNode ticker to stop
func (node *WatchNode) Stop() {
	node.ticker.Stop()
}
