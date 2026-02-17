CREATE TABLE IF NOT EXISTS gifts (
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(200)   NOT NULL,
    description    TEXT,
    point          INT            NOT NULL DEFAULT 0,
    stock          INT            NOT NULL DEFAULT 0,
    image_url      VARCHAR(500),
    is_new         BOOLEAN        NOT NULL DEFAULT FALSE,
    is_best_seller BOOLEAN        NOT NULL DEFAULT FALSE,
    avg_rating     NUMERIC(3, 2)  NOT NULL DEFAULT 0,
    total_reviews  INT            NOT NULL DEFAULT 0,
    created_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMPTZ
);

CREATE INDEX idx_gifts_deleted_at  ON gifts(deleted_at);
CREATE INDEX idx_gifts_avg_rating  ON gifts(avg_rating) WHERE deleted_at IS NULL;
CREATE INDEX idx_gifts_created_at  ON gifts(created_at) WHERE deleted_at IS NULL;