

CREATE TABLE users
(
  id             UUID PRIMARY KEY,
  username       TEXT NOT NULL,
  email          TEXT NOT NULL,
  password       TEXT NOT NULL,
  activated_at   TIMESTAMPTZ,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),


  CONSTRAINT users_email_unique
    UNIQUE ( email)
);
