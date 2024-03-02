CREATE TABLE IF NOT EXISTS messages (
    id bigserial NOT NULL PRIMARY KEY,
    content text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
