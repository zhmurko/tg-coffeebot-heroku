CREATE TABLE IF NOT EXISTS public.coffee(
   id serial PRIMARY KEY,
   name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO public.coffee (name)
    VALUES ('Espresso');

INSERT INTO public.coffee (name)
    VALUES ('Latte');