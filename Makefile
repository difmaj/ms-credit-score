# Use this file to run the project with docker-compose
start: 
	@echo "- Checking air..."
	@if ! command -v air &> /dev/null; then \
			echo "MISSING DEPENDENCY: air is not installed"; \
			go install github.com/air-verse/air@latest; \
	fi
	air

# Use this command to generate swagger documentation
swag:
	swag fmt
	nice go run github.com/swaggo/swag/cmd/swag@latest init --parseDependency --parseInternal --generalInfo cmd/main.go --output docs --outputTypes go,json

# Use this command to run the application
goapp:
	docker-compose up --build score-worker

# Use this command to generate migration files
generate-migration-files:
	epoch=$$(date +%s00); \
	name=$$(echo $(name) | tr A-Z a-z | tr ' ' '_'); \
	touch ./migrations/up/$$epoch"_$$name.up.sql"; \
	touch ./migrations/down/$$epoch"_$$name.down.sql";

# Use this command to generate mock files for a given interface
generate-mock:
	@if [ -z "$(FILE)" ]; then \
		echo "Error: FILE variable must be set."; \
		exit 1; \
	fi
	@if [ ! -f "$(FILE)" ]; then \
		echo "Error: FILE does not exist."; \
		exit 1; \
	fi
	mockgen -source=$(FILE) -destination=$(shell dirname $(FILE))/mock/$(notdir $(FILE)) -package=mock