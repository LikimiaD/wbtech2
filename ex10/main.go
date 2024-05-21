package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	timeout = flag.String("timeout", "10s", "Connection timeout")
)

func handleClient(conn net.Conn, done chan bool) {
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			_, err := conn.Write(scanner.Bytes())
			if err != nil {
				slog.Warn("Failed to write to connection", "error", err.Error())
				done <- true
				return
			}
		}
		if err := scanner.Err(); err != nil {
			slog.Warn("Error reading from stdin", "error", err.Error())
		}
		done <- true
	}()

	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		slog.Warn("Failed to read from connection", "error", err.Error())
	}
	done <- true
}

func connection(done chan bool, duration time.Duration, host, port string) error {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, duration)
	if err != nil {
		return err
	}
	defer conn.Close()

	slog.Info("Connected to server", "address", address)

	handleClient(conn, done)
	return nil
}

func main() {
	flag.Parse()

	var (
		duration time.Duration = 10 * time.Second
		err      error
	)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		<-signalChannel
		slog.Info("Received signal, closing app...")
		done <- true
	}()

	if flag.NArg() < 2 {
		fmt.Println("usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	if *timeout != "" {
		duration, err = time.ParseDuration(*timeout)
		if err != nil {
			fmt.Printf("Error while parsing timeout: %s\n", err.Error())
			os.Exit(1)
		}
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	if err := connection(done, duration, host, port); err != nil {
		slog.Error("Error while connecting", "error", err.Error())
	}

	<-done
	slog.Info("Connection closed, exiting")
}
