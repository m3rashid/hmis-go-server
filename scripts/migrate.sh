#! /bin/bash
# Path: internal/db/migrate.sh

cd internal/db

printf "\n::::: Generating Prisma Schema :::::\n"
yarn prisma migrate dev --name init --create-only --skip-generate

printf "\n::::: Moving generated SQL as schema input to SQLC :::::"
cp ./migrations/$(ls ./migrations | sort -k2 | head -n 1)/migration.sql ../sqlc/schema.sql

# echo "Archiving generated SQL : (to be used later)"
# cp ./migrations/$(ls ./migrations | sort -k2 | head -n 1)/migration.sql ./archives/$(ls ./migrations | sort -k2 | head -n 1).sql
rm -rf ./migrations || true

cd ../../
printf "\n::::: Generating SQLC :::::\n"
sqlc generate
printf "DONE\n"
