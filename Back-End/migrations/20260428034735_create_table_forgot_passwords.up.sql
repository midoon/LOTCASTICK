CREATE TABLE forgot_passwords
(
  id           UUID PRIMARY KEY,
  user_id      UUID NOT NULL,
  code         VARCHAR(255) NOT NULL,
  expired_at   TIMESTAMPTZ NOT NULL,
  used_at      TIMESTAMPTZ,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_password_reset_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX uniq_active_password_reset_per_user
ON forgot_passwords (user_id)
WHERE used_at IS NULL;

CREATE INDEX idx_password_reset_expired_at
ON forgot_passwords (expired_at);
