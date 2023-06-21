package transaction

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

type TransactionImpl struct {
	db *sql.DB
}

func NewTransactionImpl(db *sql.DB) *TransactionImpl {
	return &TransactionImpl{
		db: db,
	}
}

func (tsn *TransactionImpl) StartTx(ctx context.Context, opt *sql.TxOptions) (*sql.Tx, error) {
	tx, err := tsn.db.Begin()
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (tsn *TransactionImpl) RunTransaction(ctx context.Context, opt *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := tsn.StartTx(ctx, opt)
	if err != nil {
		log.Err(err).Msg("cannot start tx")
		return err
	}

	err = fn(tx)
	if err != nil {
		log.Err(err).Msg("error")
		if err := tx.Rollback(); err != nil {
			log.Err(err).Msg("error in rollback")
			// return err
		}
		return err
	}

	return tx.Commit()
}
