package main

import (
	"flag"
	"log"

	"github.com/luanphandinh/blockchain/blockchain"
)

func main() {
	flag.Parse()
	if *debug {
		blockchain.SetTracer(&simpleTracer{})
	}

	db, err := newBaderDbStorage("./tmp/db", []byte("last_hash"))
	if err != nil {
		log.Fatal(err)
	}

	chain, err := blockchain.InitBlockChain(db, nil)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &CommandLine{chain}
	cmd.Run()
}
