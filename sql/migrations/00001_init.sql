-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE USER_STATUS AS ENUM('blocked', 'unblocked');

CREATE TABLE users (
    id UUID default gen_random_uuid() NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) NOT NULl,
    status USER_STATUS default 'unblocked' NOT NULL,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now(),

    CONSTRAINT pk_users_id PRIMARY KEY (id),
    CONSTRAINT un_users_email UNIQUE(email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE USER_STATUS;
-- +goose StatementEnd
