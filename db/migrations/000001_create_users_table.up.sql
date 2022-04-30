CREATE TABLE IF NOT EXISTS public.users(
    id serial PRIMARY KEY,
    telegarm_id integer UNIQUE NOT NULL,
    username VARCHAR (50) NOT NULL
);