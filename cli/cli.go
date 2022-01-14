package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/kadragon/nomadcoin/explorer"
	"github.com/kadragon/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to NomadCoin\n\n")
	fmt.Printf("Please use the follwing flags:\n\n")
	fmt.Printf("-port:	Set the Port of the server\n")
	fmt.Printf("-mode:	Choose between 'html' or 'rest' or 'all'\n")

	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of server")
	mode := flag.String("mode", "rest", "Choose between 'html' or 'rest' or 'all'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		usage()
	}
}
