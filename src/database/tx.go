package database

import (
	"context"
	"fmt"

	"aiisx.com/src/ent"
	"github.com/apex/log"
)

func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}

	return err
}

type TxFn func(ctx context.Context, logger log.Interface, db *ent.Tx) error

func RunWithTx(ctx context.Context, logger log.Interface, db *ent.Client, fn TxFn) error {
	tx, err := db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	err = fn(ctx, logger, tx)

	if err != nil {
		return Rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
