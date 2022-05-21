package main

import (
	"context"
	"fmt"
	"log"

	"github.com/luanphandinh/blockchain/blockchain"
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
	c.chain.AddBlock(ctx, []byte(data))
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

		fmt.Printf("PrevHash: %x\n", block.GetPrevHash())
		fmt.Printf("Data: %s\n", block.GetData())
		fmt.Printf("Hash: %x\n", block.GetHash())

		validated, err := c.chain.ValidateBlock(ctx, block)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("validated: %v\n", validated)
	}
}
