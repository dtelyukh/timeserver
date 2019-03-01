package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)

	if len(host) == 0 {
		log.Fatalln("The time server host is not specified")
	}
	if len(port) == 0 {
		log.Fatalln("The time server port is not specified")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port specification")
	}

	conn, err := net.DialTimeout("tcp", host+":"+port, 5*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	var i int32

	buf := bufio.NewReader(conn)
	err = binary.Read(buf, binary.BigEndian, &i)
	if err != nil {
		log.Fatalln(err)
	}

	println(i)
}
