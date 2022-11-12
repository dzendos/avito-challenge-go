package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type operationsDB struct {
	db *sql.DB
}

func NewOperationsDB(db *sql.DB) *operationsDB {
	return &operationsDB{
		db: db,
	}
}

func (db *operationsDB) AddOperation(ctx context.Context, userID, serviceID, orderID, amount int64) error {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"AddOperation",
	)
	defer span.Finish()

	const query = `
		INSERT INTO operations (
			s_user_id,     
			service_id,    
			order_id,      
			op_date,     
			amount,        
			code          
		) values (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := db.db.ExecContext(ctx, query,
		userID,
		serviceID,
		orderID,
		time.Now(),
		amount,
		0,
	)

	if err != nil {
		return errors.Wrap(err, "cannot ExecContent")
	}

	return nil
}

// code = -1 - operation was cancelled
// code =  0 - operation has been created
// code =  1 - operation was approved
func (db *operationsDB) ModifyOperation(ctx context.Context, userID, serviceID, orderID, amount int64, code int) error {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"ModifyOperation",
	)
	defer span.Finish()

	const query = `
		UPDATE operations
		SET
			code = $1
		WHERE
			s_user_id  = $2 AND     
			service_id = $3 AND    
			order_id   = $4 AND       
			amount     = $5 AND
			code       = 0
	`

	_, err := db.db.ExecContext(ctx, query,
		code,
		userID,
		serviceID,
		orderID,
		amount,
	)

	if err != nil {
		return errors.Wrap(err, "cannot ExecContent")
	}

	return nil
}
