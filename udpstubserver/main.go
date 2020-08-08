package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := flag.String("p", "8125", "Listening port")
	stdLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lmsgprefix)
	flag.Parse()
	if flag.NArg() > 0 {
		flag.PrintDefaults()
		return
	}
	addr, err := net.ResolveUDPAddr("udp", ":"+*port)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	stdLogger.Printf("Server started and listen on port %s...", *port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		buf := make([]byte, 10240)
		for {
			n, addr, err := listener.ReadFromUDP(buf)
			if err != nil {
				stdLogger.Printf("Error receiving packet: %v", err)
				continue
			}
			stdLogger.Printf("Received: %q from %s", string(buf[0:n]), addr)
		}
	}()

	<-c
	stdLogger.Println("Received interrupt signal. Exiting...")
}
