package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	loadConf()
	geoipStart()
	setupPihole()

	listen, err := net.ListenPacket("udp", conf.PublicIP+":53")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	defer loggerEnd()
	loggerStart()

	log.Println("Listening on", listen.LocalAddr().String())

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
	dns, err := net.Dial("udp", "127.0.0.1:53")
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
