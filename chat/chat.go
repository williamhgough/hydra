package chat

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/williamhgough/hydra/logger"
)

var loggy = logger.GetInstance()

// Run hydra chat service
func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	if err != nil {
		loggy.Rec("Error connecting to chat client", err)
		return err
	}
	r := CreateRoom("HydraChat")
	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			loggy.Rec("Error accepting connection from chat client", err)
			break
		}
		go handleConnection(r, conn)
	}

	return err
}

func handleConnection(r *Room, c net.Conn) {
	loggy.Rec("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
