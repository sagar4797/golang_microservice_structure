#!/bin/sh
set -e

echo "Running db migrations"
/app/migrate -path /app/migration -database $DB_URL -verbose up

echo "Starting search service . . ."
./main