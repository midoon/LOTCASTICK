# Trading Journal + Prop Firm Simulator
## Feature Documentation

> **Version:** 1.0.0  
> **Status:** Pre-Development Specification  
> **Last Updated:** 2025  
> **Audience:** Engineering Team, Product Team

---

## Table of Contents

1. [Overview](#overview)
2. [Authentication](#1-authentication)
3. [User Profile & Settings](#2-user-profile--settings)
4. [Simulation Management](#3-simulation-management)
5. [Prop Firm Challenge Rules Engine](#4-prop-firm-challenge-rules-engine)
6. [Trading Journal](#5-trading-journal)
7. [Risk Engine](#6-risk-engine)
8. [Dashboard](#7-dashboard)
9. [Analytics](#8-analytics)
10. [Challenge Templates](#9-challenge-templates)
11. [Non-Functional Requirements](#10-non-functional-requirements)

---

## Overview

**Trading Journal + Prop Firm Simulator** is a SaaS web application that allows retail traders to:

- Simulate prop firm challenges (FTMO, MyFundedFx, The5ers, etc.) using their own trading journal data.
- Track live P&L, drawdown, and challenge status across multiple simulations simultaneously.
- Analyze trading performance with professional-grade analytics.
- Maintain a structured trade journal with tagging, screenshots, and strategy mapping.

### Core User Flow

```
Register / Login
    └─► Create Simulation (choose template or custom rules)
            └─► Log Trades (manual entry or bulk import)
                    └─► Monitor Dashboard (equity curve, drawdown, status)
                            └─► Review Analytics (performance breakdown)
```

---

## 1. Authentication

### 1.1 Registration

- User registers with **email + password**.
- Password requirements: minimum 8 characters, at least 1 uppercase, 1 number.
- Email must be unique in the system.
- Upon successful registration, a **JWT access token** and **refresh token** are returned.
- System creates a default user profile on registration.

### 1.2 Login

- User authenticates with email + password.
- Returns **JWT access token** (short-lived, 1 hour) and **refresh token** (long-lived, 7 days).
- Failed login attempts are tracked; account is soft-locked after 5 consecutive failures (30-minute lockout).

### 1.3 Logout

- Invalidates the active refresh token server-side.
- Client is responsible for clearing the access token from local storage.

### 1.4 Token Refresh

- Client sends refresh token to obtain a new access token.
- Refresh token is rotated on each use (rolling refresh).
- Expired or revoked refresh tokens return `401 Unauthorized`.

### 1.5 Session Management

- All protected endpoints require `Authorization: Bearer <access_token>` header.
- Access token payload contains: `user_id`, `email`, `issued_at`, `expires_at`.
- Stateless JWT validation on every request — no server-side session store required.

---

## 2. User Profile & Settings

### 2.1 Profile Management

- User can update: display name, timezone, default currency (USD, EUR, GBP).
- Timezone setting affects: daily reset calculations, session analytics grouping.

### 2.2 Password Management

- Change password (requires current password confirmation).

---

## 3. Simulation Management

### 3.1 Create Simulation

- User creates a simulation by providing:
  - **Simulation name** (e.g., "FTMO $100K Phase 1")
  - **Account size** (e.g., 100,000)
  - **Currency** (USD, EUR, GBP)
  - **Challenge rules** (from template or fully custom)
  - **Start date** (default: today)
- Each simulation is independent with its own trade journal, equity tracking, and rule state.
- Maximum simulations per user: **20 active simulations**.

### 3.2 List Simulations

- Returns paginated list of user's simulations.
- Each item includes: name, account size, current equity, challenge status, progress summary.
- Filterable by: status (active, passed, failed, paused).
- Sortable by: created date, equity, name.

### 3.3 View Simulation Detail

- Returns full simulation metadata including:
  - All challenge rules (with current state per rule)
  - Account starting balance and current equity
  - Total trading days count
  - Challenge start/end date
  - Violation history (if any)

### 3.4 Update Simulation

- User can update: simulation name, notes.
- Challenge rules **cannot be modified** after the first trade is logged (immutable for integrity).
- Exception: simulation can be **reset** (deletes all trades, resets equity, restarts challenge clock).

### 3.5 Delete Simulation

- Soft delete only — data is retained for 30 days before permanent deletion.
- All associated trades, images, and statistics are cascaded in the soft delete.

### 3.6 Simulation Status Lifecycle

```
ACTIVE ──► PASSED   (profit target reached + all conditions met)
       ──► FAILED   (drawdown violated or other rule breach)
       ──► PAUSED   (manually paused by user)
PAUSED ──► ACTIVE   (manually resumed)
```

---

## 4. Prop Firm Challenge Rules Engine

### 4.1 Drawdown Rules

Each simulation can have **one drawdown type** configured. Types are mutually exclusive.

#### 4.1.1 Maximum Drawdown (Static from Initial Balance)

- **Definition:** Maximum allowed loss measured from the **initial account balance**.
- **Formula:** `current_equity >= initial_balance - max_drawdown_amount`
- **Example:** $100,000 account, 10% max drawdown → account cannot go below $90,000.
- Violation: account equity falls below the static floor.

#### 4.1.2 Daily Drawdown

- **Definition:** Maximum loss allowed within a single trading day.
- **Reset time:** Configurable per simulation (e.g., 00:00 UTC, 17:00 EST).
- **Formula:** `daily_pnl >= -daily_drawdown_amount` (calculated from EOD equity of previous day OR daily open equity — configurable).
- Violation: daily loss exceeds the daily drawdown limit at any point during the day.

#### 4.1.3 Trailing Drawdown (High-Water Mark)

- **Definition:** Maximum drawdown calculated from the **highest equity ever reached**.
- **Formula:** `current_equity >= max_equity_ever_reached - trailing_drawdown_amount`
- **Example:** Account grows to $105,000, then trailing floor becomes $95,000. If account then hits $100,000 (new high), floor moves to $90,000.
- The floor **only moves up**, never down.
- Violation: current equity falls below the trailing floor.

#### 4.1.4 Combined Drawdown

- Simulation can optionally combine **max drawdown + daily drawdown** — both rules must be satisfied simultaneously (common in real prop firms).

### 4.2 Profit Target

- User sets a **profit target percentage or absolute amount**.
- **Formula:** `current_equity >= initial_balance + profit_target_amount`
- Profit target must be reached **while all other rules remain unviolated**.
- Reaching profit target does NOT automatically pass the simulation — minimum trading days must also be satisfied.

### 4.3 Minimum Trading Days

- **Definition:** Minimum number of calendar days on which at least one trade must have been executed.
- **Formula:** `COUNT(DISTINCT trade_date) >= min_trading_days`
- Trading day = any day where at least one closed trade exists.
- Example: 10 minimum trading days required.

### 4.4 Consistency Rule

- **Definition:** No single trading day can account for more than X% of total profit.
- **Formula:** `max_single_day_profit / total_profit <= consistency_threshold`
- Applies only when total profit is positive.
- Example: 30% consistency rule → no single day can represent more than 30% of total realized profit.

### 4.5 Challenge Pass Conditions

All of the following must be true simultaneously:

1. Profit target reached.
2. Minimum trading days completed.
3. No drawdown rule ever violated.
4. Consistency rule satisfied (if enabled).
5. Simulation status is ACTIVE (not PAUSED).

### 4.6 Timezone Reset

- Daily drawdown and daily statistics reset at a configurable time.
- Supported formats: UTC offset (e.g., `UTC-5`, `UTC+0`) or named timezone (e.g., `America/New_York`).
- Default reset time: `00:00` in the configured timezone.

---

## 5. Trading Journal

### 5.1 Add Trade

User logs a closed trade with the following fields:

| Field | Type | Required | Description |
|---|---|---|---|
| symbol | string | Yes | Trading instrument (e.g., EURUSD, XAUUSD, NQ) |
| direction | enum | Yes | LONG or SHORT |
| entry_price | decimal | Yes | Entry price |
| exit_price | decimal | Yes | Exit price |
| lot_size | decimal | Yes | Position size in lots |
| entry_time | datetime | Yes | Trade open timestamp (UTC) |
| exit_time | datetime | Yes | Trade close timestamp (UTC) |
| strategy_id | UUID | No | Link to user-defined strategy |
| session | enum | No | ASIA, LONDON, NEW_YORK, OVERLAP |
| pnl | decimal | Auto | Calculated: (exit - entry) × lot_size × pip_value |
| rr | decimal | Auto | Risk-Reward ratio (requires SL/TP) |
| stop_loss | decimal | No | Stop loss price |
| take_profit | decimal | No | Take profit price |
| notes | text | No | Free-form trade notes |
| tags | array | No | Array of tag strings |

### 5.2 Edit Trade

- All fields are editable.
- Editing a trade triggers **recalculation** of:
  - Daily statistics for affected date(s)
  - Equity curve data points
  - Drawdown calculations
  - Challenge status re-evaluation

### 5.3 Delete Trade

- Soft delete with cascade recalculation (same as edit).
- Trade images are also soft-deleted.

### 5.4 Trade Images (Screenshots)

- User can attach up to **5 images per trade**.
- Supported formats: JPEG, PNG, WEBP.
- Maximum file size: **5 MB per image**.
- Images are stored with a label field (e.g., "Entry", "Exit", "4H Context").
- Images do not block trade creation — they can be uploaded after the trade is logged.

### 5.5 Trade Filtering & Sorting

Trades list supports:

- **Filter by:** date range, symbol, direction, session, strategy, tags, win/loss.
- **Sort by:** entry time, P&L, R:R, lot size.
- **Pagination:** default 50 per page, max 200.
- **Search:** free-text search on notes.

### 5.6 Strategies

- User defines custom strategies (e.g., "ICT Breaker Block", "Supply & Demand", "News Fade").
- Strategy has: name, description, color label.
- Strategies are **user-scoped**, shared across all simulations.
- Each trade can be linked to one strategy.

### 5.7 Tags

- Free-form tags on trades (e.g., "revenge-trade", "FOMO", "high-conviction").
- Tags are auto-suggested based on user's existing tags.
- Tags are stored normalized (lowercase, trimmed).

---

## 6. Risk Engine

The Risk Engine runs automatically whenever a trade is added, edited, or deleted.

### 6.1 Real-Time Drawdown Calculation

On every trade mutation:

1. Recalculate equity curve from first trade to current.
2. Recalculate high-water mark (for trailing drawdown).
3. Recalculate daily P&L for the affected trading day.
4. Evaluate all active drawdown rules against updated values.

### 6.2 Challenge Violation Detection

- If any rule is violated, simulation status immediately transitions to `FAILED`.
- Violation event is recorded: which rule, at what equity, at what timestamp, which trade triggered it.
- Violation is **irreversible** unless the trade causing it is deleted/edited.

### 6.3 Challenge Pass Detection

- After every trade mutation, check if all pass conditions are met.
- If all conditions satisfied, simulation status transitions to `PASSED`.

### 6.4 Remaining Drawdown Calculation

Exposed in API for dashboard display:

- `remaining_max_drawdown`: current_equity - static_floor
- `remaining_daily_drawdown`: daily_drawdown_limit - abs(todays_pnl)
- `remaining_trailing_drawdown`: current_equity - trailing_floor
- All values expressed in both absolute amount and percentage of initial balance.

### 6.5 Account Status Summary

Returned with every dashboard request:

```
{
  status: ACTIVE | PASSED | FAILED | PAUSED,
  equity: 102500.00,
  drawdown_used_pct: 2.5,
  drawdown_remaining_pct: 7.5,
  daily_drawdown_used_pct: 0.8,
  daily_drawdown_remaining_pct: 4.2,
  profit_target_pct: 25.0,
  profit_achieved_pct: 12.5,
  trading_days_completed: 6,
  trading_days_required: 10,
  consistency_score_pct: 22.3,
  violations: []
}
```

---

## 7. Dashboard

### 7.1 Equity Curve Chart

- Line chart: equity value over time (per-trade granularity).
- Overlays: drawdown floor line (static and/or trailing), profit target line.
- Supports zoom: All / 1M / 1W / 1D.

### 7.2 P&L Chart

- Bar chart: daily P&L (green for positive, red for negative).
- Cumulative P&L line overlay.

### 7.3 Calendar Heatmap

- Monthly calendar view where each day is colored by daily P&L.
- Color scale: deep red (large loss) → neutral (breakeven) → deep green (large win).
- Clicking a day filters the trade list to that date.

### 7.4 Challenge Progress Bar

- Visual progress toward: profit target, minimum trading days.
- Real-time drawdown gauge (used vs remaining).

### 7.5 Statistics Summary Cards

| Metric | Description |
|---|---|
| Total Trades | Count of all closed trades |
| Win Rate | Winning trades / Total trades × 100 |
| Profit Factor | Gross profit / Gross loss |
| Expectancy | (Win rate × avg win) - (loss rate × avg loss) |
| Average R:R | Mean risk-reward ratio across all trades |
| Best Day | Highest single-day P&L |
| Worst Day | Lowest single-day P&L |
| Avg Daily P&L | Mean daily P&L across all trading days |

### 7.6 Recent Trades Widget

- Last 10 trades with: symbol, direction, P&L, R:R, timestamp.
- Quick link to full trade list.

---

## 8. Analytics

### 8.1 Win Rate Analysis

- Overall win rate (%).
- Win rate by: symbol, session, strategy, tag, day-of-week, month.
- Win/Loss/Breakeven distribution chart.

### 8.2 Profit Factor

- Overall profit factor.
- Profit factor trend over rolling 20-trade window.

### 8.3 Expectancy

- Mathematical expectancy per trade (in $).
- Expectancy by strategy and symbol.

### 8.4 Performance by Symbol

For each traded symbol:

- Total trades, win rate, total P&L, avg P&L per trade, profit factor, avg R:R.

### 8.5 Performance by Session

Group trades by market session (Asia, London, New York, London/NY Overlap):

- Total trades, win rate, total P&L, best session, worst session.

### 8.6 Performance by Strategy

For each linked strategy:

- Total trades, win rate, total P&L, expectancy, avg R:R.
- Trade list filtered by strategy.

### 8.7 Performance by Day of Week

- Win rate, avg P&L, total trades per day (Monday–Friday).
- Identify best and worst trading days.

### 8.8 Drawdown Analytics

- Maximum drawdown reached (absolute and %).
- Drawdown duration chart (periods spent below high-water mark).
- Recovery factor: net profit / max drawdown.

### 8.9 Streak Analysis

- Current win streak / loss streak.
- Longest win streak / loss streak on record.

---

## 9. Challenge Templates

### 9.1 Built-In Templates

Pre-configured rule sets based on popular prop firms:

| Template Name | Account | Max DD | Daily DD | Profit Target | Min Days |
|---|---|---|---|---|---|
| FTMO Standard | $100,000 | 10% | 5% | 10% | 10 |
| MyFundedFx | $100,000 | 12% trailing | 5% | 10% | 5 |
| The5ers Hyper | $100,000 | 6% static | — | 8% | — |
| TopStep Futures | $150,000 | trailing $3000 | — | $9,000 | 10 |

> Note: Templates are read-only references. Creating a simulation from a template copies the rules into the simulation.

### 9.2 Custom Challenge Rules

User can define fully custom rules:

- Choose any drawdown type (static, daily, trailing, combined).
- Set any profit target (percentage or absolute).
- Set any minimum trading days.
- Enable/disable consistency rule with custom threshold.
- Set custom timezone for daily reset.

---

## 10. Non-Functional Requirements

### 10.1 Performance

- API response time: < 200ms for read endpoints (p95).
- Dashboard data: served from pre-computed `daily_statistics` table (not calculated on-the-fly).
- Trade submission: < 500ms including risk engine recalculation.

### 10.2 Security

- All endpoints protected with JWT Bearer authentication.
- Row-level authorization: users can only access their own simulations and trades.
- Input sanitization and validation on all endpoints.
- Rate limiting: 100 req/min per user on write endpoints.

### 10.3 Data Integrity

- All financial values stored as `NUMERIC(20,8)` — never floating point.
- Trade mutations are wrapped in database transactions.
- Risk engine recalculation is atomic with trade save.

### 10.4 Scalability

- Stateless backend — horizontally scalable.
- Database indexes on all foreign keys and common filter columns.
- Pagination enforced on all list endpoints.
