package main

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/luanphandinh/blockchain/blockchain"
)

type BadgerDbStorage struct {
	db          *badger.DB
	lastHashKey []byte
}

func newBaderDbStorage(path string, lastHashKey []byte) (blockchain.Storage, error) {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &BadgerDbStorage{
		db:          db,
		lastHashKey: lastHashKey,
	}, nil
}

func (s *BadgerDbStorage) GetLastBlock() (*blockchain.Block, error) {
	return s.GetBlock(s.lastHashKey)
}

func (s *BadgerDbStorage) AddBlock(b *blockchain.Block) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		bytes, err := b.Serialize()
		if err != nil {
			return err
		}

		err = txn.Set(s.lastHashKey, bytes)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *BadgerDbStorage) GetBlock(key []byte) (*blockchain.Block, error) {
	var block *blockchain.Block
	err := s.db.View(func(txn *badger.Txn) error {
		encodedBlock, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil
			}
			return err
		}

		err = encodedBlock.Value(func(val []byte) error {
			block, err = blockchain.DeserializeBlock(val)
			return err
		})

		return err
	})

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *BadgerDbStorage) GetBlocks() ([]*blockchain.Block, error) {
	return nil, nil
}
