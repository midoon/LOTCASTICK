CREATE TABLE email_verifications
(
  id           UUID PRIMARY KEY,
  user_id      UUID NOT NULL,
  token        TEXT NOT NULL,
  expired_at   TIMESTAMPTZ NOT NULL,
  used_at      TIMESTAMPTZ,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_email_token_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

  CONSTRAINT unique_email_verification_tokens
    UNIQUE (token)
);

CREATE INDEX idx_email_verification_user_id
ON email_verifications (user_id);

CREATE INDEX idx_email_verification_expired_at
ON email_verifications (expired_at);

CREATE INDEX idx_email_verification_unused
ON email_verifications (used_at)
WHERE used_at IS NULL;
