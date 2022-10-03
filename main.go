package main

import (
	"github.com/BigNutJaa/countryCode"
	"log"
)

func main() {
	//country code (TH,HK,JP,CN,US)
	result, err := countryCode.Country("jP")
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(result)
}
