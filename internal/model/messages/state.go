package messages

import "github.com/dzendos/avito-challenge/internal/types"

var state = make(map[int64]types.Currency)

func SetUserCurrency(userID int64, currency types.Currency) {
	state[userID] = currency
}

func GetUserCurrency(userID int64) types.Currency {
	currency, ok := state[userID]
	if !ok {
		state[userID] = types.RUB
		currency = types.RUB
	}

	return currency
}
