CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL
        CHECK (char_length(btrim(full_name)) BETWEEN 3 AND 100),
    phone_number VARCHAR(15)
        CHECK (phone_number ~ '^[+][0-9]{9,14}$')
);

CREATE TABLE todoapp.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    user_id INTEGER NOT NULL REFERENCES todoapp.users(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL
        CHECK (char_length(btrim(title)) BETWEEN 1 AND 100),
    description VARCHAR(1000)
        CHECK (char_length(btrim(description)) BETWEEN 1 AND 1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMPTZ,

    CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    )
);
