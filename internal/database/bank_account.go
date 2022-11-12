package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/dzendos/avito-challenge/internal/types"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type bankAccountDB struct {
	db *sql.DB
}

func NewBankAccountDB(db *sql.DB) *bankAccountDB {
	return &bankAccountDB{
		db: db,
	}
}

func (db *bankAccountDB) GetAmount(ctx context.Context, userID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"GetAmount",
	)
	defer span.Finish()

	const query = `
		SELECT 
			SUM(amount)
		FROM bank_account_refill
		WHERE
			s_user_id = $1
		GROUP BY 
			s_user_id
	`

	var amount int64
	err := db.db.QueryRowContext(ctx, query,
		userID,
	).Scan(&amount)

	if err != nil {
		return 0, errors.Wrap(err, "cannot Scan amount")
	}

	return amount, nil
}

func (db *bankAccountDB) AddRefill(ctx context.Context, userID int64, amount int64) error {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"AddRefill",
	)
	defer span.Finish()

	const query = `
		INSERT INTO bank_account_refill(
			s_user_id,
			refill_date,
			amount
		) values (
			$1, $2, $3
		)
	`

	_, err := db.db.ExecContext(ctx, query,
		userID,
		time.Now(),
		amount,
	)

	if err != nil {
		return errors.Wrap(err, "cannot ExecContent")
	}

	return nil
}

func (db *bankAccountDB) GetBalanceHistory(ctx context.Context, userID int64) ([]types.BalanceHistoryUnit, error) {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"GetBalanceHistory",
	)
	defer span.Finish()

	const query = `
		SELECT
			refill_date,
			amount
		FROM 
			bank_account_refill	
		WHERE
			s_user_id = $1
	`

	rows, err := db.db.QueryContext(ctx, query,
		userID,
	)

	if err != nil {
		return nil, errors.Wrap(err, "cannot QueryContext")
	}
	defer rows.Close()

	history := make([]types.BalanceHistoryUnit, 0)
	for rows.Next() {
		var date time.Time
		var amount int64

		if err := rows.Scan(&date, &amount); err != nil {
			return nil, errors.Wrap(err, "cannot Scan")
		}

		history = append(history, types.BalanceHistoryUnit{
			Date:   date,
			Amount: amount,
		})
	}

	return history, nil
}
