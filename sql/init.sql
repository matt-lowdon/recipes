\c recipes

CREATE TYPE UNIT AS ENUM (
    'ml',
    'l',
    'g',
    'kg',
    'tsp',
    'tbsp',
    'whole'
);

CREATE TYPE INGREDIENT AS (
    name TEXT,
    quantity INTEGER,
    unit UNIT
);

CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    ingredients INGREDIENT[] NOT NULL
    servings INTEGER NOT NULL,
    cook_time_minutes INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

INSERT INTO recipes VALUES (
    'Toast',
    '(Bread, 2, whole)' :: INGREDIENT
