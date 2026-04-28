CREATE TABLE plan_features
(
  id       UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
  plan_id    UUID    NOT NULL,
  feature_id UUID    NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),


  FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE CASCADE,
  FOREIGN KEY (feature_id) REFERENCES features(id) ON DELETE CASCADE
);


