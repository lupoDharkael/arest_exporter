// Copyright 2017 Alejandro Sirgo Rica
//
// This file is part of aREST_exporter.
//
//     aREST_exporter is free software: you can redistribute it and/or modify
//     it under the terms of the GNU General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     aREST_exporter is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU General Public License
//     along with aREST_exporter.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/csv"
	"flag"
	"log"
	"net"
	"os"
	"strings"
)

var (
	listenAddress = flag.String("listen-address", ":9009",
		"The address to listen on for HTTP requests.")
	configFile = flag.String("config.file", "",
		"Sets the configuration file.")
	targets = flag.String("config.targets", "", "Sets the scraping targets.")
)

type Variable struct {
	Name  string
	Value string
}

type Query struct {
	Variables Variable `json:"variables"`
	Connected bool     `json:"connected"`
	Hardware  string   `json:"hardware"`
}

var targetsList []string

func ScrapeIP(ip string) {

}

func main() {
	flag.Parse()
	// get the targets
	switch {
	case *configFile == "" && *targets == "":
		log.Fatalln("No targets found")
	case *configFile != "":
		f, err := os.Open(*configFile)
		if err != nil {
			log.Fatalln(err)
		}
		targetsList, err = csv.NewReader(f).Read()
		if err != nil {
			log.Fatalln(err)
		}
	case *targets != "":
		var err error
		targetsList, err = csv.NewReader(strings.NewReader(*targets)).Read()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// check ips
	for _, ip := range targetsList {
		IPv4 := net.ParseIP(ip).To4()
		if IPv4 == nil {
			log.Fatalf("%s is an invalid IP address", ip)
		}
	}
	// start scraping
	for _, ip := range targetsList {
		go ScrapeIP(ip)
	}
	// infinite wait
	select {}
}
