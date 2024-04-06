FROM postgres:latest

RUN echo "CREATE TABLE IF NOT EXISTS users (\
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\
    email VARCHAR(255) UNIQUE,\
    username VARCHAR(255) UNIQUE,\
    first_name VARCHAR(255),\
    last_name VARCHAR(255),\
    phone VARCHAR(20) UNIQUE,\
    password VARCHAR(255)\
    );" > /docker-entrypoint-initdb.d/init.sql

