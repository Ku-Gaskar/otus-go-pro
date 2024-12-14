
-- Создаем новую базу данных
-- CREATE DATABASE calendar_db
--     WITH TEMPLATE = template0
--     ENCODING = 'UTF8'
--     LOCALE_PROVIDER = libc
--     LOCALE = 'en_US.utf8'
--     OWNER = postgres;

--\connect calendar_db

-- -- Установка общих параметров
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SET check_function_bodies = false;
-- SET xmloption = content;
-- SET client_min_messages = warning;

-- Создание таблиц с первичными и внешними ключами
CREATE TABLE IF NOT EXISTS public.users (
                              id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                              user_name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS public.events (
                               id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               title character varying,
                               start_event time with time zone,
                               end_event time with time zone,
                               description text,
                               "user" bigint REFERENCES public.users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.notification (
                                     id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                     title_event character varying,
                                     time_event time with time zone,
                                     event bigint REFERENCES public.events(id) ON DELETE CASCADE
);

-- Добавление дефолтного пользователя если таблица пустая
-- INSERT INTO public.users (user_name) VALUES ('admin');
INSERT INTO public.users (user_name) SELECT 'admin'
WHERE NOT EXISTS (
    SELECT 1 FROM public.users
);