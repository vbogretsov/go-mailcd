package mock

import (
	"sync"

	"github.com/vbogretsov/go-mailcd"
)

// Sender represent mock for mailcd.Sender.
type Sender struct {
	mutex sync.Mutex
	Inbox []mailcd.Request
	Error error
}

// New creates new sender mock.
func New() *Sender {
	return &Sender{
		mutex: sync.Mutex{},
		Inbox: []mailcd.Request{},
		Error: nil,
	}
}

// Send implements mailcd.Sender.Send.
func (s *Sender) Send(req mailcd.Request) error {
	s.mutex.Lock()
	s.Inbox = append(s.Inbox, req)
	s.mutex.Unlock()
	return s.Error
}

// Close implements mailcd.Sender.Close.
func (s *Sender) Close() error {
	return nil
}

// Reset clears inbox.
func (s *Sender) Reset() {
	s.mutex.Lock()
	s.Inbox = []mailcd.Request{}
	s.Error = nil
	s.mutex.Unlock()
}
