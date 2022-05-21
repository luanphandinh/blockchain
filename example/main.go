package main

import (
	"context"
	"flag"
	"log"

	"github.com/luanphandinh/blockchain/blockchain"
)

var (
	addBlock             = flag.String("add", "", "add new block with data")
	printChain           = flag.Bool("print", false, "print out the chain")
	usePersistantStorage = flag.Bool("persistant_storage", true, "use persistant_storage")
	debug                = flag.Bool("debug", false, "debug")
)

func main() {
	ctx := context.Background()
	flag.Parse()
	if *debug {
		blockchain.SetTracer(&simpleTracer{})
	}

	opts := &blockchain.BlockChainConfigs{}
	if *usePersistantStorage {
		db, err := newBaderDbStorage("./tmp/db", []byte("last_hash"), *debug)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		opts.Storage = db
	}

	chain, err := blockchain.InitBlockChain(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &CommandLine{chain}
	cmd.Run(ctx)
}
