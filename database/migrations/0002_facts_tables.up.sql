BEGIN;

CREATE TABLE institution (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE occasion (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    institution_id UUID NOT NULL REFERENCES institution(id) ON DELETE CASCADE,
    occasion_id UUID NOT NULL REFERENCES occasion(id) ON DELETE CASCADE,
    time DATE NOT NULL
);

CREATE TABLE persons (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL
);

COMMIT;