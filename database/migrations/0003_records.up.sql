BEGIN;

CREATE TABLE records (
    id UUID PRIMARY KEY,
    session_id UUID NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    speaker_id UUID NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    events JSONB,
    sequence_number INT NOT NULL
);

CREATE TABLE events (
    id UUID PRIMARY KEY,
    link TEXT,
    record_id UUID REFERENCES records(id) ON DELETE CASCADE
);

COMMIT;