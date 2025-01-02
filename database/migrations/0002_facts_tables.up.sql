BEGIN;

CREATE TABLE institution (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE occasion (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    institution_id INT NOT NULL REFERENCES institution(id) ON DELETE CASCADE,
    occasion_id INT NOT NULL REFERENCES occasion(id) ON DELETE CASCADE,
    time DATE NOT NULL
);

CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL
);

COMMIT;