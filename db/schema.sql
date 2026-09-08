CREATE TABLE inbox
(
    id                    BIGSERIAL PRIMARY KEY,
    event_id              UUID UNIQUE NOT NULL,
    card_id               UUID        NOT NULL,
    chain_id              UUID UNIQUE NOT NULL,
    amount                NUMERIC     NOT NULL,
    currency              TEXT        NOT NULL,
    status                TEXT        NOT NULL CHECK (status IN ('APPROVED', 'REJECTED', 'IN_PROGRESS')),
    transaction_date_time TIMESTAMPTZ NOT NULL,
    done_at               TIMESTAMPTZ,
    created_at            TIMESTAMPTZ,
    inbox_status          TEXT        NOT NULL CHECK (inbox_status IN ('NEW', 'DONE', 'ERROR'))
);

CREATE TABLE card_operations
(
    id                    BIGSERIAL PRIMARY KEY,
    event_id              UUID        NOT NULL,
    card_id               UUID        NOT NULL,
    chain_id              UUID        NOT NULL,
    amount                NUMERIC     NOT NULL,
    currency              TEXT        NOT NULL,
    status                TEXT        NOT NULL CHECK (status IN ('APPROVED', 'REJECTED', 'IN_PROGRESS')),
    transaction_date_time TIMESTAMPTZ NOT NULL
);
