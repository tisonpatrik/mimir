BEGIN;

CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    session_id INT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    speaker_id INT NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    events JSONB,
    sequence_number INT NOT NULL
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    link TEXT,
    record_id INT REFERENCES records(id) ON DELETE CASCADE
);

COMMIT;