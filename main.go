package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kadragon/nomadcoin/blockchain"
	"github.com/kadragon/nomadcoin/utils"
)

const port string = ":4000"

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}

	rw.Header().Add("Content-Type", "Application/json")

	err := json.NewEncoder(rw).Encode(data)
	utils.HandleErr(err)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "Application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().Allblocks())

	case "POST":
		var addBlockBody AddBlockBody

		err := json.NewDecoder(r.Body).Decode(&addBlockBody)
		utils.HandleErr(err)

		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)

		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
