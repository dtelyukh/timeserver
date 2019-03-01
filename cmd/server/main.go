package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "37", "The time server port")
	flag.Parse()

	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port specification")
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	log.Println("Request from " + conn.RemoteAddr().String())
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(time.Now().Unix()))
	conn.Write(bs)
}
