package gowatch

import (
	"testing"
	"time"
)

const (
	SuccessAddr = "www.google.com:80"

	//i ASSUME THIS WILL FAIL
	FailAddr = "localhost:33333"
)

func TestShouldSuccess(t *testing.T) {
	watcher := NewWatcher()
	err := watcher.Append("tcp", SuccessAddr, time.Second)

	if err != nil {
		t.Errorf("Whatcher create err %q", err)
	}

	select {

	case node := <-watcher.WatchChan:
		if node.Status != "" {
			t.Errorf("Whatcher test should running with networking, watcher fail %q", node.Status)
		}
	case <-time.After(time.Second * 2):
		return
	}

}

func TestShouldFail(t *testing.T) {

	watcher := NewWatcher()
	err := watcher.Append("tcp", FailAddr, time.Second)

	if err != nil {
		t.Errorf("Whatcher create err %q", err)
	}

	select {

	case node := <-watcher.WatchChan:
		if node.Status != "" {
			return
		}
	case <-time.After(time.Second * 2):
		t.Errorf("Whatcher should have fail Status for %q", FailAddr)
	}
}
