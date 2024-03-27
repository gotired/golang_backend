FROM postgres:latest

# Create a SQL initialization script
RUN echo "CREATE TABLE IF NOT EXISTS users (" \
    "id SERIAL PRIMARY KEY," \
    "username VARCHAR(50) NOT NULL," \
    "email VARCHAR(100) NOT NULL," \
    "created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP" \
    ");" > /docker-entrypoint-initdb.d/init.sql

