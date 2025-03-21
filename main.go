package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
}

type Blockchain struct {
	Blocks []*Block
}

func main() {
	run()
}

func (block *Block) culcalateHash() {
	time := []byte(strconv.FormatInt(block.Timestamp, 10))
	heders := bytes.Join([][]byte{block.PrevHash, block.Data, time}, []byte{})
	h := sha256.Sum256(heders)

	block.Hash = h[:]
}

func generateBlock(oldBlock *Block, data []byte) (*Block, error) {
	newBlock := &Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().Unix(),
		Data:      data,
		Hash:      []byte{},
		PrevHash:  oldBlock.PrevHash,
	}

	return newBlock, nil
}

func NewBlock(data []byte) error {

	return nil
}

func run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("main page")
	})
	mux.HandleFunc("/block/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("post to create block")
		var body []byte
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		err = NewBlock(body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	})
	mux.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("all blocks page")
	})
	mux.HandleFunc("/block/last", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("last blocks page")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
