#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    SELECT 'CREATE DATABASE bike_sharing_app' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'bike_sharing_app');

EOSQL