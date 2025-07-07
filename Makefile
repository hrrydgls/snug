include .env

MIGRATE=migrate -path db/migrations -database "$(DB_URL)"

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-drop:
	$(MIGRATE) drop -f

migrate-new:
	@read -p "Enter name: " name; \
	migrate create -ext sql -dir db/migrations -seq $$name
	
migrate-status:
	$(MIGRATE) version	