package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kadragon/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(wr http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	data := homeData{"Home", blockchain.GetBlockchain().Allblocks()}
	tmpl.Execute(wr, data)
}

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
