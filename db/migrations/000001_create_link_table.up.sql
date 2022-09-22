-- Active: 1663765487031@@127.0.0.1@5432@postgres@public
CREATE TABLE IF NOT EXISTS public.link (
    id serial PRIMARY KEY,
    long_link VARCHAR(200) NOT NULL,
    short_link VARCHAR(200) NOT NULL
);

INSERT INTO link (long_link,  short_link) VALUES ('https://www.google.ru/234', 'https://www.google.ru/dfg');
INSERT INTO link (long_link,  short_link) VALUES ('https://www.google.ru/23234', 'https://www.google.ru/234dfg');
INSERT INTO link (long_link,  short_link) VALUES ('https://www.google.ru/111234', 'https://www.google.ru/11213dfg');