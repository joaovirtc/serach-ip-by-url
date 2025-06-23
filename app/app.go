package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate creates a new CLI application.
func Generate() *cli.App {

	app := cli.NewApp()

	app.Name = "Comand Line App"
	app.Usage = "A simple command line application"
	app.Version = "1.0.0"

	flags := []cli.Flag {
				cli.StringFlag {
					Name : "host",
					Value: "Google.com", // Default host to lookup
				},	
	}

	app.Commands = []cli.Command {
		{
			Name : "ip",
			Usage : "Get the IP address of the machine",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name : "server",
			Usage: "Search server ",
			Flags: flags,
			Action: searchServer ,
		},

	}
	
	return app

}



func searchIps ( c * cli.Context) {

	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Printf("Public IP address of %s: %s\n", host, ip.String())
	}

}

func searchServer ( c* cli.Context) {

	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Printf("Server for %s: %s\n", host, server.Host)
	}
}

