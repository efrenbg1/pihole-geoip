package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	loadConf()
	geoipStart()
	setupPihole()

	listen, err := net.ListenPacket("udp", conf.Listen)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	defer loggerEnd()
	loggerStart()

	fmt.Println("Listening on", listen.LocalAddr().String(), "--->", conf.Target)

	for {
		buff := make([]byte, 1220)
		n, addr, err := listen.ReadFrom(buff)
		if err != nil {
			continue
		}
		ip := strings.Split(addr.String(), ":")[0]
		if geoipCheck(ip) {
			clients <- ip
			go forward(listen, addr, buff[:n])
		} else {
			blocked <- ip
		}
	}
}

func forward(listen net.PacketConn, addr net.Addr, tx []byte) {
	defer recover()
	dns, err := net.Dial("udp", conf.Target)
	dns.SetDeadline(time.Now().Add(time.Duration(2) * time.Second))
	dns.Write(tx)

	rx := make([]byte, 1220)
	_, err = bufio.NewReader(dns).Read(rx)

	if err != nil {
		log.Println(err)
		return
	}

	listen.WriteTo(rx, addr)
}
