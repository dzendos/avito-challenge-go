-- +goose Up
-- +goose StatementBegin

CREATE TABLE bank_account_refill
(
    s_user_id   BIGINT,
    refill_date DATE,
    amount      BIGINT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE bank_account_refill;

-- +goose StatementEnd
