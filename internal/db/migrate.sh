#! /bin/bash
# Path: internal/db/migrate.sh

yarn prisma migrate dev --name init --create-only
cp ./migrations/$(ls ./migrations | sort -k2 | head -n 1)/migration.sql ../sqlc/schema.sql
cp ./migrations/$(ls ./migrations | sort -k2 | head -n 1)/migration.sql ./archives/$(ls ./migrations | sort -k2 | head -n 1).sql
rm -rf ./migrations || true
