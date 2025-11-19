# Script: cmd.sh
# Purpose:
#   Convenience wrapper to run database migrations, start the Go application,
#   and run tests for the project.
#
# Configuration:
#   DB_URL       - Postgres connection string used by the migration tool.
#                  Example: postgres://user:pass@host:port/dbname?sslmode=disable
#   MIGRATION_DIR- Filesystem path (relative to repo root) containing migration files.
#
# External dependencies:
#   - migrate (golang-migrate CLI) must be installed and on PATH for migrations.
#   - Go toolchain (go) must be installed to run the app and tests.
#   - A running PostgreSQL instance reachable using DB_URL.
#
# Provided actions (invoke as: ./cmd.sh <action>):
#   migrate:up    -> Apply all up migrations from MIGRATION_DIR to DB_URL.
#   migrate:down  -> Apply down migrations (rollback) from MIGRATION_DIR on DB_URL.
#   start         -> Run the Go application entrypoint (go run ./src/cmd/main.go).
#   test          -> Run all Go tests in the module (go test ./... -v).
#
# Usage examples:
#   ./cmd.sh migrate:up        # apply migrations
#   ./cmd.sh migrate:down      # rollback migrations
#   ./cmd.sh start             # start the application
#   ./cmd.sh test              # execute tests
#
# Notes and recommendations:
#   - Ensure the script is executable (chmod +x cmd.sh).
#   - Validate DB_URL and MIGRATION_DIR values before running migrations.
#   - Back up important data before running down/rollback migrations.
#   - CI/CD pipelines can call these actions to manage schema and run tests.
#   - The script delegates to external tools; check their output for errors.
#
# Exit behavior:
#   The script invokes external commands and will reflect the exit status of
#   those commands (non-zero exit indicates failure of the invoked action).
#!/bin/bash

# --- CONFIGURATION ---

DB_URL="postgres://root:root@localhost:5432/lopply_command_db?sslmode=disable"
MIGRATION_DIR="src/adapters/database/migrations/"

# --- FUNCTIONS ---

function migrate_up() {
    echo "Running migrations UP..."
    migrate -database "$DB_URL" -path "$MIGRATION_DIR" up
}

function migrate_down() {
    echo "Running migrations DOWN..."
    migrate -database "$DB_URL" -path "$MIGRATION_DIR" down
}

function run_app() {
    echo "Starting Go app..."
    go run ./src/cmd/main.go
}

function run_tests() {
    echo "Running Go tests..."
    go test ./... -v
}

function help() {
    echo "Usage: ./run.sh [migrate:up|migrate:down|start|test]"
    echo "  migrate:up    -> apply migrations"
    echo "  migrate:down  -> rollback migrations"
    echo "  start         -> run Go application"
    echo "  test          -> run all Go tests"
}

# --- MAIN ---

case "$1" in
    migrate:up)
        migrate_up
    ;;
    migrate:down)
        migrate_down
    ;;
    start)
        run_app
    ;;
    test)
        run_tests
    ;;
    *)
        help
    ;;
esac
