CREATE TABLE features
(
  id             UUID PRIMARY KEY,
  name           VARCHAR(255)   NOT NULL,
  code           VARCHAR(255)   NOT NULL,
  created_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);


ALTER TABLE features
ADD CONSTRAINT features_name_unique UNIQUE (code);
