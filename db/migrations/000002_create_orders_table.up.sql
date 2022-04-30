CREATE TABLE IF NOT EXISTS public.orders(
    id serial PRIMARY KEY,
    user_id integer UNIQUE NOT NULL,
    coffee_id integer UNIQUE NOT NULL
);