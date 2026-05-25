# Trading Journal + Prop Firm Simulator
## Database Documentation

> **Version:** 1.0.0  
> **Database:** PostgreSQL 15+  
> **Schema:** `public`  
> **Conventions:** UUID primary keys, `snake_case` naming, soft deletes via `deleted_at`

---

## Table of Contents

1. [Entity Relationship Overview](#entity-relationship-overview)
2. [Table: users](#table-users)
3. [Table: refresh_tokens](#table-refresh_tokens)
4. [Table: simulations](#table-simulations)
5. [Table: simulation_rules](#table-simulation_rules)
6. [Table: simulation_violations](#table-simulation_violations)
7. [Table: trades](#table-trades)
8. [Table: trade_images](#table-trade_images)
9. [Table: strategies](#table-strategies)
10. [Table: trade_tags](#table-trade_tags)
11. [Table: daily_statistics](#table-daily_statistics)
12. [Table: sessions (reference)](#table-sessions)
13. [Table: challenge_templates](#table-challenge_templates)
14. [Indexes Summary](#indexes-summary)
15. [Enums](#enums)

---

## Entity Relationship Overview

```
users
  ├── refresh_tokens          (1:N)
  ├── strategies              (1:N)
  └── simulations             (1:N)
        ├── simulation_rules  (1:1)
        ├── simulation_violations (1:N)
        ├── daily_statistics  (1:N)
        └── trades            (1:N)
              ├── trade_images  (1:N)
              └── trade_tags    (N:M via trade_tags join)

challenge_templates            (lookup, system-owned)
sessions                       (lookup, reference enum table)
```

---

## Table: `users`

**Purpose:** Core user account table. Stores authentication credentials and profile preferences. One record per registered user.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `email` | `VARCHAR(255)` | NOT NULL | — | Unique login email |
| `password_hash` | `VARCHAR(255)` | NOT NULL | — | bcrypt hash of password |
| `display_name` | `VARCHAR(100)` | NOT NULL | — | User's display name |
| `timezone` | `VARCHAR(64)` | NOT NULL | `'UTC'` | IANA timezone string (e.g., `America/New_York`) |
| `default_currency` | `CHAR(3)` | NOT NULL | `'USD'` | ISO 4217 currency code |
| `login_attempts` | `SMALLINT` | NOT NULL | `0` | Consecutive failed login counter |
| `locked_until` | `TIMESTAMPTZ` | NULL | — | Account lockout expiry (NULL = not locked) |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Last update timestamp |
| `deleted_at` | `TIMESTAMPTZ` | NULL | — | Soft delete timestamp |

### Constraints

```sql
PRIMARY KEY (id)
UNIQUE (email) WHERE deleted_at IS NULL
CHECK (default_currency ~ '^[A-Z]{3}$')
```

### Indexes

```sql
CREATE UNIQUE INDEX idx_users_email ON users (email) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_deleted_at ON users (deleted_at);
```

---

## Table: `refresh_tokens`

**Purpose:** Stores active refresh tokens for JWT session management. Enables server-side token revocation and rotation.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `user_id` | `UUID` | NOT NULL | — | FK → users.id |
| `token_hash` | `VARCHAR(255)` | NOT NULL | — | SHA-256 hash of the raw refresh token |
| `expires_at` | `TIMESTAMPTZ` | NOT NULL | — | Token expiry timestamp |
| `revoked_at` | `TIMESTAMPTZ` | NULL | — | Set when token is revoked (logout) |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Token issued timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
UNIQUE (token_hash)
```

### Indexes

```sql
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);
CREATE INDEX idx_refresh_tokens_token_hash ON refresh_tokens (token_hash);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens (expires_at);
```

---

## Table: `simulations`

**Purpose:** Represents a single prop firm challenge simulation. Each simulation has its own trade journal, rules, and lifecycle state. A user can own multiple simulations.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `user_id` | `UUID` | NOT NULL | — | FK → users.id |
| `template_id` | `UUID` | NULL | — | FK → challenge_templates.id (if created from template) |
| `name` | `VARCHAR(150)` | NOT NULL | — | User-defined simulation name |
| `account_size` | `NUMERIC(20,8)` | NOT NULL | — | Initial account balance |
| `current_equity` | `NUMERIC(20,8)` | NOT NULL | — | Current account equity (updated on each trade) |
| `currency` | `CHAR(3)` | NOT NULL | `'USD'` | Account currency |
| `status` | `simulation_status` | NOT NULL | `'ACTIVE'` | Lifecycle state enum |
| `started_at` | `DATE` | NOT NULL | `CURRENT_DATE` | Challenge start date |
| `passed_at` | `TIMESTAMPTZ` | NULL | — | Timestamp when challenge was passed |
| `failed_at` | `TIMESTAMPTZ` | NULL | — | Timestamp when challenge was failed |
| `notes` | `TEXT` | NULL | — | Optional user notes about simulation |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Last update timestamp |
| `deleted_at` | `TIMESTAMPTZ` | NULL | — | Soft delete timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
FOREIGN KEY (template_id) REFERENCES challenge_templates(id) ON DELETE SET NULL
CHECK (account_size > 0)
CHECK (current_equity >= 0)
CHECK (currency ~ '^[A-Z]{3}$')
```

### Indexes

```sql
CREATE INDEX idx_simulations_user_id ON simulations (user_id);
CREATE INDEX idx_simulations_status ON simulations (status);
CREATE INDEX idx_simulations_deleted_at ON simulations (deleted_at);
CREATE INDEX idx_simulations_user_status ON simulations (user_id, status) WHERE deleted_at IS NULL;
```

---

## Table: `simulation_rules`

**Purpose:** Stores the challenge rules for a simulation. One-to-one with simulations. Rules are immutable once the first trade is logged (enforced at application layer).

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `simulation_id` | `UUID` | NOT NULL | — | FK → simulations.id |
| `drawdown_type` | `drawdown_type_enum` | NOT NULL | — | Type: `STATIC`, `DAILY`, `TRAILING`, `COMBINED` |
| `max_drawdown_pct` | `NUMERIC(6,4)` | NULL | — | Max drawdown % of initial balance (e.g., 10.0000 = 10%) |
| `max_drawdown_amount` | `NUMERIC(20,8)` | NULL | — | Absolute max drawdown amount (derived from pct at save time) |
| `daily_drawdown_pct` | `NUMERIC(6,4)` | NULL | — | Daily drawdown % limit |
| `daily_drawdown_amount` | `NUMERIC(20,8)` | NULL | — | Absolute daily drawdown amount (derived at save time) |
| `trailing_drawdown_pct` | `NUMERIC(6,4)` | NULL | — | Trailing drawdown % |
| `trailing_drawdown_amount` | `NUMERIC(20,8)` | NULL | — | Absolute trailing drawdown |
| `trailing_high_water_mark` | `NUMERIC(20,8)` | NULL | — | Current high-water mark for trailing DD (updated on trade) |
| `profit_target_pct` | `NUMERIC(6,4)` | NOT NULL | — | Profit target as % of initial balance |
| `profit_target_amount` | `NUMERIC(20,8)` | NOT NULL | — | Absolute profit target (derived at save time) |
| `min_trading_days` | `SMALLINT` | NOT NULL | `0` | Minimum trading days required (0 = no minimum) |
| `consistency_rule_enabled` | `BOOLEAN` | NOT NULL | `FALSE` | Whether consistency rule is active |
| `consistency_threshold_pct` | `NUMERIC(6,4)` | NULL | — | Max % of total profit from single day (e.g., 30.0000) |
| `daily_reset_timezone` | `VARCHAR(64)` | NOT NULL | `'UTC'` | Timezone for daily drawdown/stats reset |
| `daily_reset_time` | `TIME` | NOT NULL | `'00:00:00'` | Time of day for daily reset |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Last update timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (simulation_id) REFERENCES simulations(id) ON DELETE CASCADE
UNIQUE (simulation_id)
CHECK (profit_target_pct > 0)
CHECK (min_trading_days >= 0)
CHECK (
  CASE WHEN consistency_rule_enabled = TRUE
    THEN consistency_threshold_pct IS NOT NULL AND consistency_threshold_pct > 0
    ELSE TRUE
  END
)
```

### Indexes

```sql
CREATE UNIQUE INDEX idx_sim_rules_simulation_id ON simulation_rules (simulation_id);
```

---

## Table: `simulation_violations`

**Purpose:** Audit log of every rule violation that has occurred in a simulation. Immutable once written. Used for violation history display and debugging.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `simulation_id` | `UUID` | NOT NULL | — | FK → simulations.id |
| `trade_id` | `UUID` | NULL | — | FK → trades.id (trade that triggered violation) |
| `violation_type` | `VARCHAR(50)` | NOT NULL | — | e.g., `MAX_DRAWDOWN`, `DAILY_DRAWDOWN`, `TRAILING_DRAWDOWN` |
| `equity_at_violation` | `NUMERIC(20,8)` | NOT NULL | — | Account equity at time of violation |
| `rule_floor_at_violation` | `NUMERIC(20,8)` | NOT NULL | — | The drawdown floor that was breached |
| `occurred_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Timestamp of violation |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (simulation_id) REFERENCES simulations(id) ON DELETE CASCADE
FOREIGN KEY (trade_id) REFERENCES trades(id) ON DELETE SET NULL
```

### Indexes

```sql
CREATE INDEX idx_violations_simulation_id ON simulation_violations (simulation_id);
CREATE INDEX idx_violations_occurred_at ON simulation_violations (occurred_at);
```

---

## Table: `trades`

**Purpose:** Core trade journal entries. Each row represents one completed (closed) trade within a simulation. Financial values stored as NUMERIC to avoid floating-point drift.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `simulation_id` | `UUID` | NOT NULL | — | FK → simulations.id |
| `user_id` | `UUID` | NOT NULL | — | FK → users.id (denormalized for query perf) |
| `strategy_id` | `UUID` | NULL | — | FK → strategies.id |
| `symbol` | `VARCHAR(20)` | NOT NULL | — | Trading instrument (e.g., EURUSD, XAUUSD) |
| `direction` | `trade_direction` | NOT NULL | — | `LONG` or `SHORT` |
| `entry_price` | `NUMERIC(20,8)` | NOT NULL | — | Entry price |
| `exit_price` | `NUMERIC(20,8)` | NOT NULL | — | Exit price |
| `stop_loss` | `NUMERIC(20,8)` | NULL | — | Stop loss price |
| `take_profit` | `NUMERIC(20,8)` | NULL | — | Take profit price |
| `lot_size` | `NUMERIC(12,4)` | NOT NULL | — | Position size in lots |
| `pnl` | `NUMERIC(20,8)` | NOT NULL | — | Realized P&L (calculated field, stored for performance) |
| `pnl_pct` | `NUMERIC(10,6)` | NOT NULL | — | P&L as % of account equity at entry |
| `risk_amount` | `NUMERIC(20,8)` | NULL | — | Dollar risk (entry → SL × lot_size × pip_value) |
| `rr_ratio` | `NUMERIC(10,4)` | NULL | — | Risk:Reward ratio (requires SL and TP) |
| `session` | `trade_session` | NULL | — | Market session: `ASIA`, `LONDON`, `NEW_YORK`, `OVERLAP` |
| `entry_time` | `TIMESTAMPTZ` | NOT NULL | — | Trade open timestamp (UTC) |
| `exit_time` | `TIMESTAMPTZ` | NOT NULL | — | Trade close timestamp (UTC) |
| `trade_date` | `DATE` | NOT NULL | — | Date of exit (derived from exit_time, stored for fast grouping) |
| `notes` | `TEXT` | NULL | — | Free-form trade notes |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Last update timestamp |
| `deleted_at` | `TIMESTAMPTZ` | NULL | — | Soft delete timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (simulation_id) REFERENCES simulations(id) ON DELETE CASCADE
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE SET NULL
CHECK (lot_size > 0)
CHECK (entry_time < exit_time)
CHECK (entry_price > 0)
CHECK (exit_price > 0)
```

### Indexes

```sql
CREATE INDEX idx_trades_simulation_id ON trades (simulation_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_trades_user_id ON trades (user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_trades_trade_date ON trades (simulation_id, trade_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_trades_symbol ON trades (simulation_id, symbol) WHERE deleted_at IS NULL;
CREATE INDEX idx_trades_strategy_id ON trades (strategy_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_trades_exit_time ON trades (simulation_id, exit_time DESC) WHERE deleted_at IS NULL;
```

---

## Table: `trade_images`

**Purpose:** Screenshot attachments for trade journal entries. Stores metadata and file reference path; binary data stored in object storage (S3-compatible).

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `trade_id` | `UUID` | NOT NULL | — | FK → trades.id |
| `user_id` | `UUID` | NOT NULL | — | FK → users.id (denormalized for auth checks) |
| `storage_key` | `VARCHAR(500)` | NOT NULL | — | Object storage path (e.g., `uploads/{user_id}/trades/{trade_id}/{id}.jpg`) |
| `original_filename` | `VARCHAR(255)` | NOT NULL | — | Original uploaded filename |
| `mime_type` | `VARCHAR(50)` | NOT NULL | — | MIME type: `image/jpeg`, `image/png`, `image/webp` |
| `file_size_bytes` | `INTEGER` | NOT NULL | — | File size in bytes |
| `label` | `VARCHAR(100)` | NULL | — | User-defined label (e.g., "Entry Setup", "4H Context") |
| `sort_order` | `SMALLINT` | NOT NULL | `0` | Display order within a trade |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Upload timestamp |
| `deleted_at` | `TIMESTAMPTZ` | NULL | — | Soft delete timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (trade_id) REFERENCES trades(id) ON DELETE CASCADE
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
UNIQUE (storage_key)
CHECK (file_size_bytes > 0 AND file_size_bytes <= 5242880)  -- max 5 MB
CHECK (mime_type IN ('image/jpeg', 'image/png', 'image/webp'))
```

### Indexes

```sql
CREATE INDEX idx_trade_images_trade_id ON trade_images (trade_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_trade_images_storage_key ON trade_images (storage_key);
```

---

## Table: `strategies`

**Purpose:** User-defined trading strategy definitions. Used to categorize and analyze trades by strategy. Scoped to a user, accessible across all simulations.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `user_id` | `UUID` | NOT NULL | — | FK → users.id |
| `name` | `VARCHAR(100)` | NOT NULL | — | Strategy name (e.g., "ICT OB Reversal") |
| `description` | `TEXT` | NULL | — | Strategy description |
| `color_hex` | `CHAR(7)` | NOT NULL | `'#6366F1'` | UI color label (hex format) |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Last update timestamp |
| `deleted_at` | `TIMESTAMPTZ` | NULL | — | Soft delete timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
UNIQUE (user_id, name) WHERE deleted_at IS NULL
CHECK (color_hex ~ '^#[0-9A-Fa-f]{6}$')
```

### Indexes

```sql
CREATE INDEX idx_strategies_user_id ON strategies (user_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_strategies_user_name ON strategies (user_id, name) WHERE deleted_at IS NULL;
```

---

## Table: `trade_tags`

**Purpose:** Many-to-many join table between trades and tags. Tags are stored as plain strings (normalized). No separate tag master table — tags are inferred at query time for autocomplete.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `trade_id` | `UUID` | NOT NULL | — | FK → trades.id |
| `user_id` | `UUID` | NOT NULL | — | FK → users.id (for fast tag autocomplete per user) |
| `tag` | `VARCHAR(50)` | NOT NULL | — | Tag value (lowercase, trimmed) |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (trade_id) REFERENCES trades(id) ON DELETE CASCADE
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
UNIQUE (trade_id, tag)
CHECK (tag = lower(trim(tag)))
CHECK (char_length(tag) >= 1 AND char_length(tag) <= 50)
```

### Indexes

```sql
CREATE UNIQUE INDEX idx_trade_tags_trade_tag ON trade_tags (trade_id, tag);
CREATE INDEX idx_trade_tags_user_tag ON trade_tags (user_id, tag);
```

---

## Table: `daily_statistics`

**Purpose:** Pre-computed daily aggregated statistics per simulation per day. Updated by the Risk Engine on every trade mutation. Serves dashboard and analytics endpoints without real-time aggregation.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `simulation_id` | `UUID` | NOT NULL | — | FK → simulations.id |
| `trade_date` | `DATE` | NOT NULL | — | The trading day |
| `trade_count` | `INTEGER` | NOT NULL | `0` | Number of closed trades on this day |
| `winning_trades` | `INTEGER` | NOT NULL | `0` | Count of profitable trades |
| `losing_trades` | `INTEGER` | NOT NULL | `0` | Count of losing trades |
| `breakeven_trades` | `INTEGER` | NOT NULL | `0` | Count of breakeven trades (pnl = 0) |
| `gross_profit` | `NUMERIC(20,8)` | NOT NULL | `0` | Sum of positive P&L |
| `gross_loss` | `NUMERIC(20,8)` | NOT NULL | `0` | Sum of negative P&L (stored as negative value) |
| `daily_pnl` | `NUMERIC(20,8)` | NOT NULL | `0` | Net P&L for the day (gross_profit + gross_loss) |
| `cumulative_pnl` | `NUMERIC(20,8)` | NOT NULL | `0` | Running total P&L from simulation start |
| `equity_open` | `NUMERIC(20,8)` | NOT NULL | — | Account equity at start of trading day |
| `equity_close` | `NUMERIC(20,8)` | NOT NULL | — | Account equity at end of trading day (last trade) |
| `equity_high` | `NUMERIC(20,8)` | NOT NULL | — | Highest intraday equity |
| `equity_low` | `NUMERIC(20,8)` | NOT NULL | — | Lowest intraday equity |
| `max_drawdown_intraday` | `NUMERIC(20,8)` | NOT NULL | `0` | Largest intraday drawdown from high |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Last recalculation timestamp |

### Constraints

```sql
PRIMARY KEY (id)
FOREIGN KEY (simulation_id) REFERENCES simulations(id) ON DELETE CASCADE
UNIQUE (simulation_id, trade_date)
CHECK (trade_count >= 0)
CHECK (winning_trades + losing_trades + breakeven_trades = trade_count)
```

### Indexes

```sql
CREATE UNIQUE INDEX idx_daily_stats_sim_date ON daily_statistics (simulation_id, trade_date);
CREATE INDEX idx_daily_stats_simulation_id ON daily_statistics (simulation_id);
CREATE INDEX idx_daily_stats_trade_date ON daily_statistics (trade_date);
```

---

## Table: `sessions`

**Purpose:** Reference/lookup table for market session definitions. Used to enrich trade session labels and for analytics grouping. System-owned data, not user-editable.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `name` | `VARCHAR(30)` | NOT NULL | — | Session name: `ASIA`, `LONDON`, `NEW_YORK`, `OVERLAP` |
| `open_time_utc` | `TIME` | NOT NULL | — | Session open in UTC |
| `close_time_utc` | `TIME` | NOT NULL | — | Session close in UTC |
| `description` | `VARCHAR(200)` | NULL | — | Human-readable description |

### Constraints

```sql
PRIMARY KEY (id)
UNIQUE (name)
```

### Seed Data

```sql
INSERT INTO sessions (name, open_time_utc, close_time_utc, description) VALUES
  ('ASIA',     '00:00:00', '09:00:00', 'Tokyo/Sydney session'),
  ('LONDON',   '08:00:00', '17:00:00', 'London/Frankfurt session'),
  ('NEW_YORK', '13:00:00', '22:00:00', 'New York session'),
  ('OVERLAP',  '13:00:00', '17:00:00', 'London/New York overlap — highest liquidity');
```

---

## Table: `challenge_templates`

**Purpose:** System-defined challenge templates based on real prop firms. Read-only for users. Creating a simulation from a template copies rule values into simulation_rules.

### Columns

| Column | Type | Nullable | Default | Description |
|---|---|---|---|---|
| `id` | `UUID` | NOT NULL | `gen_random_uuid()` | Primary key |
| `name` | `VARCHAR(150)` | NOT NULL | — | Template display name (e.g., "FTMO $100K Standard") |
| `firm_name` | `VARCHAR(100)` | NOT NULL | — | Prop firm name |
| `account_size` | `NUMERIC(20,8)` | NOT NULL | — | Reference account size |
| `currency` | `CHAR(3)` | NOT NULL | `'USD'` | Account currency |
| `drawdown_type` | `drawdown_type_enum` | NOT NULL | — | Drawdown rule type |
| `max_drawdown_pct` | `NUMERIC(6,4)` | NULL | — | Max drawdown percentage |
| `daily_drawdown_pct` | `NUMERIC(6,4)` | NULL | — | Daily drawdown percentage |
| `trailing_drawdown_pct` | `NUMERIC(6,4)` | NULL | — | Trailing drawdown percentage |
| `profit_target_pct` | `NUMERIC(6,4)` | NOT NULL | — | Profit target percentage |
| `min_trading_days` | `SMALLINT` | NOT NULL | `0` | Minimum trading days |
| `consistency_rule_enabled` | `BOOLEAN` | NOT NULL | `FALSE` | Consistency rule toggle |
| `consistency_threshold_pct` | `NUMERIC(6,4)` | NULL | — | Consistency rule threshold |
| `is_active` | `BOOLEAN` | NOT NULL | `TRUE` | Whether template is visible to users |
| `created_at` | `TIMESTAMPTZ` | NOT NULL | `NOW()` | Record creation timestamp |

### Constraints

```sql
PRIMARY KEY (id)
UNIQUE (name)
CHECK (profit_target_pct > 0)
CHECK (account_size > 0)
```

---

## Indexes Summary

| Table | Index | Type | Purpose |
|---|---|---|---|
| users | `idx_users_email` | UNIQUE PARTIAL | Fast login lookup |
| refresh_tokens | `idx_refresh_tokens_token_hash` | UNIQUE | Token validation |
| simulations | `idx_simulations_user_status` | COMPOSITE PARTIAL | User's sim list filtered by status |
| simulation_rules | `idx_sim_rules_simulation_id` | UNIQUE | 1:1 lookup |
| trades | `idx_trades_simulation_id` | PARTIAL | Trade list per simulation |
| trades | `idx_trades_trade_date` | COMPOSITE PARTIAL | Daily stats calculation |
| trades | `idx_trades_exit_time` | COMPOSITE PARTIAL | Chronological trade queries |
| daily_statistics | `idx_daily_stats_sim_date` | UNIQUE COMPOSITE | Daily stats upsert + lookup |
| trade_tags | `idx_trade_tags_user_tag` | COMPOSITE | Tag autocomplete per user |

---

## Enums

```sql
-- Simulation lifecycle status
CREATE TYPE simulation_status AS ENUM ('ACTIVE', 'PASSED', 'FAILED', 'PAUSED');

-- Drawdown rule type
CREATE TYPE drawdown_type_enum AS ENUM ('STATIC', 'DAILY', 'TRAILING', 'COMBINED');

-- Trade direction
CREATE TYPE trade_direction AS ENUM ('LONG', 'SHORT');

-- Market session
CREATE TYPE trade_session AS ENUM ('ASIA', 'LONDON', 'NEW_YORK', 'OVERLAP');
```

---

## Design Notes

### Financial Precision

All monetary and price values use `NUMERIC(20,8)`:
- 20 total digits supports account sizes up to $999,999,999,999 with 8 decimal places.
- Avoids IEEE 754 floating-point drift in P&L calculations.
- **Never use** `FLOAT` or `DOUBLE PRECISION` for financial data.

### Soft Deletes

All user-owned tables include `deleted_at TIMESTAMPTZ`. All application queries must filter `WHERE deleted_at IS NULL`. Partial indexes with this filter are used on all high-frequency query paths.

### Denormalization

`trades.user_id` is denormalized (also available via `simulations.user_id`) to avoid a join on row-level authorization checks.

### Risk Engine Write Path

When a trade is inserted/updated/deleted:

1. Recalculate equity curve within simulation (transaction).
2. Upsert `daily_statistics` for affected date(s) (`ON CONFLICT (simulation_id, trade_date) DO UPDATE`).
3. Update `simulations.current_equity`.
4. Evaluate and update `simulation_rules.trailing_high_water_mark`.
5. Check for violations → insert into `simulation_violations`, update `simulations.status` if needed.
6. Check for pass conditions → update `simulations.status` and `simulations.passed_at` if met.

All steps execute within a single database transaction.
