BEGIN;

CREATE TABLE record (
    id UUID PRIMARY KEY,
    session_id UUID NOT NULL REFERENCES session(id) ON DELETE CASCADE,
    speaker_id UUID NOT NULL REFERENCES person(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    events JSONB,
    sequence_number INT NOT NULL
);

CREATE TABLE event (
    id UUID PRIMARY KEY,
    link TEXT NOT NULL,
    record_id UUID NOT NULL REFERENCES record(id) ON DELETE CASCADE
);

COMMIT;