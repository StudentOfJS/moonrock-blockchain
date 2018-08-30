package main

import (
	"encoding/json"
	"io"
	"net"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type Message struct {
	BPM int
}

// HandleWriteBlock writes data to the blockchain
func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.BPM)
	if err != nil {
		RespondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		spew.Dump(Blockchain)
	}

	RespondWithJSON(w, r, http.StatusCreated, newBlock)

}

// HandleGetBlockchain returns the full blockchsin as JSON
func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// HandleConn takes care of the incoming blockchain connection requests
func HandleConn(conn net.Conn) {
	defer conn.Close()
}
