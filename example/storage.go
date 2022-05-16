package main

import (
	"context"

	"github.com/dgraph-io/badger/v3"
	"github.com/luanphandinh/blockchain/blockchain"
)

type BadgerDbStorage struct {
	db          *badger.DB
	lastHashKey []byte
}

func newBaderDbStorage(path string, lastHashKey []byte, debug bool) (*BadgerDbStorage, error) {
	opts := badger.DefaultOptions(path)
	if !debug {
		opts.Logger = nil
	}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &BadgerDbStorage{
		db:          db,
		lastHashKey: lastHashKey,
	}, nil
}

func (s *BadgerDbStorage) GetLastBlock(ctx context.Context) (blockchain.Block, error) {
	return s.GetBlock(ctx, s.lastHashKey)
}

func (s *BadgerDbStorage) AddBlock(ctx context.Context, b blockchain.Block) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		bytes, err := b.Marshal()
		if err != nil {
			return err
		}

		err = txn.Set(b.GetHash(), bytes)
		err = txn.Set(s.lastHashKey, bytes)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *BadgerDbStorage) GetBlock(ctx context.Context, key []byte) (blockchain.Block, error) {
	block := &blockchain.SimpleBlock{}
	err := s.db.View(func(txn *badger.Txn) error {
		encodedBlock, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil
			}
			return err
		}

		err = encodedBlock.Value(func(val []byte) error {
			err := block.Unmarshal(val)
			return err
		})

		return err
	})

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *BadgerDbStorage) Close() error {
	return s.db.Close()
}
