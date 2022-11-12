-- +goose Up
-- +goose StatementBegin

CREATE TABLE operations
(
    s_user_id     BIGINT,
    service_id    BIGINT,
    order_id      BIGINT,
    op_date       DATE,
    amount        BIGINT,
    code          INTEGER
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE operations;

-- +goose StatementEnd
