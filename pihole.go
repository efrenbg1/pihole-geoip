package main

import (
	"log"
	"os/exec"
)

func setupPihole() {
	if len(conf.Dnsmasq) == 0 {
		return
	}
	cmd("sed -i 's/interface=.*/interface=lo/' " + conf.Dnsmasq)
	cmd("sed -i '/bind-interfaces/d' " + conf.Dnsmasq)
	cmd("sed -i '${/^$/d;}' " + conf.Dnsmasq)
	cmd("echo -e \"\nbind-interfaces\" >> " + conf.Dnsmasq)
	cmd("pihole restartdns")
	log.Println("Pi-hole now listening on localhost only")
}

func cmd(command string) ([]byte, error) {
	out, err := exec.Command("bash", "-c", command).Output()
	return out, err
}
