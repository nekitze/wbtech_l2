package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleCloseConnectionSignal(sigChan <-chan os.Signal, conn net.Conn) {
	go func() {
		for {
			select {
			case <-sigChan:
				disconnect(conn)
			default:
				_, err := conn.Read(nil)
				if err == io.EOF || err != nil {
					disconnect(conn)
				}
			}
		}
	}()
}

func disconnect(conn net.Conn) {
	_ = conn.Close()
	fmt.Print("disconnected from host.")
	os.Exit(0)
}

func startResponseHandler(conn net.Conn) {
	go func() {
		for _, err := io.Copy(os.Stdout, conn); err == nil; {
		}
	}()
}

func startTelnetCommunication(conn net.Conn) {
	startResponseHandler(conn)
	for _, err := io.Copy(conn, os.Stdin); err == nil; {
	}
}

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: telnet-client [--timeout] host port")
		return
	}
	host := flag.Args()[0]
	port := flag.Args()[1]

	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	handleCloseConnectionSignal(sigChan, conn)
	startTelnetCommunication(conn)
}
