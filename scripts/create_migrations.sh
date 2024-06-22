#!/bin/bash

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="${{ secrets.POSTGRES_PASSWORD }}"
DB_NAME="librarymanagementsystem_test"

MIGRATIONS_PATH="../migrations"

for file in $MIGRATIONS_PATH/*.sql; do
    echo "Running migration: $file"
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f $file
done
