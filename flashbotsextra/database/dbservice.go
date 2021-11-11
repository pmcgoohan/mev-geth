package database

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DatabaseService struct {
	DB *sqlx.DB
}

func NewDatabaseService(postgresDSN string) (*DatabaseService, error) {
	db, err := sqlx.Connect("postgres", postgresDSN)
	if err != nil {
		return nil, err
	}
	return &DatabaseService{
		DB: db,
	}, nil
}

func (s *DatabaseService) AddMegabundle(blockNumber int64, parentHash string, blockProfit *big.Rat, bundleProfit *big.Rat, bundleLength int, processingTimeMicros int64, bundle []*types.Transaction, revertingTxHashes []common.Hash) error {
	query := "INSERT INTO megabundles (block_number, parent_hash, block_profit_eth, bundle_profit_eth, bundle_length, processing_time_micros, bundle_txs, tx_hashes, reverting_tx_hashes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	bundlesRaw := []string{}
	txHashes := []string{}
	for _, tx := range bundle {
		b, err := tx.MarshalBinary()
		if err != nil {
			return err
		}
		bundlesRaw = append(bundlesRaw, "0x"+hex.EncodeToString(b))
		txHashes = append(txHashes, tx.Hash().Hex())
	}

	revertingTxHashesHex := []string{}
	for _, hash := range revertingTxHashes {
		revertingTxHashesHex = append(revertingTxHashesHex, hash.Hex())
	}

	_, err := s.DB.Exec(query, blockNumber, parentHash, blockProfit.FloatString(18), bundleProfit.FloatString(18), bundleLength, processingTimeMicros, pq.StringArray(bundlesRaw), pq.StringArray(txHashes), pq.StringArray(revertingTxHashesHex))
	return err
}
