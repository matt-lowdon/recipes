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
    'whole',
    'bulb',
    'clove'
);

CREATE TABLE IF NOT EXISTS recipes_api.recipes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    servings INTEGER NOT NULL,
    cook_time_minutes INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS recipes_api.ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS recipes_api.recipe_ingredients (
    recipe_id INT REFERENCES recipes_api.recipes(id) ON DELETE CASCADE,
    ingredient_id INT REFERENCES recipes_api.ingredients(id),
    quantity INTEGER NOT NULL,
    unit UNIT,
    preparation TEXT,
    PRIMARY KEY(recipe_id, ingredient_id)
);

CREATE TABLE IF NOT EXISTS recipes_api.recipe_instructions (
    recipe_id INTEGER REFERENCES recipes_api.recipes(id) ON DELETE CASCADE,
    position INTEGER NOT NULL,
    step_text TEXT NOT NULL,
    PRIMARY KEY(recipe_id, position)
);

INSERT INTO recipes_api.recipes(name, servings, cook_time_minutes)
VALUES ('toast', 1, 5)
ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name;

INSERT INTO recipes_api.ingredients(name)
VALUES ('bread')
ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name;

INSERT INTO recipes_api.recipe_ingredients (recipe_id, ingredient_id, quantity, unit, preparation)
VALUES
    (
        (SELECT id FROM recipes_api.recipes WHERE name = 'toast'),
        (SELECT id FROM recipes_api.ingredients WHERE name = 'bread'),
        1,
        'whole',
        'sliced'
    )
ON CONFLICT DO NOTHING;

INSERT INTO recipes_api.recipe_instructions(recipe_id, position, step_text)
VALUES
    (
        (SELECT id FROM recipes_api.recipes WHERE name = 'toast'),
        1,
        'Put bread in the toaster for 2 minutes'
    ),
    (
        (SELECT id FROM recipes_api.ingredients WHERE name = 'bread'),
        2,
        'Apply butter and toppings'
    )
ON CONFLICT DO NOTHING;
