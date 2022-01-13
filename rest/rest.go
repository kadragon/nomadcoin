package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kadragon/nomadcoin/blockchain"
	"github.com/kadragon/nomadcoin/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
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
		var addBlockBody addBlockBody

		err := json.NewDecoder(r.Body).Decode(&addBlockBody)
		utils.HandleErr(err)

		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)

		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
}

func Start(aPort int) {
	// gorilla mux 사용
	handler := mux.NewRouter()

	port = fmt.Sprintf(":%d", aPort)

	handler.HandleFunc("/", documentation).Methods("GET")
	handler.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	handler.HandleFunc("/blocks/{id:[0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
