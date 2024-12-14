-- Создаем новую базу данных

CREATE DATABASE calendar_db
    WITH TEMPLATE = template0
    ENCODING = 'UTF8'
    LOCALE_PROVIDER = libc
    LOCALE = 'en_US.utf8'
    OWNER = postgres;
