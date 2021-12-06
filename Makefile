# define variables
DB_NAME=db/db.sqlite
MIGRATION_DIR=backend/db/migrations
args = $(filter-out $@,$(MAKECMDGOALS))


# start the backend
start:
	go run main.go


# database migrations
migration:
	goose -dir ${MIGRATION_DIR} create $(call args,defaultstring) sql

migrate-status:
	goose -dir ${MIGRATION_DIR} sqlite3 ${DB_NAME} status

migrate-up:
	goose -dir ${MIGRATION_DIR} sqlite3 ${DB_NAME} up

migrate-down:
	goose -dir ${MIGRATION_DIR} sqlite3 ${DB_NAME} down

migrate-rollback:
	goose -dir ${MIGRATION_DIR} sqlite3 ${DB_NAME} reset

migrate-reset:
	goose -dir ${MIGRATION_DIR} sqlite3 ${DB_NAME} reset
