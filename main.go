package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	mode      string
	zkAddress string
)

func init() {
	flag.StringVar(&mode, "mode", "consul", "the mode to operate in")
	flag.StringVar(&zkAddress, "zk", ":2181", "the ip:port for zookeeper to listen on")
}

func main() {
	flag.Parse()

	validateConfiguration()

	listen, err := net.Listen("tcp", zkAddress)
	if err != nil {
		fmt.Printf("Failed to open socket on %s: %q\n", zkAddress, err)
		os.Exit(1)
	}

	server, err := NewServer(listen)
	if err != nil {
		fmt.Printf("During server creation: %q\n", err)
		os.Exit(1)
	}

	if err := server.Listen(); err != nil {
		fmt.Printf("During listening: %q\n", err)
	}
}

func validateConfiguration() {

}
