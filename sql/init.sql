DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'recipes_user') THEN
        CREATE ROLE recipes_user LOGIN PASSWORD 'super_safe_password';
    END IF;
END
$$;

DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'recipes_db') THEN
        CREATE DATABASE recipes_db;
    END IF;
END
$$;

\c recipes_db

CREATE SCHEMA IF NOT EXISTS recipes_api AUTHORIZATION recipes_user;

GRANT ALL PRIVILEGES ON DATABASE recipes_db TO recipes_user;

-- Leaving out litre and kilogram, since they'll be converted
CREATE TYPE UNIT AS ENUM (
    'ml',
    'g',
    'tsp',
    'tbsp',
    'whole'
);

CREATE TYPE INGREDIENT AS (
    name TEXT,
    quantity INTEGER,
    unit UNIT,
    preparation TEXT
);

CREATE TABLE IF NOT EXISTS recipes_api.recipes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    ingredients INGREDIENT[] NOT NULL,
    instructions TEXT[] NOT NULL,
    servings INTEGER NOT NULL,
    cook_time_minutes INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

INSERT INTO recipes_api.recipes (name, ingredients, instructions, servings, cook_time_minutes)
VALUES (
    'Toast',
    ARRAY [('Bread', 2, 'whole', 'sliced')::INGREDIENT],
    ARRAY ['Put bread in toaster for 2 minutes', 'Apply butter and toppings'],
    1,
    5
)
