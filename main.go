package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	port int
	dir  string
)

func appMain() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		return
	}
}

func main() {
	app := cli.NewApp()
	app.Usage = "Easy launch HTTP Server"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port",
			Usage:       "Set port number",
			Value:       8080,
			Destination: &port,
		},
	}
	app.Action = func(c *cli.Context) {
		dir = c.Args().First()
		if strings.Compare(dir, "") == 0 {
			dir = "."
		}
		appMain()
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
