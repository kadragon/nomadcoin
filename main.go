package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kadragon/nomadcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(wr http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().Allblocks()}
	templates.ExecuteTemplate(wr, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
