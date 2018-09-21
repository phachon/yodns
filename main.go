package main

import (
	"github.com/spf13/pflag"
	"net"
	"log"
	"os"
	"io"
)

// your own dns server

var (
	local = pflag.String("local", ":53", "please input dns server listen addr")
	remote = pflag.String("remote", "", "please remote dns server addr")
)

func init()  {
	pflag.Parse()
	if *remote == "" {
		log.Println("remote dns server addr is not empty!")
		os.Exit(100)
	}
}

func main()  {

	addr, err := net.ResolveUDPAddr("udp", *local)
	if err != nil {
		log.Printf("reslove udp addr error: %s ", err.Error())
		os.Exit(100)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Printf("udp listen error: %s ", err.Error())
		os.Exit(100)
	}

	handle(conn)
}

func handle(conn *net.UDPConn) {

	log.Println("udp server start ")

	// remote dns server
	remoteConn, err := net.Dial("udp", *remote)
	if err != nil {
		log.Printf("remote dns server conn error %s", err.Error())
		os.Exit(100)
	}

	// io bind
	go func() {
		io.Copy(conn, remoteConn)
	}()
	io.Copy(remoteConn, conn)
}

