package types

import "time"

type BalanceHistoryUnit struct {
	Date   time.Time
	Amount int64
}
