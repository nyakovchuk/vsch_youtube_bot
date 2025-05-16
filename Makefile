include .env.prod

LOCAL_BIN:=$(CURDIR)/bin

IMAGE_TELEGRAM=tg-vsch-youtube-bot
GOOSE_DBSTRING=./db.sqlite

ARGUMENTS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

build-telegram:
	docker build -f Dockerfile.telegram -t $(IMAGE_TELEGRAM) .

run-telegram:
	docker run --rm $(IMAGE_TELEGRAM)

install-migration:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.24.2

migration-create:
	$(LOCAL_BIN)/goose create -dir $(MIGRATION_DIR) $(ARGUMENTS) sql

migration-status:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) $(DB_DRIVER) status 

migration-up:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(GOOSE_DBSTRING) up

migration-up-to:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) $(DB_DRIVER) up-to $(ARGUMENTS) -v

migration-down:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) $(DB_DRIVER) down

migration-down-to:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) $(DB_DRIVER) down-to $(ARGUMENTS)

migration-fresh:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) $(DB_DRIVER) down-to 0

#%::
#	@true
