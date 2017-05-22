package chat

import (
	"fmt"
	"io"
	"sync"
)

// Room - chat room structure
type Room struct {
	name    string
	Msgch   chan string
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

// CreateRoom - make new chat room
func CreateRoom(name string) *Room {
	r := &Room{
		name:    name,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[chan<- string]struct{}),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

// AddClient - add new client
func (r *Room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	//remove client when done is signalled
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

// ClCount - room count
func (r *Room) ClCount() int {
	return len(r.clients)
}

// RemoveClient - get rid of client
func (r *Room) RemoveClient(wc chan<- string) {
	loggy.Rec("Removing client ", nil)
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()
	select {
	case <-r.Quit:
		if len(r.clients) == 0 {
			close(r.Msgch)
		}
	default:
	}
}

// Run chat room
func (r *Room) Run() {
	loggy.Rec("Starting chat room", r.name)
	go func() {
		for msg := range r.Msgch {
			r.broadcastMsg(msg)
		}
	}()
}

func (r *Room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	for wc := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
