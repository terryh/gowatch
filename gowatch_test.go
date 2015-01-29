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
	watcher, err := NewWatcher("tcp", SuccessAddr, time.Second)

	if err != nil {
		t.Errorf("Whatcher create err", err)
	}

	select {

	case msg := <-watcher.Status:
		if msg != "" {
			t.Errorf("Whatcher test should running with networking, watcher fail", msg)
		}
	case <-time.After(time.Second * 2):
		return
	}

}

func TestShouldFail(t *testing.T) {

	watcher, err := NewWatcher("tcp", FailAddr, time.Second)

	if err != nil {
		t.Errorf("Whatcher create err", err)
	}

	select {

	case msg := <-watcher.Status:
		if msg != "" {
			return
		}
	case <-time.After(time.Second * 2):
		t.Errorf("Whatcher should have fail Status for", FailAddr)
	}
}
