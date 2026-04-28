CREATE TABLE plans
(
  id             UUID PRIMARY KEY,
  name           TEXT           NOT NULL,
  price_monthly  NUMERIC(15,2)  NOT NULL DEFAULT 0,
  price_yearly   NUMERIC(15,2)  NOT NULL DEFAULT 0,
  created_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

ALTER TABLE plans
ADD CONSTRAINT plans_name_unique UNIQUE (name);
