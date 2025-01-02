set dotenv-load

# Definice proměnných
COUNT := "1"
MIGRATE_CMD := "docker run --rm -v $(pwd)/database/migrations:/migrations --network mimir_default migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@db:5432/postgres?sslmode=disable"

migrate_up:
    @{{MIGRATE_CMD}} -verbose up

migrate_down:
    @{{MIGRATE_CMD}} -verbose down {{COUNT}}

migrate_to_version:
    @{{MIGRATE_CMD}} -verbose goto {{COUNT}}

migrate_force:
    @{{MIGRATE_CMD}} -verbose force {{COUNT}}

drop_database:
    @{{MIGRATE_CMD}} -verbose drop -f
