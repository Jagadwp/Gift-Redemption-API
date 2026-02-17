CREATE TABLE IF NOT EXISTS redemptions (
    id          SERIAL PRIMARY KEY,
    user_id     INT         NOT NULL REFERENCES users(id),
    gift_id     INT         NOT NULL REFERENCES gifts(id),
    quantity    INT         NOT NULL DEFAULT 1,
    total_point INT         NOT NULL,
    redeemed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_redemptions_user_id ON redemptions(user_id);
CREATE INDEX idx_redemptions_gift_id ON redemptions(gift_id);

CREATE TABLE IF NOT EXISTS ratings (
    id             SERIAL PRIMARY KEY,
    user_id        INT         NOT NULL REFERENCES users(id),
    gift_id        INT         NOT NULL REFERENCES gifts(id),
    redemption_id  INT         NOT NULL REFERENCES redemptions(id),
    score          NUMERIC(2,1) NOT NULL CHECK (score BETWEEN 1 AND 5),
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    -- one rating per redemption
    CONSTRAINT uq_rating_redemption UNIQUE (redemption_id)
);

CREATE INDEX idx_ratings_gift_id ON ratings(gift_id);
CREATE INDEX idx_ratings_user_id ON ratings(user_id);
