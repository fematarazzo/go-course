package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search IPs and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IPs from internet addresses",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "duckduckgo.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
