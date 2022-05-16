package main

import (
	"context"
	"flag"
	"log"

	"github.com/luanphandinh/blockchain/blockchain"
)

func main() {
	ctx := context.Background()
	flag.Parse()
	if *debug {
		blockchain.SetTracer(&simpleTracer{})
	}

	db, err := newBaderDbStorage("./tmp/db", []byte("last_hash"), *debug)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	chain, err := blockchain.InitBlockChain(ctx, db, nil)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &CommandLine{chain}
	cmd.Run(ctx)
}
