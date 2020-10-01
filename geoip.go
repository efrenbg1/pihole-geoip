package main

import (
	"log"

	"github.com/abh/geoip"
)

var (
	geo, errGeo = geoip.Open("GeoIP/GeoIP.dat", "GeoIP/GeoIPCity.dat")
	areas       = make(map[string]bool)
)

func geoipStart() {
	if errGeo != nil {
		log.Fatal(errGeo)
	}
	for _, i := range conf.Areas {
		areas[i] = true
	}
}

func geoipCheck(ip string) bool {
	gps := geo.GetRecord(ip)
	if gps == nil {
		return areas["nil"]
	}
	code := gps.CountryCode + ":" + gps.PostalCode
	return areas[code]
}
