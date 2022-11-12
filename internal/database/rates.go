package database

import (
	"database/sql"
	"time"

	"github.com/dzendos/avito-challenge/internal/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type ratesDB struct {
	db *sql.DB
}

func NewRatesDB(db *sql.DB) *ratesDB {
	return &ratesDB{
		db: db,
	}
}

func (db *ratesDB) SetCurrencyRate(ctx context.Context, currency types.CurrencyRate) error {
	const query = `
		INSERT INTO currency_rate(
			char_code,
			base,
			rate,
			rate_date
		) values (
			$1, $2, $3, $4
		)
		ON CONFLICT(char_code)
		DO UPDATE 
		SET
			rate = $3,
			rate_date = $4
	`
	_, err := db.db.ExecContext(ctx, query,
		currency.CharCode,
		currency.BaseCurrency,
		currency.Rate,
		currency.Date,
	)

	if err != nil {
		return errors.Wrap(err, "cannot ExecContent")
	}

	return nil
}

func (db *ratesDB) GetCurrencyRate(ctx context.Context, currency types.Currency, date time.Time) (int64, error) {
	const query = `
		SELECT 
			rate
		FROM 
			currency_rate
		WHERE
			char_code = $1 AND
			rate_date = $2
	`
	var rate int64
	err := db.db.QueryRowContext(ctx, query,
		currency,
		date,
	).Scan(&rate)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, types.ErrNoCurrencyRate
		} else {
			return 0, errors.Wrap(err, "cannot QueryRowContext")
		}
	}

	return rate, nil
}
