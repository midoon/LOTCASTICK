CREATE TYPE simulation_status AS ENUM (
    'ACTIVE',
    'PASSED',
    'FAILED',
    'PAUSED'
);

CREATE TYPE drawdown_type_enum AS ENUM (
    'STATIC',
    'DAILY',
    'TRAILING',
    'COMBINED'
);

CREATE TYPE trade_direction AS ENUM (
    'LONG',
    'SHORT'
);

CREATE TYPE trade_session AS ENUM (
    'ASIA',
    'LONDON',
    'NEW_YORK',
    'OVERLAP'
);