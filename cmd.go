package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/luanphandinh/blockchain/blockchain"
)

var (
	addBlock   = flag.String("add", "", "add new block with data")
	printChain = flag.Bool("print", false, "print out the chain")
	debug      = flag.Bool("debug", false, "debug")
)

type CommandLine struct {
	chain *blockchain.BlockChain
}

func (c *CommandLine) Run(ctx context.Context) {
	if *addBlock != "" {
		c.AddBlock(ctx, *addBlock)
	}

	if *printChain {
		c.Print(ctx)
	}

	fmt.Println("Done...")
}

func (c *CommandLine) AddBlock(ctx context.Context, data string) {
	c.chain.AddBlock(ctx, data)
}

func (c *CommandLine) Print(ctx context.Context) {
	iterator := c.chain.NewIterator()
	for {
		block, err := iterator.Next(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if block == nil {
			return
		}

		fmt.Printf("PrevHash: %x\n", block.Prevhash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		validated, err := c.chain.ValidateBlock(ctx, block)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("validated: %v\n", validated)
	}
}
