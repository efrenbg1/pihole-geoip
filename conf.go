package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type confStruct struct {
	PublicIP     string   `json:"publicIP"`
	Dnsmasq      string   `json:"dnsmasq"`
	LoggerPeriod int      `json:"loggerPeriod"`
	Areas        []string `json:"areas"`
}

var (
	conf confStruct
)

func loadConf() {
	confFile, err := os.Open("config.json")
	if err != nil {
		log.Panic("Error while openning config file. Are permissions right?")
	}
	defer confFile.Close()
	bytes, _ := ioutil.ReadAll(confFile)
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		log.Panic("Error loading config file!")
	}
}
